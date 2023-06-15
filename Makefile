export GONOPROXY=https://github.com/AnimusPEXUS/*

all: get build

get:
		$(MAKE) -C tests/test01 get
		go get -u -v "./..."
		go mod tidy

build:
		$(MAKE) -C tests/test01 build
		go build

