.PHONY: build build-all clean

# Variables
NAME=marketxiv
SKIP_COMPRESS?=true

# Build a single binary for the current system.
build:
ifeq (, $(shell which goreleaser))
	@echo "WARNING: You are using an unsupported build system. Please install 'goreleaser' to build this project as intended."
	rm -rf dist
	go build -ldflags="-s -w" -o ./dist/$(NAME)
else
	SKIP_COMPRESS=$(SKIP_COMPRESS) goreleaser build --snapshot --rm-dist --single-target
endif

# Build binaries for all supported systems.
build-all:
ifeq (, $(shell which goreleaser))
	@echo "WARNING: You are using an unsupported build system. Please install 'goreleaser' to build this project as intended."
	rm -rf ./dist

	@echo "Building for linux"
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ./dist/$(NAME)_linux_amd64
	GOOS=linux GOARCH=arm go build -ldflags="-s -w" -o ./dist/$(NAME)_linux_arm
	GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -o ./dist/$(NAME)_linux_arm64
	GOOS=linux GOARCH=386 go build -ldflags="-s -w" -o ./dist/$(NAME)_linux_386
	
	@echo "Building for darwin"
	GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o ./dist/$(NAME)_darwin_amd64
	GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -o ./dist/$(NAME)_darwin_arm64

	@echo "Building for windows"
	GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o ./dist/$(NAME)_windows_amd64.exe
	GOOS=windows GOARCH=386 go build -ldflags="-s -w" -o ./dist/$(NAME)_windows_386.exe
	GOOS=windows GOARCH=arm go build -ldflags="-s -w" -o ./dist/$(NAME)_windows_arm.exe
	GOOS=windows GOARCH=arm64 go build -ldflags="-s -w" -o ./dist/$(NAME)_windows_arm64.exe

	@echo "Building for netbsd"
	GOOS=netbsd GOARCH=amd64 go build -ldflags="-s -w" -o ./dist/$(NAME)_netbsd_amd64
	GOOS=netbsd GOARCH=arm go build -ldflags="-s -w" -o ./dist/$(NAME)_netbsd_arm
	GOOS=netbsd GOARCH=386 go build -ldflags="-s -w" -o ./dist/$(NAME)_netbsd_386
	GOOSS=netbsd GOARCH=arm64 go build -ldflags="-s -w" -o ./dist/$(NAME)_netbsd_arm64

	@echo "Building for freebsd"
	GOOS=freebsd GOARCH=amd64 go build -ldflags="-s -w" -o ./dist/$(NAME)_freebsd_amd64
	GOOS=freebsd GOARCH=arm go build -ldflags="-s -w" -o ./dist/$(NAME)_freebsd_arm
	GOOSS=freebsd GOARCH=386 go build -ldflags="-s -w" -o ./dist/$(NAME)_freebsd_386
 else 
	SKIP_COMPRESS=$(SKIP_COMPRESS) goreleaser build --snapshot --rm-dist
endif

clean:
	rm -rf ./dist

# Run the application and pass all arguments to it.
run:
	go run main.go $(filter-out $@,$(MAKECMDGOALS))