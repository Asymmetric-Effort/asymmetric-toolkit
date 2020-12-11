package dictionary

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"bytes"
	"encoding/binary"
	"fmt"
)

func (o *File) WordWrite(offset int64, word *Word) (err error) {
	//
	// Write the word struct into a bytes buffer then write the buffer
	// to disk.
	//
	var buf bytes.Buffer
	var header uint32 = 0
	//
	// create the header bitmap.
	//
	parentPtrSz := getPointerSize(word.Parent)
	lhsPtrSz := getPointerSize(word.Lhs)
	rhsPtrSz := getPointerSize(word.Rhs)
	header = (parentPtrSz << 16) | (lhsPtrSz << 8) | rhsPtrSz
	//
	// Write header
	//
	fmt.Printf("\tfile::WordWrite() - header:%b\n", header)
	if err= binary.Write(&buf, binary.BigEndian, header); err != nil {
		return fmt.Errorf("\tError  writing header:%v", err)
	}
	//
	// Write wordsz
	//
	if err = writeWord16Buf(&buf, len(word.Word)); err != nil {
		return err
	}
	//
	// Write Parent address
	//
	fmt.Printf("\tword.Parent (%x) of size %d bytes\n",word.Parent,parentPtrSz)
	if err = writeVarLenAddressToBuf(&buf, word.Parent, parentPtrSz); err != nil {
		return fmt.Errorf("\tError  writing Parent (%x):%v", word.Parent, err)
	}
	//
	// Write LHS address
	//
	fmt.Printf("\tword.Lhs (%x) of size %d bytes\n",word.Lhs,lhsPtrSz)
	if err = writeVarLenAddressToBuf(&buf, word.Lhs, lhsPtrSz); err != nil {
		return fmt.Errorf("\tError  writing Lhs (%x):%v", word.Parent, err)
	}
	//
	// Write RHS Address
	//
	fmt.Printf("\tword.Rhs (%x) of size %d bytes\n",word.Rhs,rhsPtrSz)
	if err = writeVarLenAddressToBuf(&buf, word.Rhs, rhsPtrSz); err != nil {
		return fmt.Errorf("\tError  writing Rhs (%x):%v", word.Parent, err)
	}
	//
	// Write Word []byte array
	//
	if err = binary.Write(&buf, binary.BigEndian, word.GetWordBytes()); err != nil {
		return fmt.Errorf("Error  writing Word (%x):%v", word.GetWordBytes(), err)
	}
	fmt.Printf("\tFile::WriteWord(): Writing Buffer to disk at %x\n", offset)
	//
	// Write the buffer to disk.
	//
	_, o.Err = o.Handle.WriteAt(buf.Bytes(), offset)
	errors.Assert(o.Err == nil, fmt.Sprintf("File::WriteWord(): post-write error: %v", o.Err))
	//
	// Done
	//
	fmt.Printf("\tFile::WriteWord(): Write successful.\n")
	return o.Err
}
