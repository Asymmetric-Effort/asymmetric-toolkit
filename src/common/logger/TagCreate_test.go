package logger

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"fmt"
	"testing"
)

func TestLogger_TagCreate(t *testing.T) {
	var o TagTracker
	tests := []struct {
		tagName string
		create  bool
		pass    bool
		result  string
	}{
		{
			tagName: "test1", // OK
			create:  true,
			pass:    true,
			result:  "OK",
		}, {
			tagName: "test2", // OK
			create:  true,
			pass:    true,
			result:  "OK",
		}, {
			tagName: "test3", // OK
			create:  true,
			pass:    true,
			result:  "OK",
		}, {
			tagName: "test4", // OK
			create:  true,
			pass:    true,
			result:  "OK",
		}, {
			tagName: "test5", // OK
			create:  true,
			pass:    true,
			result:  "OK",
		}, {
			tagName: "test6", // OK
			create:  true,
			pass:    true,
			result:  "OK",
		}, {
			tagName: "test7", // OK
			create:  true,
			pass:    true,
			result:  "OK",
		}, {
			tagName: "test 7", // BAD TAG NAME
			create:  true,
			pass:    false,
			result:  "Bad Tag Name (space)",
		}, {
			tagName: "testbadspace01234012345678901234567898567890132365498765432132165465464", // BAD TAG NAME
			create:  true,
			pass:    false,
			result:  "Bad Tag Name (too long)",
		}, {
			tagName: "test-7", // BAD TAG NAME
			create:  true,
			pass:    false,
			result:  "Bad Tag Name (bad char-hyphen)",
		}, {
			tagName: "test_7", // BAD TAG NAME
			create:  true,
			pass:    false,
			result:  "Bad Tag Name (bad char-underscore)",
		}, {
			tagName: "test+7", // BAD TAG NAME
			create:  true,
			pass:    false,
			result:  "Bad Tag Name (bad char-plus sign)",
		}, {
			tagName: "test.7", // BAD TAG NAME
			create:  true,
			pass:    false,
			result:  "Bad Tag Name (bad char-period)",
		}, {
			tagName: "", // BAD TAG NAME
			create:  true,
			pass:    false,
			result:  "Bad Tag Name (too short)",
		}, {
			tagName: "test8", // OK
			create:  true,
			pass:    false,
			result:  "",
		}, {
			tagName: "test8", // DUPLICATE
			create:  true,
			pass:    false,
			result:  "",
		}, {
			tagName: "test9", // NOT CREATED
			create:  false,
			pass:    false,
			result:  "",
		}, {
			tagName: "test10", // OK
			create:  true,
			pass:    true,
			result:  "OK",
		},
	}

	recoverOnError := func() {
		fmt.Print("\trecoverOnError()")
		recover()
	}

	panicOnError := func() {
	}

	f := panicOnError

	testFunc := func(i int, tagName string, create bool, pass bool, result string) {
		if (i%5 == 0) || !pass {
			fmt.Printf("\ni:%2d tagName:%-8.8s create:%6t, pass:%6t, result:%-50.50s",
				i, tagName, create, pass, result)
		}
		if !create {
			f = recoverOnError
		}
		defer f()
		tagId := o.Create(tagName)
		if !pass {
			panic(fmt.Sprintf("Expected Error (pass):  %s", result))
		}
		if !create {
			panic(fmt.Sprintf("Expected error (create):%s", result))
		}
		if create {
			errors.Assert(tagId >= 0, "expect non-zero tagId")
		}
	}

	for i, test := range tests {
		if test.pass {
			f = panicOnError
		} else {
			f = recoverOnError
		}
		testFunc(i, test.tagName, test.create, test.pass, test.result)
	}
	for i := len(tests); i < maxTagTrackerDictionarySize+10; i++ {
		tagName := fmt.Sprintf("LimitTest%d", i)
		pass := true
		result := "OK"

		if i < maxTagTrackerDictionarySize {
			f = panicOnError
			pass = true
			result = "over limit"
		} else {
			f = recoverOnError
		}
		testFunc(i, tagName, true, pass, result)
	}
	fmt.Println("")
}
