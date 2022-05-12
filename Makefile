SHELL=/bin/bash -o pipefail
pwd=$(shell pwd)
project_module=github.com/panther-labs/stix2


verify: fmt importorder vet staticcheck

verify-test: verify test

test:
	go test -race ./...

# staticcheck requires staticcheck. to install run the following outside of the repo:
# GO111MODULE="off" go get honnef.co/go/tools/cmd/staticcheck
#
staticcheck: ## runs staticcheck on our packages
	staticcheck $(project_module)/...

vet:
	go vet ./...

# importorder requires impi. to install run the following outside of the repo:
# GO111MODULE="off" go get github.com/pavius/impi/cmd/impi
#
importorder: ## Verifies all code has correct import orders (stdlib, 3rd party, internal)
	@STATUS=0 ; \
	for f in `find . -type d | grep -v wasm` ; do \
		file=$$(cd $$f && impi --local $(project_module) --scheme stdThirdPartyLocal --ignore-generated=true .) ; \
		if [[ $$file ]] ; then \
		echo "$$f/$$file" ; \
			STATUS=1 ; \
		fi ; \
	done ; \
	if [ $$STATUS -ne 0 ] ; then \
		exit 1 ; \
	fi

fmt: ## Verifies all code is gofmt'ed
	@STATUS=0 ; \
	for f in `find . -type f -name "*.go"` ; do \
		file=$$(gofmt -l $$f) ; \
		if [[ $$file ]] ; then \
			echo "file not gofmt'ed: $$f" ; \
			STATUS=1 ; \
		fi ; \
	done ; \
	if [ $$STATUS -ne 0 ] ; then \
		exit 1 ; \
	fi