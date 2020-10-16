package main

import (
	"testing"
	"time"
)
func TestDnsEnum(t *testing.T){
	go func(){
		<-time.After(time.Minute * 15)
		t.Fatal("Main() Test: Terminating after timeout (15min)")
	}()
}