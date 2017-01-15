NAME := github-notifier
SRCS := $(shell find . -name '*.go' -type f)

.DEFAULT_GOAL := bin/$(NAME)

bin/$(NAME): $(SRCS)
	go build -o bin/$(NAME)

.PHONY: clean
clean:
	rm -rf bin/*
	rm -rf vendor/*

.PHONY: install
install:
	go install
