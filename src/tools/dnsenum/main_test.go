package main

import (
	"fmt"
	"testing"
	"time"
)

func TestDnsEnum(t *testing.T) {
	go func() {
		<-time.After(time.Minute * 30)
		t.Fatal("Main() Test: Terminating after timeout (30min)")
	}()
}

func TestProcessSpecification(t *testing.T) {
	func() {
		args := []string{"--help"}
		config, exitProgram, err := ProcessSpecification(args)
		if err != nil {
			t.Errorf("[main_test]: TestProcessSpecifcation(): %v", err)
		}
		if exitProgram {
			fmt.Println("Exit Program as expected")
			return
		}
		if config == nil {
			t.Error("Internal error nil config encountered.")
		}
	}()

	func() {
		fmt.Println("---Process Specification -h")
		args := []string{"-h"}
		config, exitProgram, err := ProcessSpecification(args)
		if err != nil {
			t.Error(err)
		}
		if exitProgram {
			fmt.Println("Exit Program as expected")
			return
		}
		if config == nil {
			t.Error("Internal error nil config encountered.")
		}
	}()

	func() {
		fmt.Println("---Process Specification --version")
		args := []string{"--version"}
		config, exitProgram, err := ProcessSpecification(args)
		if err != nil {
			t.Error(err)
		}
		if exitProgram {
			fmt.Println("Exit Program as expected")
			return
		}
		if config == nil {
			t.Error("Internal error nil config encountered.")
		}
	}()
}

func TestMain(m *testing.M){
	m.Run()
}