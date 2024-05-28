all: set-arch build

# ------------------------ golang --------------------------
# golang's order
GO_CMD=go
GO_SET=set
GO_BUILD=$(GO_CMD) build
GO_CLEAN=$(GO_CMD) clean
GO_TEST=$(GO_CMD) test
GO_ARCH=amd64

# product's info exp: name & version & runtime os
BINARY_NAME=ginfx-templete
BINARY_VERSION=_v1.0
OS := $(shell $(GO_CMD) env GOOS)

set-arch:
	$(GO_SET) GOARCH=$(GO_ARCH)

# should be set GOOS when before this building
# like:
#	set GOOS=windows
#	make all OR make build
ifeq ($(OS),windows)
build:
	$(GO_BUILD) -o $(BINARY_NAME)$(BINARY_VERSION).exe -v -x
else
build:
	$(GO_BUILD) -o $(BINARY_NAME)$(BINARY_VERSION) -v -x
endif

test:
	go test -covermode=count -coverpkg=./... -coverprofile=coverage.out


clean:
	$(GO_CLEAN)
	rm -f $(BINARY_NAME)$(BINARY_VERSION).exe
	rm -f $(BINARY_NAME)$(BINARY_VERSION)


# ------------------------ angular --------------------------
frontend:
	cd ui && yarn run build

# ensures that these targets do not conflict with real file names and guarantees that they are always executed,
# regardless of the existence of files with the same name
.PHONY: *