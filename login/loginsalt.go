package login

import (
	"crypto/md5"
	"fmt"
	"io"
)

func Addsalt(password string) string {
	h := md5.New()
	io.WriteString(h, password)
	//log.Println(h)
	pwmd5 := fmt.Sprintf("%x", h.Sum([]byte("eric.liu")))

	salt1 := "!@#$%^&*()"
	salt2 := "1234567890"

	io.WriteString(h, salt1)
	io.WriteString(h, salt2)
	io.WriteString(h, pwmd5)

	last := fmt.Sprintf("%x", h.Sum([]byte(salt1)))

	return last

}
