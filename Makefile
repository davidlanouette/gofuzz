
go=go1.18beta2

.PHONY: run
run:
	$(go) run .

.PHONY: test
test:
	$(go) test

.PHONY: testv
testv:
	$(go) test -v .

.PHONY: fuzz
fuzz:
	$(go) test -v -fuzz .

.PHONY: fuzztimed
fuzztimed:
	$(go) test -v -fuzz=Fuzz -fuzztime 30s
