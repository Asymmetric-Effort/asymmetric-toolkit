Common/CommandLine Processor
----------------------------

(c) 2020 Sam Caldwell.  See LICENSE.txt.


## Purpose
* Provide a reusable commandline facility which can be used across the asymmetric-toolkit project
  with minimal effort.
* Provide standardized arguments for consistency across the code repository (all tools developed)
  to reduce the learning curve for users.

## Usage

### Creating a Specification
The following code snippet illustrates how to use the common/cli facility:
```
var o Specification

o.Version = "1.1.1"
o.ProgramName = "program"
o.Copyright = "copyright"
o.Argument = make(map[string]ArgumentDescriptor)
o.Argument["myFlag1"] = ArgumentDescriptor{
    1000,
    String,
    "myDefault1",
    "MyHelpString1",
    nil,
    ExpectNone,
}
```

In the above, we see that the first step is to create a `Specification`.  After that we need to--
1. Configure the `Specification` with `Version`, `ProgramName`, `Copyright` adn `Description`.
2. Add our `ArgumentDescriptor` objects to the `Specification` Argument map.

When we define an `ArgumentDescriptor`, we define the following elements--
* `FlagId` (type `ArgumentFlag`): indicating the internal identifier for a given flag (see below).
* `Type` (type `ArgumentType`): defining the datatype of the argument (e.g. Integer).
* `Default` (type `string`): specifying the default value in string format.
* `Help` (type `string`):  declaring the help text which will appear when `-h` or `--help` is used.
* `Parse` (type `func`): referencing the Parser Factory Function used to parse/validate the commandline
  string value for later extraction to its final internal state (datatype).
* `Expects` (type `NextExpected`): instructs the `Commandline::Parse()` function what token type to expect
  next (e.g. `ExpectNone`, `ExpectEnd`, `ExpectFlag` and `ExpectValue`).
  
### Using the Specification
Once the `Specification` is defined, we need to execute `Commandline::Parse()` to parse the commandline arguments
from `os.Args` into the `cli.CommandLine` internal state:
```
var ui cli.CommandLine
exit, err := ui.Parse(&spec, os.Args[1:])
if err != nil {
    fmt.Printf("Error:%v", err)
    os.Exit(cli.ErrArgumentParseError)
}
if exit {
    fmt.Printf("%v", err)
    os.Exit(cli.ErrArgumentParseError)
}
```
We can see in this snippet (above) that we define `ui` as `cli.Commandline`.  We then call `ui.Parse()` and 
pass a reference to the `specification` object along with the raw commandline arguments from `os.Args`.  Note
that we pass only `os.Args[1:]`,not the program name (`os.Args[0]`).  This keeps things simple.

Once the `CommandLine::Parse()` function finishes, the resulting `CommandLine` object has the parsed (but 
un-extracted) commandline state.  At this point the program may extract the state into an application-specific
structure and free the `Commandline` object from memory.

