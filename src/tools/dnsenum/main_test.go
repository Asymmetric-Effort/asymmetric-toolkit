package main

import (
	"testing"
	"time"
)
func TestMain(m *testing.M) {
	go func(){
		<-time.After(time.Minute * 30)
		m.Run()
	}()
}