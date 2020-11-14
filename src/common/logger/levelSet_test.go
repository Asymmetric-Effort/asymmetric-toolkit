package logger

import (
	"math"
	"testing"
)
func TestLogLevel_Set(t *testing.T){
	recoverOnError:=func(){recover()}
	for i:=math.MinInt8;i<math.MaxInt8;i++ {
		func() {
			pre:=func(){}
			post:=func(){}
			if i < 0 || i >4 {
				pre=recoverOnError
				post=func(){panic("expected error")}

			}
			defer pre()
			var n=Level(i)
			var L = Level(i)
			L.Set(n)
			post()
		}()
	}
}
