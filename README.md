Asymmetric-toolkit
==================
(c) 2018 Sam Caldwell.  See LICENSE.txt.

## Project Status
This is a work in progress.  It's still pretty early, and all the things are subject to change.

### Milestones:
Note these milestones are not exhaustive but provide a summary of latest efforts/achievements/priotized tasks:
* [Documentation](docs/README.md).
* [common/cli](src/common/cli/README.md) is working.  Tests written. [clidev](src/tools/clidev/README.md) is
  ready, providing an example implementation of [common/cli](src/common/cli/README.md).
* [common/types/tags](src/common/types/tags/README.md) is now complete.
* [common/logger](src/common/logger/README.md) is feature complete.

### ToDo:
* Working on [tools/dictionary/createDictionary](src/tools/dictionary/createDictionary/README.md) and the
[common/dictionary](src/common/dictionary/README.md) to finish out the dictionary writer then the dictionary reader
facilities.  These will ingest textfile dictionaries and output proprietary dictionary files.
* Continue work on [dnsEnum](src/tools/dnsenum/README.md).
  > This tool is our first attempt at and end-to-end golang solution within the framework.  Proof-of-concept efforts
  > have been the basis for all work performed within the project thus far.  
  > The tool will require [common/dictionary](src/common/dictionary/README.md) and at least one dictionary file.

## Purpose:
* To Provide a thin set of compact, container-based pen testing tools written in golang,
  compressed and capable of interoperability either on a single machine or distributed
  across many machines.  These tools should have a CI/CD build pipeline, unit and integration
  tests and be deployable either as Linux, MacOS or Windows binaries or within tool-specific
  containers.
  
## Disclaimer:
* The author specifically licenses this software under an MIT License (See LICENSE.txt)
  with one reservation.  This software will be used only for lawful purposes under the
  laws of the United States.

## Documentation
* [General Documentation Index](docs/README.md)