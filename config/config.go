package config

import (
	"os"
	"strconv"
	"strings"
	//"os"
	//"reflect"
	//"strings"
	//cfgloader "gitlab.autodoc.dev/dgt/cross-cutting-concerns/libraries/go/configloader"
	//"github.com/stretchr/testify/assert/yaml"
)

const (
	ErrorLoadingEnvFile = "error loading .env file"
	ErrorParsingEnvFile = "unable to parse environment variables: %w"
	ConsumerYamlFile    = "conf/consumers.yaml"
)

type Config struct {
	App         AppConfig
	Server      ServerConfig
	Database    DatabaseConfig
	Redis       RedisConfig
	Kafka       KafkaConfig
	ExternalAPI ExternalAPIConfig
}

type AppConfig struct {
	Name        string
	Environment string
}

type ServerConfig struct {
	Host string
	Port string
}

type DatabaseConfig struct {
	Host            string
	Port            int
	User            string
	Password        string
	Name            string
	SSLMode         string
	MaxConnections  int32
	MinConnections  int32
	MaxConnLifetime int
}

type RedisConfig struct {
	Host     string
	Port     int
	Password string
	DB       int
}

type KafkaConfig struct {
	Brokers  []string
	Username string
	Password string
	ClientID string
}

type ExternalAPIConfig struct {
	BaseURL string
	APIKey  string
	Timeout int
}

/*
type PostgresDBConfig struct {
	User     string `env:"POSTGRES_USER,required"`
	Password string `env:"POSTGRES_PASSWORD,required"`
	Address  string `env:"POSTGRES_ADDRESS,required"`
	Port     uint   `env:"POSTGRES_PORT,required"`
	Database string `env:"POSTGRES_DATABASE,required"`
}

type KafkaConfig struct {
	Broker       string `env:"KAFKA_CONNECT_STRING,required"`
	AuthUsername string `env:"KAFKA_AUTH_USERNAME,required"`
	AuthPassword string `env:"KAFKA_AUTH_PASSWORD,required"`
}

type ClickHouseConfig struct {
	Address  string `env:"CLICKHOUSE_ADDRESS,required"`
	Username string `env:"CLICKHOUSE_USERNAME,required"`
	Password string `env:"CLICKHOUSE_PASSWORD,required"`
	Database string `env:"CLICKHOUSE_DATABASE,required"`
}

type HTTPConfig struct {
	Port string `env:"HTTP_PORT,required"`
}

type AASAuthentication struct {
	ClientID     string `env:"AAS_CLIENT_ID,required"`
	ClientSecret string `env:"AAS_CLIENT_SECRET,required"`
	APIURL       string `env:"AAS_API_URL,required"`
	HostURL      string `env:"AAS_HOST_URL,required"`
	FrontURL     string `env:"AAS_FRONT_URL,required"`
	SSOURL       string `env:"AAS_SSO_URL,required"`
}

type NewRelicConfig struct {
	LicenseKey string `env:"NEW_RELIC_LICENSE_KEY"`
	AppName    string `env:"NEW_RELIC_APP_NAME"`
}

type BucketConfig struct {
	BucketName           string `env:"BUCKET_NAME,required"`
	BucketServiceAccount string `env:"BUCKET_SERVICE_ACCOUNT"`
}

type AWSAPIConfig struct {
	BaseURL string `env:"AWS_API_URL,required"`
}

type OpsGetterConfig struct {
	BaseURL     string `env:"OPS_GETTERS_URL,required"`
	StaticToken string `env:"OPS_GETTERS_STATIC_TOKEN,required"`
}

type MessageConsumerConfig struct {
	Broker        string `yaml:"Broker"`
	Topic         string `yaml:"Topic"`
	GroupID       string `yaml:"GroupId"`
	HandlerMethod string `yaml:"HandlerMethod"`
	Handler       reflect.Value
}

type ConsumersConfig struct {
	Consumers map[string]*MessageConsumerConfig `yaml:"consumers"`
}

type Consumer struct {
	Config *MessageConsumerConfig
	//Dialer *kafka.Dialer
}*/

func Load() Config {
	return Config{
		App: AppConfig{
			Name:        getEnv("APP_NAME", "my-service"),
			Environment: getEnv("APP_ENV", "local"),
		},
		Server: ServerConfig{
			Host: getEnv("SERVER_HOST", "0.0.0.0"),
			Port: getEnv("SERVER_PORT", "8080"),
		},
		Database: DatabaseConfig{
			Host:            getEnv("DB_HOST", "localhost"),
			Port:            getEnvAsInt("DB_PORT", 5432),
			User:            getEnv("DB_USER", "postgres"),
			Password:        getEnv("DB_PASSWORD", ""),
			Name:            getEnv("DB_NAME", "postgres"),
			SSLMode:         getEnv("DB_SSLMODE", "disable"),
			MaxConnections:  int32(getEnvAsInt("DB_MAX_CONNECTIONS", 20)),
			MinConnections:  int32(getEnvAsInt("DB_MIN_CONNECTIONS", 5)),
			MaxConnLifetime: getEnvAsInt("DB_MAX_CONN_LIFETIME", 3600),
		},
		Redis: RedisConfig{
			Host:     getEnv("REDIS_HOST", "localhost"),
			Port:     getEnvAsInt("REDIS_PORT", 6379),
			Password: getEnv("REDIS_PASSWORD", ""),
			DB:       getEnvAsInt("REDIS_DB", 0),
		},
		Kafka: KafkaConfig{
			Brokers:  strings.Split(getEnv("KAFKA_BROKERS", "localhost:9092"), ","),
			Username: getEnv("KAFKA_USERNAME", ""),
			Password: getEnv("KAFKA_PASSWORD", ""),
			ClientID: getEnv("KAFKA_CLIENT_ID", "my-service"),
		},
		ExternalAPI: ExternalAPIConfig{
			BaseURL: getEnv("API_BASE_URL", ""),
			APIKey:  getEnv("API_KEY", ""),
			Timeout: getEnvAsInt("API_TIMEOUT", 30),
		},
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	valueStr := getEnv(key, "")
	if valueStr == "" {
		return defaultValue
	}

	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return defaultValue
	}

	return value
}

/*
func GetAppConfigs() (*cfgloader.Configuration, error) {
	cfg, err := cfgloader.Load()
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func GetEnvs() (*Configurations, error) {
	environmentConfig, err := Load()
	if err != nil {
		return nil, err
	}

	return environmentConfig, nil
}

func GetConsumerConfigs() (*ConsumersConfig, error) {
	data, err := os.ReadFile(ConsumerYamlFile)
	if err != nil {
		return nil, err
	}

	var consumersConfig ConsumersConfig
	if err = yaml.Unmarshal(data, &consumersConfig); err != nil {
		return nil, err
	}

	for _, consumerConfig := range consumersConfig.Consumers {
		if strings.Contains(consumerConfig.Broker, "os:env:") {
			envKey := strings.Replace(consumerConfig.Broker, "os:env:", "", 1)
			envValue := os.Getenv(envKey)
			consumerConfig.Broker = envValue
		}
	}

	return &consumersConfig, nil
}
*/
