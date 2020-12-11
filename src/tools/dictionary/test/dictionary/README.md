Dictionary
==========


### General format
```
====================================================
	00 00 00 00       - 32-bit header. See above.
   	00 00             - 16-bit word size.
	00 00 ... 00 00   - ptr to parent node (determined by parentPtrSz)
   	00 00 ... 00 00   - ptr to lhs node (determined by lhsPtrSz)
	00 00   		  - ptr to rhs node (determined by rhsPtrSz)
	00 00 ... 00 00   - arbitrary byte string (word)
====================================================
```

### 	Word Header
```
====================================================
bit    definition
====================================================
0x00 - Reserved
0x01 - Reserved
0x02 - Reserved
0x03 - Reserved
0x04 - Reserved
0x05 - Reserved
0x06 - Reserved
0x07 - Reserved

0x08 - Parent Pointer Size (parentPtrSz) Byte0 msb
0x09 - Parent Pointer Size (parentPtrSz) Byte1
0x0A - Parent Pointer Size (parentPtrSz) Byte2
0x0B - Parent Pointer Size (parentPtrSz) Byte3
0x0C - Parent Pointer Size (parentPtrSz) Byte4
0x0D - Parent Pointer Size (parentPtrSz) Byte5
0x0E - Parent Pointer Size (parentPtrSz) Byte6
0x0F - Parent Pointer Size (parentPtrSz) Byte7 lsb

0x10 - LHS Pointer Size (lhsPtrSz) Byte0 msb
0x11 - LHS Pointer Size (lhsPtrSz) Byte1
0x12 - LHS Pointer Size (lhsPtrSz) Byte2
0x13 - LHS Pointer Size (lhsPtrSz) Byte3
0x14 - LHS Pointer Size (lhsPtrSz) Byte4
0x15 - LHS Pointer Size (lhsPtrSz) Byte5
0x16 - LHS Pointer Size (lhsPtrSz) Byte6
0x17 - LHS Pointer Size (lhsPtrSz) Byte7 lsb

0x18 - RHS Pointer Size (rhsPtrSz) Byte0 msb
0x19 - RHS Pointer Size (rhsPtrSz) Byte1
0x1A - RHS Pointer Size (rhsPtrSz) Byte2
0x1B - RHS Pointer Size (rhsPtrSz) Byte3
0x1C - RHS Pointer Size (rhsPtrSz) Byte4
0x1D - RHS Pointer Size (rhsPtrSz) Byte5
0x1E - RHS Pointer Size (rhsPtrSz) Byte6
0x1F - RHS Pointer Size (rhsPtrSz) Byte7 lsb
====================================================
```
```
10000000100000001
10000000100000001
MASK     11111111 
         --------
         00000001
 



```
