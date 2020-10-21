package encryption_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"fmt"
	"testing"
)

func TestEncrypt(t *testing.T){
	plainText:="We seek security through transparency by openly identifying, recognizing, prioritizing and acting " +
		"to resolve the weaknesses we find in our systems while balancing the risks against opportunities for future" +
		"growth."
	password:="Trump_Thinks_This_Is_Secure_Unless_A_hacker_Gets_15_Percent_Of_The_Password"
	var key Key
	key.Set(&password)
	cipherText:= Encrypt(&plainText,&key)
	fmt.Println("cipherText:",*cipherText)
	errors.Assert(len(*cipherText)>1,"ciphertext should not be empty.")
}