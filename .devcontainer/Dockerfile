FROM golang:1.21-alpine

# Install Git and OpenSSH.
RUN apk add --no-cache --update git openssh curl upx

# Create a non-root dev user and swap.
RUN adduser --disabled-password --gecos "" dev
USER dev

# Install useful go packages.
RUN go install github.com/goreleaser/goreleaser@latest