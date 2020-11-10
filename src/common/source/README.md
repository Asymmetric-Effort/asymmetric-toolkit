Common Source Generator
=======================
(c) 2020 Sam Caldwell.  All Rights Reserved.

> This library exists to provide a commandline flag (e.g. -source),
> its validation logic, parser and a generator of feed data.  When
> implemented in a project, the implementing project will inherit
> the commandline flag and the ability to generate feed data given
> input parameters.

## Usage

### CommandLine
```bash
<prog> [-source &lt;string&gt; [-srcFile &lt;string&gt;] [-srcRegex &lt;string&gt;][-srcCount &lt;int&gt;]
```
* `-source` (required. Default: `sequence`) specifies the type of source (random, sequence, 
  dictionary)
  
* `-srcFile` (required for `dictionary`) applies to dictionary source only and identifies the
  source file to be consumed.
  
* `-srcRegex` (optional. Default: `disabled`) specifies a filtering regular expression to be
  applied to elements fo the feed.
  
* `-srcCount` (optional. Default: `1048576)` specifies the number of elements to generate.

### Implementation
***TBD***

## Theory of Operation
On implementation the consuming program will--
1. Initialize the commandline flags.
2. Validate the commandline inputs.
3. Configure internal state.
4. Initialize the source feed generator in a 'paused' state.

When the source feed generator is 'un-paused', the generator will--
1. Produce an element from the data source (randomizer, sequencer or dictionary).
2. Apply any regular expression filter and only return matching data elements.