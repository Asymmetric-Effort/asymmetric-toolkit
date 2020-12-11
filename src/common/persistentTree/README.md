persistentTree
==============

## Purpose
    * Establish a reusable persistent binary tree file format optimized for a balance between disk space
      used and I/O operations.

## Usage
TBD

## Theory of Operation
```                          
                                    (root address)
                                           | 
                                           | 
                                           |  (0)                
                                           |   ^                 
                                           V   |                
                                (==============|==============) 
                                (      (parent address)       )
                       +------->(        <word data>          )<--------------+
                       |        (                             )               |
                       |        ( (lhs address) (rhs address) )               |
                       |        (========|=============|======)               |
                       |                 |             |                      |
                       |    +------------+             +------------------+   | 
                       |    |                                             |   |
                       |    V                                             V   |
              (========|====================)                  (==============|==============)
              (      (parent address)       )                  (      (parent address)       )
              (        <word data>          )                  (         <word data>         ) 
              (                             )                  (                             )      
              ( (lhs address) (rhs address) )                  ( (lhs address) (rhs address) )
              (========|=============|======)                  (========|=============|======)
                       |             |                                  |             |
                       V             V                                  V             V
                      (0)           (0)                                (0)           (0)

    (0) -> nil pointer address 
```


## File Format
### General format
#### File Header
```
    00 00                   - 16 bit file version header.
    00 00                   - 16-bit Flags (See below)
    00 00 00 00 00 00 00 00 - 64 bit root node address (relative to file position 0).  
    <arbitrary 'node' records>
```
#### Flags
```
    0 0 0 0 0 0 0  0 0 0 0 0 0 0 0 0
    F E D C B A 9  8 7 6 5 4 3 2 1 0
    | | | | | | |  | | | | | | | | +--- Compression enabled (GZIP) (01 -> 1)
    | | | | | | |  | | | | | | | | 
    | | | | | | |  | | | | | | | | 
    | | | | | | |  | | | | | | | +----- Encryption enabled (AES) (10 -> 2)
    | | | | | | |  | | | | | | |   
    | | | | | | |  | | | | | | |   
    +-+-+-+-+-+-+--+-+-+-+-+-+-+----------- Reserved    

```

#### Node Records
```
====================================================
    -- 32-bit Node Record Header --
    00                - 8-bit address size (ParentNodeAddrSz: Parent node address size, in bytes)
    00                - 8-bit address size (LhsNodeAddrSz: Lhs node address size, in bytes)
    00                - 8-bit address size (RhsNodeAddrSz: Rhs node address size, in bytes)
    00                - 8-bit wordSize size (DataLenSz: Size of the word size record).
    -- Arbitrary-Length Node Record Data --
    00 ... 00         - Parent Address (ParentNodeAddr: size determined by ParentNodeAddrSz)
    00 ... 00         - Lhs Node Address (LhsNodeAddr: size determined by LhsNodeAddrSz)
    00 ... 00         - Rhs Node Address (RhsNodeAddr: size determined by RhsNodeAddrSz)
    00 ... 00         - Data Size (DataLength: size determined by DataLenSz)
    00 00 ... 00 00   - Data Record (Payload data--DataLength bytes)

====================================================
```
