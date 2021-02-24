package auth

import (
	"crypto/rand"
	"github.com/qnfnypen/gokit/internal/domain"
)

// GenerateCaptcha 生成验证码
func GenerateCaptcha(figure int) string {
	captcha := make([]byte,figure)
	
	rand.Read(captcha)

	for i,v := range captcha {
		captcha[i] = domain.NumTable[int(v)%len(domain.NumTable)]
	}

	return string(captcha)
}