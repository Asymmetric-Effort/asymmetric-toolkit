package logger

import (
	assert "github.com/sam-caldwell/adrestia-assertions/src"
	"testing"
)

func TestLogger_TagTrackerCreate(t *testing.T) {
	func() {
		var o TagTracker
		tagId := o.Create("test0")
		assert.Error(tagId == 0, "expected TagId 0")
		tagId = o.Create("test1")
		assert.Error(tagId == 1, "expected TagId 1")
	}()
	func() {
		var o TagTracker
		defer func() { recover() }()
		_ = o.Create("test 1")
		panic("Expected error on bad tag.")
	}()
	func() {
		var o TagTracker
		defer func() { recover() }()
		_ = o.Create("testbadspace01234012345678901234567898567890132365498765432132165465464")
		panic("Expected error on bad tag.")
	}()
	func() {
		var o TagTracker
		defer func() { recover() }()
		_ = o.Create("test-1")
		panic("Expected error on bad tag.")
	}()
	func() {
		var o TagTracker
		defer func() { recover() }()
		_ = o.Create("")
		panic("Expected error on bad tag.")
	}()
	func() {
		var o TagTracker
		defer func() { recover() }()
		_ = o.Create("test_1")
		panic("Expected error on bad tag.")
	}()
	func() {
		var o TagTracker
		defer func() { recover() }()
		_ = o.Create("test+1")
		panic("Expected error on bad tag.")
	}()
	func() {
		var o TagTracker
		defer func() { recover() }()
		_ = o.Create("_")
		panic("Expected error on bad tag.")
	}()
	func() {
		var o TagTracker
		defer func() { recover() }()
		_ = o.Create("test!")
		panic("Expected error on bad tag.")
	}()
	func() {
		var o TagTracker
		defer func() { recover() }()
		_ = o.Create("test@")
		panic("Expected error on bad tag.")
	}()
	func() {
		var o TagTracker
		defer func() { recover() }()
		_ = o.Create("test#")
		panic("Expected error on bad tag.")
	}()
	func() {
		var o TagTracker
		defer func() { recover() }()
		_ = o.Create("test$")
		panic("Expected error on bad tag.")
	}()
	func() {
		var o TagTracker
		defer func() { recover() }()
		_ = o.Create("test%")
		panic("Expected error on bad tag.")
	}()
	func() {
		var o TagTracker
		defer func() { recover() }()
		_ = o.Create("test^")
		panic("Expected error on bad tag.")
	}()
	func() {
		var o TagTracker
		defer func() { recover() }()
		_ = o.Create("test&")
		panic("Expected error on bad tag.")
	}()
	func() {
		var o TagTracker
		defer func() { recover() }()
		_ = o.Create("test*")
		panic("Expected error on bad tag.")
	}()
	func() {
		var o TagTracker
		defer func() { recover() }()
		_ = o.Create("test$")
		panic("Expected error on bad tag.")
	}()
	func() {
		var o TagTracker
		defer func() { recover() }()
		_ = o.Create("test(")
		panic("Expected error on bad tag.")
	}()
	func() {
		var o TagTracker
		defer func() { recover() }()
		_ = o.Create("test)")
		panic("Expected error on bad tag.")
	}()
	func() {
		var o TagTracker
		defer func() { recover() }()
		_ = o.Create("test${")
		panic("Expected error on bad tag.")
	}()
	func() {
		var o TagTracker
		defer func() { recover() }()
		_ = o.Create("test}")
		panic("Expected error on bad tag.")
	}()
	func() {
		var o TagTracker
		defer func() { recover() }()
		_ = o.Create("test[")
		panic("Expected error on bad tag.")
	}()
	func() {
		var o TagTracker
		defer func() { recover() }()
		_ = o.Create("test]")
		panic("Expected error on bad tag.")
	}()
	func() {
		var o TagTracker
		defer func() { recover() }()
		_ = o.Create("test:")
		panic("Expected error on bad tag.")
	}()
	func() {
		var o TagTracker
		defer func() { recover() }()
		_ = o.Create("test;")
		panic("Expected error on bad tag.")
	}()
	func() {
		var o TagTracker
		defer func() { recover() }()
		_ = o.Create("test'")
		panic("Expected error on bad tag.")
	}()
	func() {
		var o TagTracker
		defer func() { recover() }()
		_ = o.Create("test/")
		panic("Expected error on bad tag.")
	}()
	func() {
		var o TagTracker
		defer func() { recover() }()
		_ = o.Create("test?")
		panic("Expected error on bad tag.")
	}()
	func() {
		var o TagTracker
		defer func() { recover() }()
		_ = o.Create("test>")
		panic("Expected error on bad tag.")
	}()
	func() {
		var o TagTracker
		defer func() { recover() }()
		_ = o.Create("test<")
		panic("Expected error on bad tag.")
	}()
	func() {
		var o TagTracker
		defer func() { recover() }()
		_ = o.Create("test.")
		panic("Expected error on bad tag.")
	}()
	func() {
		var o TagTracker
		defer func() { recover() }()
		_ = o.Create("test,")
		panic("Expected error on bad tag.")
	}()
	func() {
		var o TagTracker
		defer func() { recover() }()
		_ = o.Create("test|")
		panic("Expected error on bad tag.")
	}()
	func() {
		var o TagTracker
		defer func() { recover() }()
		_ = o.Create("test\\")
		panic("Expected error on bad tag.")
	}()
	func() {
		var o TagTracker
		defer func() { recover() }()
		_ = o.Create("test=")
		panic("Expected error on bad tag.")
	}()
	func() {
		var o TagTracker
		defer func() { recover() }()
		_ = o.Create("test~")
		panic("Expected error on bad tag.")
	}()
	func() {
		var o TagTracker
		defer func() { recover() }()
		_ = o.Create("test`")
		panic("Expected error on bad tag.")
	}()
}
