package main_test

import (
	"fmt"
	"os"
	"testing"
	"time"
)
func TestDnsEnum(t *testing.T){
	go func(){
		<-time.After(time.Minute * 30)
		fmt.Println("Main() Test: Terminating after timeout (30min)")
		os.Exit(1)
	}()
}