Dictionary/definition
=====================

## Purpose:
* A definition is a structure with a word, an identifier and various statistics used to prioritize
  words used in attacks.  Each definition is serialized/deserialized as it is read from or written
  to disk.
  

## Dataflow
```
    +--->(raw record length)                                            (raw record length)---->+
    |                                                                               |           |
    |                                                                               |           V
(disk)->(decrypt)->(decompress)->deserialize->(descriptor)->serialize->(compress)->(encrypt)->(disk)
```

## On-Disk Format
The following illustrates the format on disk for each definition:

```
     0 . . . . . . 31 . . . . . . . . . . . . . . . . . .n
 r   +--------------+-----------------------------------+
 0   | recordLength | compressed,encrypted record bytes |
     +--------------+-----------------------------------+
 1   | recordLength | compressed,encrypted record bytes |
     +--------------+-----------------------------------+
 2   | recordLength | compressed,encrypted record bytes |
     +--------------+-----------------------------------+
 3   | recordLength | compressed,encrypted record bytes |
 .   +--------------+-----------------------------------+
 .   .              .                                   .
 .   .              .                                   .
 .   .              .                                   .
 .   +--------------+-----------------------------------+
 r-2 | recordLength | compressed,encrypted record bytes |
     +--------------+-----------------------------------+
 r-1 | recordLength | compressed,encrypted record bytes |
     +--------------+-----------------------------------+
 r   | recordLength | compressed,encrypted record bytes |
     +--------------+-----------------------------------+

```
Data moves back and forth between this on-disk format and the in-memory runtime format through the
`Definition::Serialize()` and `Definition::Deserialize()` methods.
