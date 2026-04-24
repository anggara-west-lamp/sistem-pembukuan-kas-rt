package config

import (
    "errors"
    "os"
)

type Config struct {
    AppPort string
    AppEnv  string

    DBHost     string
    DBPort     string
    DBUser     string
    DBPassword string
    DBName     string
    DBSSLMode  string

    JWTSecret     string
    JWTTtlMinutes int

    StorageEndpoint  string
    StorageAccessKey string
    StorageSecretKey string
    StorageUseSSL    bool
    StorageBucket    string
}

func Load() (*Config, error) {
    cfg := &Config{
        AppPort:     getenv("APP_PORT", "8080"),
        AppEnv:      getenv("APP_ENV", "development"),
        DBHost:      getenv("DB_HOST", "localhost"),
        DBPort:      getenv("DB_PORT", "5432"),
        DBUser:      getenv("DB_USER", "postgres"),
        DBPassword:  getenv("DB_PASSWORD", "postgres"),
        DBName:      getenv("DB_NAME", "kasrt"),
        DBSSLMode:   getenv("DB_SSLMODE", "disable"),
        JWTSecret:   getenv("JWT_SECRET", ""),
        JWTTtlMinutes: atoi(getenv("JWT_TTL_MINUTES", "60")),
        StorageEndpoint:  getenv("STORAGE_ENDPOINT", ""),
        StorageAccessKey: getenv("STORAGE_ACCESS_KEY", ""),
        StorageSecretKey: getenv("STORAGE_SECRET_KEY", ""),
        StorageUseSSL:    getenv("STORAGE_USE_SSL", "false") == "true",
        StorageBucket:    getenv("STORAGE_BUCKET", "kasrt-proof"),
    }
    if cfg.JWTSecret == "" {
        return nil, errors.New("JWT_SECRET required")
    }
    return cfg, nil
}

func getenv(k, d string) string {
    if v := os.Getenv(k); v != "" { return v }
    return d
}

func atoi(s string) int {
    n := 0
    for _, c := range s {
        if c < '0' || c > '9' { return n }
        n = n*10 + int(c-'0')
    }
    return n
}

