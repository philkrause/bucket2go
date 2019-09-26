LDFLAGS = -ldflags "-X main.BuildVersion=$$VERSION"
GOARCH = amd64
linux: export GOOS=linux
darwin: export GOOS=darwin
windows: export GOOS=windows

all: linux darwin windows

build:
	go build ${LDFLAGS}

run:
	go build ${LDFLAGS}
	./gobit

linux:
	go build $(LDFLAGS) -o gobit
	mkdir -p release
	rm -f release/gobit-${GOOS}_${GOARCH}.zip
	zip release/gobit-${GOOS}_${GOARCH}.zip gobit
	rm -f gobit

darwin:
	go build $(LDFLAGS) -o gobit
	mkdir -p release
	rm -f release/gobit-${GOOS}_${GOARCH}.zip
	zip release/gobit-${GOOS}_${GOARCH}.zip gobit
	rm -f gobit

windows:
	go build $(LDFLAGS) -o gobit.exe
	mkdir -p release
	rm -f release/gobit-${GOOS}_${GOARCH}.zip
	zip release/gobit-${GOOS}_${GOARCH}.zip gobit.exe
	rm -f gobit.exe

.PHONY: clean
clean:
	rm -rf release
	rm -f gobit gobit.exe