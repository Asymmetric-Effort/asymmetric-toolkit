package misc

func BoolToByte(b bool) byte {
	if b {
		return byte(1)
	} else {
		return byte(0)
	}
}
