package env

import (
	"os"

	"github.com/joho/godotenv"
)

func Get(name string) string {
    file, err := godotenv.Read(".env")

    if err == nil { return file[name] }
    return os.Getenv(name)
}

func IsProd() bool {
    return Get("GOENV_PROD") == "true"
}

func Port() string {
    return ":" + Get("PORT")
}