CC=gcc
CFLAGS=-I.
SRC=src
BUILD_DIR=build
PROG_NAME=main
MKDIR=mkdir -p

all: $(PROG_NAME)
$(PROG_NAME): $(SRC)/main.c $(SRC)/parser.c $(SRC)/records.c
	$(MKDIR) $(BUILD_DIR)
	$(CC) -o $(BUILD_DIR)/$(PROG_NAME) -g $^ $(CFLAGS)

.PHONY: all clean $(SRC)
clean:
	rm -rf $(BUILD_DIR)
