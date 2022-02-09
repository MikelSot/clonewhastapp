FROM golang:1.17-alpine AS builder

WORKDIR /go/src/clonewhastapp

# add some necessary packages
RUN apk update && \
    apk add libc-dev && \
    apk add gcc && \
    apk add make

# prevent the re-installation of vendors at every change in the source code
COPY go.mod .
# if exists
#COPY go.sum .
RUN go mod download  # && go mod verify  [<- if exists]

# Install Compile Daemon for go. We'll use it to watch changes in go files
RUN go get github.com/githubnemo/CompileDaemon

# Copy and build the app
COPY . .
COPY ./entrypoint.sh /entrypoint.sh

# wait-for-it requires bash, which alpine doesn't ship with by default. Use wait-for instead
ADD https://raw.githubusercontent.com/eficode/wait-for/v2.1.0/wait-for /usr/local/bin/wait-for
RUN chmod +rx /usr/local/bin/wait-for /entrypoint.sh

ENTRYPOINT [ "sh", "/entrypoint.sh" ]















#RUN apt-get update
#ENV GO111MODULE=on \
#    CGO_ENABLED=0 \
#    GOOS=linux \
#    GOARCH=amd64
#
#WORKDIR /go/src/clonewhastapp
#COPY go.mod .
#RUN go mod download
#COPY . .
#RUN go build cmd/main.go

#FROM scratch
#COPY --from=builder /go/src/clonewhastapp .
#ENTRYPOINT ["./main"]
#CMD ["cd cmd", 'CompileDaemon -command="./cmd"']



