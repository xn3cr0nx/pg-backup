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
SCRIPTS_PATH=scripts

LNX_BUILD=$(build)/$(BINARY_NAME)
WIN_BUILD=$(build)/$(BINARY_NAME).exe

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

# backup
.PHONY: backup
backup:
	$(GORUN) $(BACKUP) export

.PHONY: build_backup
build_backup: 
	$(GOBUILD) -o $(BUILD_PATH)/$(BACKUP_BINARY) -v $(BACKUP)

.PHONY: install_backup
install_backup:
	$(GOINSTALL) $(BACKUP)

.PHONY: build
build: build_backup

.PHONY: install
install: install_backup

# # Cross compilation
# linux: $(LNX_BUILD)
# windows: $(WIN_BUILD)
# # deploy:
# # 	ansible-playbook -i deploy/inventory.txt deploy/deploy.yml

# $(LNX_BUILD):
# 	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BUILD_PATH)/$(BACKUP_BINARY) -v $(BACKUP)
# $(WIN_BUILD):
# 	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOBUILD) -o $(BUILD_PATH)/$(BACKUP_BINARY).exe -v $(BACKUP)