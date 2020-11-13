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

### ToDo:
* Working on [common/logger](src/common/logger/README.md) to provide rich, high-cardinality structured event logging
  to stdout, file or syslog ensuring the efficient use of network and system resources.  This leaves the path open
  for a compression and encryption layer to our logging strategy.
* Continue work on [dnsEnum](src/tools/dnsenum/README.md).
  > This tool is our first attempt at and end-to-end golang solution within the framework.  Proof-of-concept efforts
  > have been the basis for all work performed within the project thus far.  Now with 
  > [common/cli](src/common/cli/README.md) finished, we can refactor existing code to implement 
  > [common/cli](src/common/cli/README.md) and move forward to the next hurdle.

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