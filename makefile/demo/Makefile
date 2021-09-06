SHELL              := /bin/bash
# go options
GO                 ?= go
LDFLAGS            :=
GOFLAGS            :=
BINDIR             ?= $(CURDIR)/bin

.PHONY: build
build: demo

DEMO_BINARIES := demo-darwin demo-linux demo-windows
$(DEMO_BINARIES): demo-%:
	@GOOS=$* $(GO) build -o $(BINDIR)/$@ $(GOFLAGS) -ldflags '$(LDFLAGS)' $(CURDIR)/main.go
	@if [[ $@ != *windows ]]; then \
	  chmod 0755 $(BINDIR)/$@; \
	else \
	  mv $(BINDIR)/$@ $(BINDIR)/$@.exe; \
	fi

.PHONY: demo
demo: $(DEMO_BINARIES)

.PHONY: clean
clean:
	@rm -rf $(BINDIR)