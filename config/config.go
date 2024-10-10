package config

type Config struct {
	// MySQL Setup
	MySQLHost     string `mapstructure: "MYSQL_HOST"`
	MySQLPort     string `mapstructure: "MYSQL_PORT"`
	MySQLUser     string `mapstructure: "MYSQL_USER"`
	MySQLPassword string `mapstructure: "MYSQL_ROOT_PASSWORD"`
	MySQLDatabase string `mapstructure: "MYSQL_DATABASE"`

	// Redis Setup
	RedisHost string `mapstructure: "REDIS_HOST"`
	RedisPort string `mapstructure: "REDIS_PORT"`
}