FROM golang:latest

LABEL maintainer="Kyle Lee <kyleleemm@gmail.com>"

WORKDIR /app

# get requirements, hash information, then install modules
COPY go.mod .
COPY go.sum .
RUN go mod download 

COPY . . 

ENV PORT 8080

RUN go build 

# we only need a single compiled file so remove all source files
RUN find . -name "*.go" -type f -delete

EXPOSE $PORT 

CMD ["./ttweb-poc"]

