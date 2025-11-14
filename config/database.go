package config

import (
	"ecom-go/models"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DB *gorm.DB

// LoadEnv loads .env; returns no error if file missing (env may come from docker compose)
func LoadEnv() {
	_ = godotenv.Load()
}

// ConnectDatabase connects to MSSQL using environment variables and sets global DB.
func ConnectDatabase() {
	LoadEnv()

	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "1433")
	user := getEnv("DB_USER", "SA")
	pass := getEnv("DB_PASSWORD", "")
	name := getEnv("DB_NAME", "ecom_go")

	// Build DSN: sqlserver://username:password@host:port?database=dbname
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", user, pass, host, port, name)

	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ failed to connect to MSSQL: %v\nDSN=%s", err, maskDSN(dsn))
	}

	// set connection pooling on underlying sql.DB
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("❌ failed to get sql.DB from gorm: %v", err)
	}

	// sensible defaults — tune for production
	sqlDB.SetMaxOpenConns(getEnvInt("DB_MAX_OPEN_CONNS", 20))
	sqlDB.SetMaxIdleConns(getEnvInt("DB_MAX_IDLE_CONNS", 10))
	sqlDB.SetConnMaxLifetime(time.Hour)

	DB = db

	DB.AutoMigrate(&models.User{})

	log.Println("📦 User table migrated.")
	log.Println("✅ Connected to MSSQL database.")
}

// helper: read environment variable with default
func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

// helper: read integer env with default
func getEnvInt(key string, fallback int) int {
	if v := os.Getenv(key); v != "" {
		if i, err := strconv.Atoi(v); err == nil {
			return i
		}
	}
	return fallback
}

// maskDSN hides password for logging
func maskDSN(dsn string) string {
	// very simple mask: remove password between : and @
	// example: sqlserver://SA:MyPass@localhost:1433?database=...
	start := "://"
	i := 0
	if pos := findNth(dsn, start, 1); pos >= 0 {
		i = pos + len(start)
	} else {
		return dsn
	}
	// find '@' after start
	at := -1
	for j := i; j < len(dsn); j++ {
		if dsn[j] == '@' {
			at = j
			break
		}
	}
	if at == -1 {
		return dsn
	}
	// find ':' between start and @
	col := -1
	for j := i; j < at; j++ {
		if dsn[j] == ':' {
			col = j
			break
		}
	}
	if col == -1 {
		return dsn
	}
	return dsn[:col+1] + "****" + dsn[at:]
}

func findNth(s, sub string, n int) int {
	pos := -1
	start := 0
	for i := 0; i < n; i++ {
		idx := -1
		if start < len(s) {
			idx = indexOf(s[start:], sub)
		}
		if idx == -1 {
			return -1
		}
		pos = start + idx
		start = pos + len(sub)
	}
	return pos
}

func indexOf(s, sub string) int {
	for i := 0; i+len(sub) <= len(s); i++ {
		if s[i:i+len(sub)] == sub {
			return i
		}
	}
	return -1
}
