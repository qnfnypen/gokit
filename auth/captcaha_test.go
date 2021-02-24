package auth

import (
	"strconv"
	"testing")


func TestGenerateCaptcha(t *testing.T) {
	cap1 := GenerateCaptcha(6)
	cap2 := GenerateCaptcha(6)

	_,err := strconv.ParseInt(cap1,10,64)
	if err != nil {
		t.Error(err)
	}

	if len(cap1) != 6 && len(cap2) != 6 {
		t.Errorf("The length is 6,but is %d %d",len(cap1),len(cap2))
	}

	if cap1 == cap2 {
		t.Error("It should be random, but it's not")
	}
}