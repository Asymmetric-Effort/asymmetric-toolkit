package persistentTree

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/file"
	"encoding/binary"
	"os"
	"testing"
)

func TestTree_GotoPosition(t *testing.T) {
	const testFileName string = "goto_position_testfile.txt"
	var err error
	func() {
		t.Log("Generate test file")
		var fileHandle *os.File
		fileHandle, err = os.Create(testFileName)
		if err != nil {
			t.Errorf("Error creating a new file.  Error: %v", err)
			t.FailNow()
		}
		t.Log("Test file created (no error)")
		defer func() {
			if fileHandle == nil {
				t.Log("File not open.  Cannot close")
				return
			}
			err := fileHandle.Close()
			if err != nil {
				t.Errorf("Error closing file: %v", err)
			}
			t.Log("File closed: " + testFileName)
		}()
		t.Log("Creating test file data")
		var b = make([]byte, 1)

		t.Log("Create 10 bytes")
		for n := 0; n < 10; n++ {
			b[0] = byte(n + 61)
			_, err := fileHandle.Write(b)
			if err != nil {
				t.Errorf("Error writing byte to disk (%x)", n)
				t.FailNow()
			}
		}
		t.Log("Test data written.")
	}()

	func() {
		t.Log("Running test")
		var tree Tree
		tree.file, err = os.Open(testFileName)
		if err != nil {
			t.Errorf("Error opening tree file:%v", err)
		}
		t.Log("Tree file is open")
		defer func() {
			if tree.file == nil {
				t.Log("tree file not open.  Cannot close")
				return
			}
			err := tree.file.Close()
			if err != nil {
				t.Errorf("Error closing tree file: %v", err)
			}
			t.Log("tree file closed: " + testFileName)
		}()

		func() {
			t.Log("Test #1: hit all the possible addresses without error")
			for n := 0; n < 10; n++ {
				t.Logf("Seek Position:%d", n)
				tree.gotoPosition(int64(n))
				var b byte
				err = binary.Read(tree.file, binary.BigEndian, &b)
				if err != nil {
					t.Errorf("Error reading byte(%d): %v", n, err)
				}

				t.Logf("\tseek %d and get %d", n, b)
				if b != byte(n+61) {
					t.Fatalf("Error in expected data.  Expected %d.  Found %d", n+61, b)
					t.FailNow()
				}
			}
		}()
	}()
	defer func() {
		if file.FileExists(testFileName) {
			err = os.Remove(testFileName)
			if err != nil {
				t.Errorf("Error closing tree file: %v", err)
			}
			t.Log("tree file removed: " + testFileName)
		}
	}()

}
