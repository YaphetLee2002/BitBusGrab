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

	// 邮件相关配置
	Mail MailConfig
}

// MailConfig 邮件配置结构体
type MailConfig struct {
	Username        string
	Password        string
	DefaultEncoding string
	Host            string
	Port            int
	UserEmail       string // 用户接收邮件的邮箱地址
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

	mailPort, err := strconv.Atoi(getEnv("MAIL_PORT", "465"))
	if err != nil {
		log.Fatal("Invalid MAIL_PORT value")
	}

	AppConfig = &Config{
		ApiHost:    getEnv("API_HOST", "hqapp1.bit.edu.cn"),
		ApiToken:   getEnv("API_TOKEN", ""),
		ApiTime:    getEnv("API_TIME", ""),
		UserID:     getEnv("USER_ID", ""),
		TotalSeats: totalSeats,
		Mail: MailConfig{
			Username:        getEnv("MAIL_USERNAME", "registercode@yaphet.top"),
			Password:        getEnv("MAIL_PASSWORD", "AYaphet677958"),
			DefaultEncoding: getEnv("MAIL_DEFAULT_ENCODING", "UTF-8"),
			Host:            getEnv("MAIL_HOST", "smtpdm.aliyun.com"),
			Port:            mailPort,
			UserEmail:       getEnv("USER_EMAIL", ""),
		},
	}
}

// getEnv 获取环境变量，如果不存在则返回默认值
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
