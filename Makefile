.DEFAULT=help
ROOT_DIR:=$(shell pwd -P)

help:
	@echo " "
	@echo "asymmetric-toolkit Makefile"
	@echo "---------------------------"
	@echo "(c) 2018 Sam Caldwell.  See LICENSE.txt"
	@echo " "
	@echo " "

routers/base:
	@docker build --tag asymmetric-toolkit:routers.base \
				  --compress \
				  -f ${ROOT_DIR}/docker/routers/base.docker .

tools/base:
	@docker build --tag asymmetric-toolkit:tools.base \
				  --compress \
				  -f ${ROOT_DIR}/docker/tools/base.docker .

osint/base:
	@docker build --tag asymmetric-toolkit:osint.base \
				  --compress \
				  -f ${ROOT_DIR}/docker/osint/base.docker .

