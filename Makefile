# Go and compilation related variables
BUILD_DIR ?= out
SOURCE_DIRS = cmd pkg test
MAC_KEY_PATH ?= /home/ariobolo/01WORKSPACE/04SOURCE/gitlab.redhat/qe-platform/qe-platform/config/private/hosts/mac-2-brno-key

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
cross: $(BUILD_DIR)/gomacx

$(BUILD_DIR)/gomacx: $(SOURCES)
	CC=clang GOARCH=amd64 GOOS=darwin go build -o $(BUILD_DIR)/gomacx .

# Create and update the vendor directory
.PHONY: remote
remote:
	scp -i $(MAC_KEY_PATH) -r ${PWD}/* crcqe@macmini-crcqe-2.tpb.lab.eng.brq.redhat.com:gomacx

# https://github.com/gshen7/macOSNotes/blob/master/README.md
# https://gist.github.com/gerad/1645235
# https://stackoverflow.com/questions/28452264/how-to-get-nsaccessibilitychildren
# https://stackoverflow.com/questions/60677904/incompatible-pointer-types-passing-struct-nsarray-to-parameter-of-type-nsar
# https://github.com/tejasmanohar/go-libproc/blob/master/libproc.go
# https://coderwall.com/p/l9jr5a/accessing-cocoa-objective-c-from-go-with-cgo


# https://www.electronjs.org/docs/latest/tutorial/accessibility#macos
# https://github.com/gngrwzrd/objc-gw/blob/master/AccessibilityHelper.m
# https://fosdem.org/2023/schedule/event/govfkit/attachments/slides/5847/export/events/attachments/govfkit/slides/5847/fosdem2023_go_devroom_vfkit.pdf



# https://gist.github.com/zaru/2405fee754d25cb16a1622a4187758bc