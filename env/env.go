package env

import (
	"os"
	"strings"
)

func IsProd() bool {
    return os.Getenv("GOENV_PROD") == "true"
}

func Port() string {
    if !IsProd() { return ":3000" }
    return ":" + os.Getenv("PORT")
}

func DbPath() string {
    if !IsProd() { return "mcsm.db" }
    return os.Getenv("DB_PATH")
}

func Admins() []string {
    if !IsProd() { return []string{"The13Doctors"} }
    res := strings.Split(os.Getenv("ADMINS"), ",")

    for i, s := range res {
        res[i] = strings.Trim(s, " ")
    }

    return res
}
