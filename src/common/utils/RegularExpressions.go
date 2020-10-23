package utils

/*
	Some common regular expressions.
*/

const (
	RegExDotPlusMan = ".+" // A tribute to Dual Core.
	RegExTcpUdp = `(udp|tcp){1}:`
	RegExAddr   = `(([01]{0,1}[0-9]{1,2}|2[0-4][0-9]|25[0-5])\\.){3}(([01]{0,1}[0-9]{1,2}|2[0-4][0-9]|25[0-5])){1}:`
	RegExPort   = `[0-5]{0,1}[0-9]{1,4}|6[0-4][0-9]{3}|65[0-4]{1}[0-9]{2}|655[0-2][0-9]|6553[0-6]`

	//Server descriptors are defined as <protocol>:<ipaddress>:<port>
	RegExServerString = `^` + RegExTcpUdp + `:` + RegExAddr + `:` + RegExPort + `$`
)
