package configs

import "os"

func ternaryMap(value string, defaultValue string) string {
	return map[bool]string{true: value, false: defaultValue}[len(value) != 0]
}

// Envs has values for environment variables and the defaults for them
var Envs = map[string]string{
	"PROXY_DESTINATION": ternaryMap(os.Getenv("PROXY_DESTINATION"), "localhost:8080"),
	"SERVER_PORT": ternaryMap(os.Getenv("SERVER_PORT"), "8080"),
	"DB_HOST":     ternaryMap(os.Getenv("DB_HOST"), "localhost"),
	"DB_PORT":     ternaryMap(os.Getenv("DB_PORT"), "5432"),
	"DB_USERNAME": ternaryMap(os.Getenv("DB_USERNAME"), "postgres"),
	"DB_PASSWORD": ternaryMap(os.Getenv("DB_PASSWORD"), "1234"),
	"DB_DATABASE": ternaryMap(os.Getenv("DB_DATABASE"), "database"),
	"DB_SSLMODE":  ternaryMap(os.Getenv("DB_SSLMODE"), "disable"),
}
