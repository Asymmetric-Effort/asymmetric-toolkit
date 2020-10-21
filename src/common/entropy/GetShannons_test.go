package entropy_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/entropy"
	"testing"
)

const (
	sshKey2048 = "b3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAABFwAAAAdzc2gtcn" +
		"NhAAAAAwEAAQAAAQEAoabfH73CYjH6v5Qw2lUCIWUkfm997GCwldNUQODocqdscPh/7p1Ii7" +
		"6s8E/VxkPDV1PVSbgVJ7gEjKnWQfKx3BcDAw0IdNLZrZ7i51JRuhB1WWRB+2k/6NJeNNFwcI" +
		"blCIZEUKn16Q90Sxkv/W9WiFazFCRREHrXwOC3vMnE0mA00vO4EUx69k+9U92utB9penMR1o" +
		"yys+ogLMpjKKcsIri88P5DQSjPxUtZ2srPtJ9f40lXZVuOVN6lnQ2ur8YBt0dZPclgMuuH8t" +
		"H9nBXmnZ3bIfZpsOPoCTI/7zhujo4QfWX+cr9oq70e/c/MEf0la1mgUXrZs6RlE04NUvmFzw" +
		"AAA9A5q0sEOatLBAAAAAdzc2gtcnNhAAABAQChpt8fvcJiMfq/lDDaVQIhZSR+b33sYLCV01" +
		"RA4Ohyp2xw+H/unUiLvqzwT9XGQ8NXU9VJuBUnuASMqdZB8rHcFwMDDQh00tmtnuLnUlG6EH" +
		"VZZEH7aT/o0l400XBwhuUIhkRQqfXpD3RLGS/9b1aIVrMUJFEQetfA4Le8ycTSYDTS87gRTH" +
		"r2T71T3a60H2l6cxHWjLKz6iAsymMopywiuLzw/kNBKM/FS1nays+0n1/jSVdlW45U3qWdDa" +
		"6vxgG3R1k9yWAy64fy0f2cFeadndsh9mmw4+gJMj/vOG6OjhB9Zf5yv2irvR79z8wR/SVrWa" +
		"BRetmzpGUTTg1S+YXPAAAAAwEAAQAAAQBIsJ8YyN3GBi95QowNQbiph6+3Yy8+wePmG5eBbS" +
		"FZnUu4KSZuCC/9FwrxLRU2CHaoqDv64Foy+B6jmiOMmDO+gmKKVqv8zGuyFQZ3ep7hilEal2" +
		"jMQvHIIgDWw29KSn5nLOk/VI4N9TMYKgVYc154BHSWlBYX0QV/6Az3ScvFxWJDvNsjrxhX8C" +
		"BmzleZDuUWeqBYUtUXMesSkTWYfTBdkYpftpMhwjuKA8L+O7ITd85S5Dgk948bN+MA977IQb" +
		"kk+41PMbB5I/QzUj8eje28K/4lwdngNtptGayDhM31VANRikqwlF2LXbckW8PVqXko6JCGn9" +
		"SQgh4e541U3RExAAAAgDX2Xv9PbyyWUro+awyXWyBpGQMd5KK6wHUMMVyFIPpeR5dtkuPDnl" +
		"4XaujFn2Lq4D8bm1CEhpGpIFazV5zc6Q9+y/ihuN0/mvH7+FES1Pz2ijPmoqYUXY9Pe8/l+o" +
		"9QcD87dAvo67ZQ98va2P8Hzh1dBkg9nH5y7uy6XMrumr5wAAAAgQDOBYGasyxj3+yTW/aH1n" +
		"gWblX9Qy/mSqDeZ/GTpuhXpTHbrzM80RcJ5Vkt+ns95PhfmL0KWvoZeRWm1J3HSdhVLdj4Dd" +
		"ZP+iMuXtAaRADZcpdKVeoko8lDhEGgWj1zpdg4rG8SfcVXN3xDHfFtupSpBGztEyx/CPhJqX" +
		"+zA6J9VQAAAIEAyN3jihHRhsptB3lCoz+nYBDKIbOqyy1xVeIi6FYYnACPM84ywW1LkEreg1" +
		"oysSScOzdlysNTQvjgli+76mLrG54VE6jpxXkuzslo8cZKj0K+M6jMAgpnowr+Z2uM3utDtd" +
		"HsaJ9Q2SYLON/uJF7m9WfmYY9Md617G6o/YiRNVpMAAAAXc2NhbGR3ZWxsQHg2ODQ4NjcubG" +
		"9jYWwBAgME"
	sshKey8192 = "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" +
		"g2ODQ4NjcubG9jYWwBAgME"
)

func TestGetShannons(t *testing.T) {
	var tests = []struct {
		input string
		score int
	}{
		{
			"",
			0,
		},
		{
			"1111111111111111111111",
			0,
		},
		{
			"1223334444",
			185,
		},
		{
			"tr,BfXBRTjPBn7xUF@csyCpCNpeGpqhzHbs9FYmGUGsmiCCEPTPgue>Bie2ftpXk",
			588,
		},
		{
			"AoXWDwkkddxEhttQtJRhP>$UXHfMZ?7GMzrjkW8DUJZV4qoHMCinozsEcsgaWjgy",
			523,
		},
		{
			sshKey2048,
			600,
		},
		{
			sshKey8192,
			600,
		},
	}
	for i, test := range tests {
		if entropy.GetShannons(test.input) > test.score {
			t.Fatalf("Test %d failed with score %d (expected %d)", i, test.score, entropy.GetShannons(test.input))
		}
	}
}
