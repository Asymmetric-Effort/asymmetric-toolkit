Cli Dev
=======

## Purpose:
* Provide developers with a reference implementation of [common/cli](../../common/cli/README.md) which can be used
  to create new tools within the Asymmetric Toolkit framework.
* Provide a high-level integration test which can be executed to prove that [common/cli](../../common/cli/README.md)
  is working as expected.
  
## Usage:
### For Software Development
This tool is intended for use by software developers seeking to learn how to implement the 
[common/cli](../../common/cli/README.md) library within their tools.  To that end, usage of this tool is simply
by reference to the code found [here](../clidev).

### For Build / Test Processes
While all unit tests may pass for the [common/cli](../../common/cli/README.md), the ultimate determination of whether
[common/cli](../../common/cli/README.md) is working should be this test program.  If [clidev](../clidev) builds and if
its tests pass, then we should be able to assert that the [common/cli](../../common/cli/README.md) module works as
intended.

### Caveat
It is not the intent of the Asymmetric Toolkit project to create a test program for all common modules in the
mono-repo.  This should be done very judiciously.  In the case of [clidev](../clidev), we created a reference 
implementation tool because the program interacts with commandline arguments and is difficult at best to test.
