package form

import (
	"net/http"
	"regexp"
	"unicode/utf8"
)

func VerifyLogin(username, password string) string {
	err := VerifyUsername(username)

	if err != "" {
		return err
	}

	err = VerifyPassword(password)

	if err != "" {
		return err
	}

	return ""
}

func VerifySignup(username, password, code string) string {
	err := VerifyLogin(username, password)

	if err != "" {
		return err
	}

	return VerifyInviteCode(code)
}

func VerifyUsername(username string) string {
	if username == "" {
		return "Username must be provided"
	}

	re := regexp.MustCompile(`^\w{3,16}$`)

	if !re.MatchString(username) {
		return "Username is not valid"
	}

	url := "https://api.mojang.com/users/profiles/minecraft/"
	res, err := http.Get(url + username)

	if err != nil || res.StatusCode == 400 {
		return "Internal server error"
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return "Username does not exist"
	}

	return ""
}

func VerifyPassword(password string) string {
	if password == "" {
		return "Password must be provided"
	}

	if utf8.RuneCountInString(password) < 10 {
		return "Password length must be at least 10 "
	}

	return ""
}

func VerifyInviteCode(code string) string {
	if code == "" {
		return "Invite Code must be provided"
	}

	return ""
}
