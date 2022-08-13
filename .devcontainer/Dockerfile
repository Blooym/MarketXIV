FROM golang:1.19-alpine

# Set the go version to use here.
ARG GOLANG_VERSION=1.19

# Install Git and OpenSSH.
RUN apk add --update git openssh

# Create a non-root dev user and swap.
RUN adduser --disabled-password --gecos "" dev
USER dev

# Install useful go packages.
RUN go install github.com/goreleaser/goreleaser@latest 
RUN go install gotests gomodifytags impl goplay dlv staticcheck gopls