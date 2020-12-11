package dictionary

/*
	dictionary.File::Add() will add a node to the Word tree.
*/

import (
	"fmt"
)

func (o *File) Add(wordStr *string) {

	o.mutex.Lock()
	defer o.mutex.Unlock()
	//
	//Create word
	//
	var newWord Word
	newWord.Set(wordStr, 0, 0)
	//
	//Read current root word and get the word string.
	//
	rootWord := func() *string {
		var offset = o.Header.RootWord
		if offset == 0 {
			fmt.Println("\tFile::Add() DEBUG: we are at root node.")
			return nil //There is no root word.  We are at root.
		}
		fmt.Printf("\tFile::Add() DEBUG: reading node at %x\n", offset)
		return &o.WordRead(offset).Word
	}()

	if rootWord == nil {
		fmt.Println("\tFile::Add() DEBUG: tree is at root (", o.Header.RootWord, ")")
	} else {
		if *rootWord < newWord.Word {
			fmt.Printf("\tFile::Add() DEBUG: tree is lhs (%x)\n", o.Header.RootWord)
			newWord.Lhs = o.Header.RootWord
		}
		if *rootWord > newWord.Word {
			fmt.Printf("\tFile::Add() DEBUG: tree is rhs (%x)\n", o.Header.RootWord)
			newWord.Rhs = o.Header.RootWord
		}
		if *rootWord == newWord.Word {
			fmt.Printf("\tFile::Add() Warning: Duplicate word found. Word %s", *wordStr)
			return
		}
	}

	//Move to EOF and get position (this is our new node position).  It is now root of our tree.
	originalState := o.Header.RootWord // Allow rollback.
	fmt.Printf("\tFile::Add() original state:%x, rootWord:%x\n", originalState, o.Header.RootWord)
	o.Header.RootWord = o.MoveEof()
	fmt.Printf("\tFile::Add() original state:%x, rootWord:%x\n", originalState, o.Header.RootWord)

	//Write word to EOF
	if err := o.WordWrite(o.Header.RootWord, &newWord); err == nil {
		fmt.Println("\tFile::Add(): no error writing word do disk.")
	} else {
		fmt.Println("\trollback to original state.  err:", err)
		o.Header.RootWord = originalState //Revert our header state to keep the tree as it was.
		//
		//ToDo: We need to reclaim space allocated in the file.
		//      Note we may not have written anything. We may also
		//      have an orphaned node in the tree.
		//
	}

	fmt.Println("\tFile::Add() Updating file header")
	o.WriteHeader(o.Header)
	fmt.Println("\tFile::Add() header update successful")
}
