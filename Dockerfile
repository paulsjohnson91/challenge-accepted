# pulling a lightweight version of golang
FROM golang:1.8-alpine
MAINTAINER Mauricio Ribeiro <maumribeiro@gmail.com>
RUN apk --update add --no-cache git

# Copy the local package files to the container's workspace.
ADD . /go/src/challenge-accepted
WORKDIR /go/src/challenge-accepted

# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN go get challenge-accepted

# Run the command by default when the container starts.
ENTRYPOINT ["/go/bin/challenge-accepted"]

# Document that the service listens on port 9000.
EXPOSE 9000
