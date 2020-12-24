GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOINSTALL=$(GOCMD) install
MAKE=make

BUILD_PATH=build
BACKUP=./cmd/backup
BACKUP_BINARY=backup
CMD=export

export GO111MODULE=on

default: build_backup backup

.PHONY: all
all: test build linux

.PHONY: test
test:
	$(GOTEST) -v ./...

.PHONY: clean
clean:
	$(GOCLEAN)
	rm -f $(BUILD_PATH)

.PHONY: run_backup
run_backup:
	$(GORUN) $(BACKUP) export

.PHONY: build_backup
build_backup: 
	$(GOBUILD) -o $(BUILD_PATH)/$(BACKUP_BINARY) -v $(BACKUP)

.PHONY: install_backup
install_backup:
	$(GOINSTALL) $(BACKUP)

.PHONY: backup
backup:
	$(BUILD_PATH)/$(BACKUP_BINARY) $(CMD)
