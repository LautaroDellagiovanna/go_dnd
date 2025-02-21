package config

type Config struct {
	ServerAddress string
	DatabaseURL   string
}

func LoadConfig() Config {
	return Config{
		//ServerAddress: os.Getenv("SERVER_ADDRESS"),
		ServerAddress: "localhost:9090",
		//DatabaseURL:   os.Getenv("DATABASE_URL"),
		DatabaseURL: "root:manager1@tcp(localhost:3306)/go_dnd",
	}
}
