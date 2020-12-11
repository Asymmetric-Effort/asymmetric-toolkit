package dictionary

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"encoding/binary"
	"fmt"
	"os"
)

const wordSizeLength = 2 //Two bytes (16 bits) 0...65535

func (o *File) WordRead(offset int64) *Word {
	var err error
	var pos int64
	var wordSz int64
	var header uint32
	var word = Word{}
	var rhsPtrSz uint8
	var lhsPtrSz uint8
	var parentPtrSz uint8
	//
	// Seek record location
	//
	fmt.Printf("\tFile::WordRead()  offset:%x\n", offset)
	pos, o.Err = o.Handle.Seek(offset, 0)
	errors.Assert(o.Err == nil, fmt.Sprintf("File::WordRead(): post-seek error: %v", o.Err))
	errors.Assert(pos == offset, fmt.Sprintf("Expected to have moved to target position:%x but at %x", offset, pos))
	fmt.Printf("\tFile::WordRead() curr position:%x\n", pos)
	//
	//
	//
	if o.Err = binary.Read(o.Handle, binary.BigEndian, &header); o.Err != nil {
		panic(fmt.Sprintf("\tFile::WordRead(): Error Reading file "+
			"(header:%x), error: %v\n", header, o.Err))
	}
	fmt.Printf("\tFile::WordRead(): OK (header:%x)\n", header)
	//
	//	Header is read.  Parsing it into the values...
	//
	parentPtrSz = uint8((header & 65280) >> 16)
	lhsPtrSz = uint8((header & 255) >> 8)
	rhsPtrSz = uint8(header & 255)
	fmt.Printf("\tFile::WordRead(): "+
		"parentPtrSz:%x "+
		"lhsPtrSz:%x "+
		"rhsPtrSz:%x\n", parentPtrSz, lhsPtrSz, rhsPtrSz)
	//
	// Reading the word size for the new field.
	//
	wordSz, err = readField(o.Handle, 2)
	if err != nil {
		panic(fmt.Sprintf("\tFile::WordRead() Error reading wordSz: %v", o.Err))
	}
	fmt.Printf("\tFile::WordRead(): wordSz:%d [%x]\n", wordSz, wordSz)
	//
	//
	//
	word.Parent, err = readField(o.Handle, parentPtrSz)
	if err != nil {
		panic(fmt.Sprintf("\tFile::WordRead() Error reading Parent: %v", o.Err))
	}
	word.Lhs, err = readField(o.Handle, lhsPtrSz)
	if err != nil {
		panic(fmt.Sprintf("\tFile::WordRead() Error reading Lhs: %v", o.Err))
	}
	word.Rhs, err = readField(o.Handle, rhsPtrSz)
	if err != nil {
		panic(fmt.Sprintf("\tFile::WordRead() Error reading Rhs: %v", o.Err))
	}
	word.Word = func() string {
		buf := make([]byte, wordSz)
		if o.Err = binary.Read(o.Handle, binary.BigEndian, &buf); o.Err != nil {
			panic(fmt.Sprintf("\tFile::WordRead() Error reading word: %v", o.Err))
		}
		return string(buf)
	}()
	if o.Err = binary.Read(o.Handle, binary.BigEndian, &word.Lhs); o.Err != nil {
		panic(o.Err)
	}
	if o.Err = binary.Read(o.Handle, binary.BigEndian, &word.Rhs); o.Err != nil {
		panic(o.Err)
	}
	errors.Assert(o.Err == nil, fmt.Sprintf("\tFile::WordRead(): post-read error: %v", o.Err))
	return &word
}

func readField(f *os.File, sz uint8) (result int64, err error) {
	//
	// Read the word file with a given size.
	//
	fmt.Printf("\tFile::readField(): starting (sz:%d)\n", sz)
	buf := make([]byte, sz)
	if err = binary.Read(f, binary.BigEndian, &buf); err != nil {
		panic(fmt.Sprintf("\t\tFile::readField() Error reading word: %v", err))
	}
	switch sz {
	case 1: // Byte
		result = int64(buf[0])

	case 2: // Word16
		result = int64(binary.BigEndian.Uint16(buf))
	case 4: // Word32
		result = int64(binary.BigEndian.Uint32(buf))
	case 8: // Word64
		result = int64(binary.BigEndian.Uint64(buf))
	default:
		panic(fmt.Sprintf("\t\t\tUnsupported field width in readField():%d", sz))
	}
	fmt.Printf("\tFile::readField(): terminate (result:%x)\n", result)
	return result, nil
}
