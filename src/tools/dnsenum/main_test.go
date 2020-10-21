package main_test

import (
	"testing"
	"time"
)
func TestDnsEnum(t *testing.T){
	go func(){
		<-time.After(time.Minute * 30)
		t.Fatal("Main() Test: Terminating after timeout (30min)")
	}()
}