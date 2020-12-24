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
BACKUP_BINARY=pg-backup
CMD=export

export GO111MODULE=on

default: build backup

.PHONY: all
all: test build backup

.PHONY: test
test:
	$(GOTEST) -v ./...

.PHONY: clean
clean:
	$(GOCLEAN)
	rm -f $(BUILD_PATH)

.PHONY: run
run:
	$(GORUN) $(BACKUP)

.PHONY: build
build: 
	$(GOBUILD) -o $(BUILD_PATH)/$(BACKUP_BINARY) -v $(BACKUP)

.PHONY: install
install:
	$(GOINSTALL) $(BACKUP)

.PHONY: backup
backup:
	$(BUILD_PATH)/$(BACKUP_BINARY) $(CMD)
