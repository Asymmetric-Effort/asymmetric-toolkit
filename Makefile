.DEFAULT=help
ROOT_DIR:=$(shell pwd -P)

help:
	@echo " "
	@echo "asymmetric-toolkit Makefile"
	@echo "---------------------------"
	@echo "(c) 2018 Sam Caldwell.  See LICENSE.txt"
	@echo " "
	@echo " This is a work in progress.  Sam should implement this...."
	@echo " "
	exit 1

dnsenum_test:
	@(cd tools/dnsenum && go test -v)
	#ToDo: Makefile needs reusable pieces.  I'm lazy, remember.

dnsenum: dnsenum_test
	@(cd tools/dnsenum/ && go build -o ../../build/tools/dnsenum main.go)
	#ToDo: We need to build macos, windows, linux spread here.
	#ToDo: Makefile needs reusable pieces.  I'm lazy, remember.
