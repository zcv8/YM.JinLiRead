package common

import (
	"regexp"
)

//验证邮箱
func ValidEmail(email string) (match bool) {
	match, _ = regexp.MatchString("^([a-zA-Z0-9_\\.\\-])+\\@(([a-zA-Z0-9\\-])+\\.)+([a-zA-Z0-9]{2,4})+$", email)
	return
}

func ValidPhone(phone string) (match bool) {
	match, _ = regexp.MatchString("^1\\d{10}$", phone)
	return
}
