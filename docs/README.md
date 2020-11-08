Documentation Index
===================

## Toolkit User Guides
* [Pen Tester's Getting Started Guide](./user/getting-started.md)
  > Documentation intended for Asymmetric Toolkit users.

* [Developer's Getting Started Guide](./developer/getting-started.md)
  > Documentation for software developers looking to help grow the Asymmetric toolkit and
  > empower its users.

## Tools 
* [dnsEnum: DNS Enumeration Tool](../src/tools/dnsenum/README.md)
* [vladTheInjector: Injection Tool](../src/tools/vladTheInjector/README.md)

## Tools: Common Code
* [tools/common/cli](../src/common/cli/README.md)
  >common/cli is the reusable commandline processor code.
                                                    
* [tools/common/counter](../src/common/counter/README.md)
  >common/counter creates a simple counter for generating sequences in various keyspaces
  >of various sizes.

* [tools/common/dictionary](../src/common/dictionary/README.md)
  >common/dictionary creates, manages and reads a proprietary, encrypted dictionary file
  >format specific to the asymmetric toolkit.

* [tools/common/entropy](../src/common/entropy/README.md)
  >common/entropy is a set of tools for measuring and working with randomness.

* [tools/common/errors](../src/common/errors/README.md)
  >common/errors is a set of tools for working with errors and general error handling.

* [tools/common/fifo](../src/common/fifo/README.md)
  >common/fifo is a simple first-in/first-out queue data structure.

* [tools/common/file](../src/common/file/README.md)
  >common/file provides some simple file-handling facilities.

* [tools/common/logger](../src/common/logger/README.md)
  >common/log provides a robust logging platform for stdout, stderr, file and syslog.

* [tools/common/misc](../src/common/misc/README.md)
  >common/misc provides miscellaneous code that doesn't fit elsewhere in the project.

* [tools/common/source](../src/common/source/README.md)
  >common/source provides a reusable data source for brute-force, dictionary or other
  >enumeration attacks.  This includes sequential data sources, random data sources as
  >a data source which consumes an asymmetric-toolkit dictionary file format source.
  >Each source provides an asynchronous feed of 'words' (character strings) which can
  >be consumed via a channel by various tool payloads.

* [tools/common/stats](../src/common/stats/README.md)
  >common/stats provides an abstraction for stats collection.

* [tools/common/types](../src/common/types/README.md)
  >common/types attempts to provide a standardized set of common data types.

