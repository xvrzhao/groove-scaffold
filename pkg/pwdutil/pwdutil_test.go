package pwdutil

import "testing"

func TestSaltHashPwd(t *testing.T) {
	p, s := SaltHashPwd("admin123", 8)
	t.Log(p, s)
}
