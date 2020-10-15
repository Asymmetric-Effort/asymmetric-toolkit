package main

import (
	"fmt"
)

const (
	dnsChars     = "0123456789"
	maxSizeLimit = 6
)

func nextSequence(n int) func() string {
	r := []rune(dnsChars)
	p := make([]rune, n)
	x := make([]int, len(p))
	rLen:=len(r)
	return func() string {
		p := p[:len(x)]
		for i, v := range x {
			p[i] = r[v]
		}
		for i := len(x) - 1; i >= 0; i-- {
			x[i]++
			if x[i] < rLen {
				break
			}
			x[i] = 0
			if i <= 0 {
				x = x[0:0]
				break
			}
		}
		return string(p)
	}
}

func bruteForce() {
	done:=make(chan bool,1)
	out:=make(chan string,16384)
	go func(){
		for {
			fmt.Println(<-out)
		}
		<-done
	}()
	for n := 0; n <= maxSizeLimit; n++ {
		np := nextSequence(n)
		for pwd := np(); len(pwd) != 0; pwd = np() {
			out <-pwd
		}
	}
	done<-true
}

func main() {
	bruteForce()
}
