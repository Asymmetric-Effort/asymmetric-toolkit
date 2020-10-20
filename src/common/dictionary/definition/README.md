Dictionary Definition
=====================

## Purpose
Where the dictionary is a collection of definitions, this
package provides the concept of a definition.  Each definition
consists at a minimum of the `id` and `word` and in the end a
collection of statistics.  The `id` is a sha256 hash of the 
cleartext `word` text used to uniquely and safely identify a
definition throughout the toolkit.  By contrast, the `word`
is the payload object used in various attacks, stored in an
encrypted, base64 encoded form.
