Counter
=======

### Objective:
* To create an arbitrary-length counter for generating sequences of words using a given character set
  where each word is in ascending sequential order.  This counter should generate each sequence regardless
  of the character set given assuming the relative value of one character to the next according to the character's
  position in the character set array given as input when initialized.

### Theory of Operation:

1. Create an array of n uint8 elements.

2. The `.Setup()` method will provide a field of characters which will
   represent the base of our counter's numbering system.

3. When `.Increment()` is called, increment the counter starting with the
   least significant element and carrying the one when the numbering base
   is exhausted, recursively.
   
4. When `.String()` is called, the counter values will be mapped to the
   provided base character set and then concatenated into a single string
   which will be returned to the caller.
   
### Example of Operation

Given the input character sequence `0123456789`, and a word size of two (2), the counter would produce the following:
00, 01, 02, 03, 04, 05, 06, 07, 08, 09, 10, ..., 99.