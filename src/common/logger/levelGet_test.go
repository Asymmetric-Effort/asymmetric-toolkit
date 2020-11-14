package logger

import (
	"math"
	"testing"
)
func TestLogLevel_Get(t *testing.T){
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
			var L = Level(i)
			_ = L.Get()
			post()
		}()
	}
}
