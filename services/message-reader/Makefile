GO=go
GOTEST=$(GO) test
GOCOVER=$(GO) tool cover
.PHONY: test/cover
test/cover:
		$(GOTEST) -v -coverprofile=tmp/coverage.out ./...
		$(GOCOVER) -func=tmp/coverage.out

test/cover/html:
		$(GOTEST) -v -coverprofile=tmp/coverage.out ./...
		$(GOCOVER) -func=tmp/coverage.out
		$(GOCOVER) -html=tmp/coverage.out

