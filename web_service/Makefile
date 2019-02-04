# All functions of makefile with go commands
CMD=go
BUILD=$(CMD) build
CLEAN=$(CMD) clean
TEST=$(CMD) test
GET=$(CMD) get
# The code output after compiling all files in the directory
BINARY=webservice
BINARY_UNIX=$(BINARY)_unix

# The parameters for Makefile
all: test build
build:
	$(BUILD) -o $(BINARY) -v
test:
	$(TEST) -v ./...
clean:
	$(CLEAN)
	rm -f $(BINARY)
	rm -f $(BINARY_UNIX)
run:
	$(BUILD) -o $(BINARY) -v ./...
	./$(BINARY)
