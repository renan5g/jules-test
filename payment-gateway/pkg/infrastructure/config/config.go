package config

// Example structure - expand as needed
type Config struct {
	ServerPort string `mapstructure:"SERVER_PORT"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`
	StripeKey  string `mapstructure:"STRIPE_KEY"`
}

// TODO: Add function to load config from file or environment variables (e.g., using Viper)
// func LoadConfig(path string) (*Config, error)
