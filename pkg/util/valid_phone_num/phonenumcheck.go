package validphonenum

import (
	"regexp"
	"strconv"
)

const regex = `^1[3456789]\d{9}$`

var re = regexp.MustCompile(regex)

func IsPhoneNumCN(phoneStr string) bool {
	return re.MatchString(phoneStr)
}

func FooCheckIsNum(phoneStr string) bool {
	_, err := strconv.Atoi(phoneStr)
	return err == nil
}
