Asymmetric-toolkit
==================
(c) 2018 Sam Caldwell.  See LICENSE.txt

## PURPOSE:
* To create a container-based toolkit for penetration testing.
 which is backed by a CI/CD pipeline for rapid development as the platform evolves.

## USAGE:
### Setup:
1. Clone this repository
2. Build the container image
```bash
make image
```

### Getting Started
This tool is intended to make starting work simple.  Just execute the following to shell into the container with all the things:
```bash
make hack
```
This will start a shell into the container with all the tools the container offers and any external resources provided.

