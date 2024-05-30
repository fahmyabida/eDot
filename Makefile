# Exporting bin folder to the path for makefile
export PATH   := $(PWD)/bin:$(PATH)
# Default Shell
export SHELL  := bash
# Type of OS: Linux or Darwin.
export OSTYPE := $(shell uname -s)

build: ## Builds binary
	@ printf "Building aplication... "
	@ go build \
		-trimpath  \
		-buildvcs=false \
		-o brick-transfer \
		./cmd/
	@ echo "done"

