package persistentTreeFlags

import (
	"testing"
)

func TestFlags_NoCompression_NoEncryption(t *testing.T) {
	var flag Flags
	flag.raw = 0
	//
	// Compression: Clear | Encryption : Clear
	//
	if flag.Get(Compression) {
		t.Error("Compression flag expected to be clear")
		t.FailNow()
	}
	if flag.Get(Encryption) {
		t.Error("Compression flag expected to be clear")
		t.FailNow()
	}
}
func TestFlags_ClearCompression_NoEncryption(t *testing.T) {
	var flag Flags
	flag.raw = 0
	//
	// Compression: Clear | Encryption : Clear
	//
	flag.Clear(Compression)
	if flag.Get(Compression) {
		t.Error("Compression flag expected to be clear")
		t.FailNow()
	}
	if flag.Get(Encryption) {
		t.Error("Encryption should have been clear")
		t.FailNow()
	}
}
func TestFlags_ClearCompression_ClearEncryption(t *testing.T) {
	var flag Flags
	flag.raw = 65535
	//
	// Compression: Clear | Encryption : Clear
	//
	flag.Clear(Compression)
	if flag.Get(Compression) {
		t.Error("Compression flag expected to be clear")
		t.FailNow()
	}
	flag.Clear(Encryption)
	if flag.Get(Encryption) {
		t.Error("Encryption should have been clear")
		t.FailNow()
	}

}
func TestFlags_Compression_Encryption(t *testing.T) {
	var flag Flags
	flag.raw = 0
	//
	// Compression: Set | Encryption : Clear
	//
	t.Logf("(1)flag.raw:%d",flag.raw)
	t.Logf("(2)flag.raw:%d",flag.raw)
	flag.Set(Compression)
	t.Logf("(3)flag.raw:%d",flag.raw)
	if !flag.Get(Compression) { //Expect Compression
		t.Logf("(4)flag.raw:%d",flag.raw)
		t.Error("Compression flag expected to be set")
		t.FailNow()
	}
	t.Logf("(5)flag.raw:%d",flag.raw)
	if flag.Get(Encryption) { //Expect no Encryption
		t.Error("Encryption flag expected to be clear")
		t.FailNow()
	}
}
func TestFlags_NoCompression_Encryption(t *testing.T) {
	var flag Flags
	flag.raw = 0
	//
	// Compression: Clear | Encryption : Set
	//
	flag.Set(Encryption)
	if flag.Get(Compression) {
		t.Error("Compression flag expected to be clear")
		t.FailNow()
	}
	if !flag.Get(Encryption) {
		t.Log("Encryption flag is set")
		t.FailNow()
	}
}
func TestFlags_SetCompression_SetEncryption(t *testing.T) {
	var flag Flags
	flag.raw = 0
	//
	// Compression: Set | Encryption : Set
	//
	t.Log("Set compression flag (1)...")

	flag.Set(Compression)

	t.Log("-->(1): Get(Compression):",flag.Get(Compression))
	t.Log("-->(1): Get(Encryption):",flag.Get(Encryption))

	if !flag.Get(Compression) {
		t.Error("Compression flag expected to be set(1)")
		t.FailNow()
	}

	t.Log("Set encryption flag (2)...")

	flag.Set(Encryption)

	t.Log("-->(2): Get(Compression):",flag.Get(Compression))
	t.Log("-->(2): Get(Encryption):",flag.Get(Encryption))

	if !flag.Get(Compression) {
		t.Error("Compression flag expected to be set(2)")
		t.FailNow()
	}

	if !flag.Get(Encryption) {
		t.Error("Encryption flag expected to be set(2)")
		t.FailNow()
	}
}
