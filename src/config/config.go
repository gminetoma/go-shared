package config

import (
	"log/slog"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type (
	Config struct {
		Host                    string
		Port                    string
		Protocol                string
		DatabaseURL             string
		JWTSecret               string
		AccessTokenExpiry       time.Duration
		RefreshTokenExpiry      time.Duration
		RefreshTokenGracePeriod time.Duration
		Env                     string
	}
)

const (
	EnvProduction  = "production"
	EnvDevelopment = "development"
	EnvTest        = "test"
)

func parseDuration(value string, fallback time.Duration) time.Duration {
	if value == "" {
		return fallback
	}

	duration, err := time.ParseDuration(value)
	if err != nil {
		return fallback
	}

	return duration
}

func LoadEnv() Config {
	godotenv.Load()

	// Domain
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	host := os.Getenv("HOST")
	if host == "" {
		host = "localhost"
	}

	protocol := os.Getenv("PROTOCOL")
	if protocol == "" {
		protocol = "http"
	}

	// Database
	databaseURL := os.Getenv("DATABASE_URL")

	// Auth Tokens
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		slog.Error("JWT_SECRET is required")
		os.Exit(1)
	}

	accessTokenExpiry := os.Getenv("ACCESS_TOKEN_EXPIRY")
	refreshTokenExpiry := os.Getenv("REFRESH_TOKEN_EXPIRY")
	refreshTokenGracePeriod := os.Getenv("REFRESH_TOKEN_GRACE_PERIOD")

	// Env
	env := os.Getenv("ENV")
	if env == "" {
		env = EnvProduction
	}

	return Config{
		Host:                    host,
		Port:                    port,
		Protocol:                protocol,
		DatabaseURL:             databaseURL,
		JWTSecret:               jwtSecret,
		AccessTokenExpiry:       parseDuration(accessTokenExpiry, 15*time.Minute),
		RefreshTokenExpiry:      parseDuration(refreshTokenExpiry, 7*24*time.Hour),
		RefreshTokenGracePeriod: parseDuration(refreshTokenGracePeriod, 10*time.Second),
		Env:                     env,
	}
}
