CC=gcc
CGFLAGS=-I.
SRC=src
BUILD_DIR=build
PROG_NAME=mirrors
MKDIR=mkdir -p

all: $(PROG_NAME)
$(PROG_NAME): $(SRC)/main.c $(SRC)/parser.c
	$(MKDIR) $(BUILD_DIR)
	$(CC) -o $(BUILD_DIR)/$(PROG_NAME) -g $^ $(CGFLAGS)

.PHONY: clean all
clean:
	rm -rf $(BUILD_DIR)