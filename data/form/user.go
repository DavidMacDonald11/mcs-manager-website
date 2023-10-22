package form

import (
	"net/http"
	"regexp"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

const INTERNAL_ERROR = "Internal server error"

type User struct {
    Username string `json:"Username"`
    Password string `json:"Password"`
    Code string `json:"Code"`
}

func (u User) VerifyUsername() *bool {
    re := regexp.MustCompile(`^\w{3,16}$`)
    ok := re.MatchString(u.Username)
    if !ok { return &ok }

    url := "https://api.mojang.com/users/profiles/minecraft/"
    res, err := http.Get(url + u.Username)

    if err != nil { return nil }
    defer res.Body.Close()

    if res.StatusCode == 400 { return nil }
    ok = res.StatusCode == 200
    return &ok
}

func (u User) VerifyPassword() bool {
    p := []byte(u.Password)
    if len(p) < 16 || len(p) > 72 { return false }
    return true
}

func (u User) HashPassword() *string {
    hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
    h := string(hash)

    if err == nil { return &h }
    return nil
}

func (u User) RenderAuth(c *fiber.Ctx, tmpl string, err string) error {
    return c.Render(tmpl, fiber.Map {
        "Username": u.Username,
        "Error": err,
    })
}

func ParseAuth(c *fiber.Ctx, tmpl string, needed int) *User {
    u := new(User)
    err := c.BodyParser(u)

    switch {
    case err != nil:
        u.RenderAuth(c, tmpl, INTERNAL_ERROR)
    case needed > 0 && u.Username == "":
        u.RenderAuth(c, tmpl, "Must provide a username")
    case needed > 1 && u.Password == "":
        u.RenderAuth(c, tmpl, "Must provide a password")
    case needed > 2 && u.Code == "":
        u.RenderAuth(c, tmpl, "Must provide an invite code")
    default:
        return u
    }

    return nil
}
