# pulling a lightweight version of golang
FROM golang:1.8-alpine
MAINTAINER Paul Johnson <paulsjohnson91@gmail.com>
RUN apk --update add --no-cache git

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/paulsjohnson91/challenge-accepted
WORKDIR /go/src/github.com/paulsjohnson91/challenge-accepted

# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN go get github.com/paulsjohnson91/challenge-accepted
RUN go build .

# Run the command by default when the container starts.
#ENTRYPOINT ["go run /go/bin/github.com/paulsjohnson91/challenge-accepted/main.go"]

CMD ["./challenge-accepted"]

# Document that the service listens on port 3333.
EXPOSE 3333
