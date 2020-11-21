package tags

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"bytes"
	"fmt"
	"testing"
)

func TestTag_SerializeBoolean1(t *testing.T) {
	for i := 0; i < 100000; i++ {

		o := NewBoolean()

		errors.Assert(o.Count() == 0, "expect 0")

		o.Add("test", true)

		errors.Assert(o.Count() == 1, "expect 1")

		tag := o.SerializeBoolean()
		expectedTag := []byte{1, 1, 4, 1, 116, 101, 115, 116}

		if !bytes.Equal(tag, expectedTag) {
			fmt.Println("")
			fmt.Println("Test #", i)
			fmt.Println("\ttag(e):", expectedTag)
			fmt.Println("\ttag(a):", tag)
			fmt.Printf("\tKey-value mismatch. v:%s\n", string(tag))
		}
	}
}

func TestTag_SerializeBoolean2(t *testing.T) {
	for i := 0; i < 100000; i++ {
		o := NewBoolean()

		fmt.Println("Test 2(", i, ")")
		fmt.Println("")

		o.Add("test1", true)
		o.Add("test2", true)

		errors.Assert(o.Count() == 2, "expect 2 on count")
		errors.Assert(o.Find("test1"), "expect true on find")
		errors.Assert(o.Delete("test2"), "expect true on delete")
		errors.Assert(o.Count() == 1, "expect 1 on count")
		o.Add("test3", true)
		errors.Assert(o.Count() == 2, "expect 2 on count")
	}
}

func TestTag_SerializeBoolean3(t *testing.T) {
	for n := 0; n < 10; n++ {
		fmt.Println("Test 3(", n, ")")
		fmt.Println("")
		o := NewBoolean()

		o.Add("test1", false)
		o.Add("test2", false)
		o.Add("test3", false)

		fmt.Print("        :")
		for i := 0; i < 40; i++ {
			fmt.Print(". ")
		}
		fmt.Println("")

		tag := o.SerializeBoolean()
		//
		// tagCount = 3
		// tag #1:
		// 		type: 			BooleanTag (1)
		// 		length(key): 	5 (chars)
		// 		value: 			0 (false)
		//		value([]byte):  {116(t), 101(e), 115(s), 116(t),49(1)}
		// tag #2:
		// 		type: 			BooleanTag (1)
		// 		length(key): 	5 (chars)
		// 		value: 			0 (false)
		//		value([]byte):  {116(t), 101(e), 115(s), 116(t),50(2)}
		// tag #3:
		// 		type: 			BooleanTag (1)
		// 		length(key): 	5 (chars)
		// 		value: 			0 (false)
		//		value([]byte):  {116(t), 101(e), 115(s), 116(t),51(3)
		//
		expectedTag := []byte{3, 1, 5, 0, 116, 101, 115, 116, 49, 1, 5, 0, 116, 101, 115, 116, 50, 1, 5, 0, 116, 101, 115, 116, 51}

		fmt.Println("tag(e):", expectedTag)
		fmt.Println("tag(a):", tag)

		errors.Assert(bytes.Equal(tag, expectedTag), fmt.Sprintf("Key-value mismatch. v:%s", string(tag)))
	}
}
