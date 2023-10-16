package env

import "os"

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

func JwtSecret() string {
    if !IsProd() { return "Y1]iH*[~swp/uPNoP7+S-q9wiBR;Oy@@:IH!OLziOPv[l;lu,J" }
    return os.Getenv("JWT_SECRET")
}
