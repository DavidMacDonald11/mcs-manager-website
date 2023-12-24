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

func DbPath() string {
	return Get("DB_PATH")
}

func SessionKey() string {
	return Get("SESSION_KEY")
}

func AccessControlOrigin() []string {
	return []string{Get("ACCESS_CONTROL_ORIGIN")}
}

func BootstrapInviteCode() string {
	return Get("BOOTSTRAP_INVITE_CODE")
}

func Admin() string {
	return Get("ADMIN")
}
