
# Go and compilation related variables
BUILD_DIR ?= out
SOURCE_DIRS = cmd pkg test


# https://golang.org/cmd/link/
# LDFLAGS := $(VERSION_VARIABLES) -extldflags='-static' ${GO_EXTRA_LDFLAGS}

.PHONY: clean ## Remove all build artifacts
clean:
        rm -rf $(BUILD_DIR)

# Create and update the vendor directory
.PHONY: vendor
vendor:
        go mod tidy
        go mod vendor

.PHONY: cross ## Cross compiles all binaries
cross: $(BUILD_DIR)/gomacx.exe

$(BUILD_DIR)/gomacx.exe: $(SOURCES)
        CC=clang GOARCH=amd64 GOOS=darwin go build -o $(BUILD_DIR)/gomacx.exe .



# https://github.com/gshen7/macOSNotes/blob/master/README.md
# https://gist.github.com/gerad/1645235
# https://stackoverflow.com/questions/28452264/how-to-get-nsaccessibilitychildren
# https://stackoverflow.com/questions/60677904/incompatible-pointer-types-passing-struct-nsarray-to-parameter-of-type-nsar