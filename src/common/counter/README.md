Counter
=======

### Objective:
* To create an arbitrary length counter for mapping sequential numbers
  to arbitrary character sets.  This is used, for example, by the source
  module in this repo.
  
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