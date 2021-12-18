help:		## Show this help.
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

build:		## Build the gh-discussions binary
	go build -ldflags "-X github.com/pete0emerson/gh-discussions/cmd.version=$(VERSION)" .

install:	## Install the gh-discussions binary as a gh extension
	gh extension remove discussions || true
	gh extension install .


clean:		## Remove the gh-discussions binary
	rm -f ./gh-discussions

.PHONY: help
