package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config 配置结构体
type Config struct {
	// API相关配置
	ApiHost  string
	ApiToken string
	ApiTime  string

	// 用户相关配置
	UserID string

	// 车辆相关配置
	TotalSeats int
}

var AppConfig *Config

// LoadConfig 加载配置
func LoadConfig() {
	// 加载.env文件
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	totalSeats, err := strconv.Atoi(getEnv("TOTAL_SEATS", "51"))
	if err != nil {
		log.Fatal("Invalid TOTAL_SEATS value")
	}

	AppConfig = &Config{
		ApiHost:    getEnv("API_HOST", "hqapp1.bit.edu.cn"),
		ApiToken:   getEnv("API_TOKEN", ""),
		ApiTime:    getEnv("API_TIME", ""),
		UserID:     getEnv("USER_ID", ""),
		TotalSeats: totalSeats,
	}
}

// getEnv 获取环境变量，如果不存在则返回默认值
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
