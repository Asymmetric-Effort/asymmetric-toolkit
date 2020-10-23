package encryption

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestDecrypt(t *testing.T) {
	var key Key
	cipherText:="c819da91a775de90daedce31b780e58ffe65764104755807c4a286671e494955b5d6aff0275885b7871697529e381" +
				"fdcb833e1e21a9ec4a94664927e9f2670747bd36509a1a4961263875af9d2d4b074e1a1749b70f3993abba92431f6c" +
		"bd35df2c1c15d03953206797dc40a09a6f3e3050976c02490259a1839f8a17b8de60311ac1c943026425197625f7c65be318b7" +
		"9b2e077aad42f64dea23b95b0da74ccde67f428a16f6c7920120fe3c895d6a4fac0d1894d92418ccdb84d14109714c41447091" +
		"18a2cc74f417168e804067b9d2ca062d5e12d582bdf82e205d5d6444de9ca85008fd9d3e5ca976941a307f661"

	plainText := "We seek security through transparency by openly identifying, recognizing, prioritizing and acting " +
		"to resolve the weaknesses we find in our systems while balancing the risks against opportunities for future" +
		"growth."
	password := "Trump_Thinks_This_Is_Secure_Unless_A_hacker_Gets_15_Percent_Of_The_Password"
	key.Set(&password)
	decryptedText := Decrypt(&cipherText, &key)
	errors.Assert(plainText == *decryptedText,"plaintext does not match decrypted cipher")
}
