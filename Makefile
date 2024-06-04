all: set-arch build

# ------------------------ golang --------------------------
# golang's order
GO_CMD=go
GO_SET=set
GO_LINT=$(GO_CMD)lint
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

# 注释风格：导出的函数和类型应该有适当的注释。
# 命名规范：变量、函数、类型等的命名应该符合 Go 的命名惯例。
lint:
	$(GO_LINT) ./...

# Printf 格式检查：检查 Printf、Sprintf 等函数的格式化字符串与参数类型是否匹配。
# 结构标签：检查结构体字段标签的格式是否正确。
# 无效的操作：例如，对 nil 指针进行操作。
# 导入路径：检查导入路径是否正确。
vet:
	$(GO_CMD) vet ./...

test:
	$(GO_CMD) test -covermode=count -coverpkg=./... -coverprofile=coverage.out


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