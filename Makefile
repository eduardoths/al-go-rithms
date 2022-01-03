GO ?= go
GORUN ?= $(GO) run
GOTEST ?= gotest
GOTOOL ?= $(GO) tool

DIR_COVER ?= ./.cover
COVER_UNIT_FILE ?= unit.out
GO_COVERPKG ?= ./...

#-----------------------------------------#
# Tests
#-----------------------------------------#

.PHONY: test
test: _cover/setup test/run _cover/patch

.PHONY: test/run
test/run: GO_PKGS_UNIT_TEST ?= `go list $(GO_COVERPKG)`
test/run:
	@echo "Running go unit tests..."
	@$(GOTEST) -race -failfast -v \
		-coverpkg=$(GO_COVERPKG) -covermode=atomic -coverprofile=$(DIR_COVER)/$(COVER_UNIT_FILE).tmp \
		$(GO_PKGS_UNIT_TEST)


.PHONY: _cover/unit
_cover/unit:
	@echo "Coverage for: $(COVER_UNIT_FILE)"
	@$(GOTOOL) cover -func=$(DIR_COVER)/$(COVER_UNIT_FILE) | { grep -v "100.0%" || true; }

.PHONY: _cover/setup
_cover/setup:
	@rm -rf $(DIR_COVER)
	@mkdir -p $(DIR_COVER)

.PHONY: _cover/patch
_cover/patch:
	@touch $(DIR_COVER)/$(COVER_UNIT_FILE).tmp
	@cat $(DIR_COVER)/$(COVER_UNIT_FILE).tmp | { grep -v "src/tests/" > $(DIR_COVER)/$(COVER_UNIT_FILE) || true; }