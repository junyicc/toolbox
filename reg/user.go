package reg

import "regexp"

const (
	emailRegStr    = ``
	phoneRegStr    = `^(?:\+?86)?1(?:3\d{3}|5[^4\D]\d{2}|8\d{3}|7(?:[35678]\d{2}|4(?:0\d|1[0-2]|9\d))|9[189]\d{2}|66\d{2})\d{6}$`
	usernameRegStr = ``
	idCardRegStr   = ``
)

var (
	emailReg    = regexp.MustCompile(emailRegStr)
	phoneReg    = regexp.MustCompile(phoneRegStr)
	usernameReg = regexp.MustCompile(usernameRegStr)
	idCardReg   = regexp.MustCompile(idCardRegStr)
)

func IsValidEmail(email string) bool {
	return emailReg.MatchString(email)
}

func IsValidPhone(phone string) bool {
	return phoneReg.MatchString(phone)
}

func IsValidUsername(username string) bool {
	return usernameReg.MatchString(username)
}

func IsValidIdCard(idCard string) bool {
	return idCardReg.MatchString(idCard)
}
