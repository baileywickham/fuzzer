FROM golang:1.15.0-alpine3.12
RUN apk add --no-cache git

WORKDIR /go/src/app
RUN git clone https://github.com/strongcourage/fuzzing-corpus.git
COPY . .
RUN go get -d -v ./...

EXPOSE 8080
CMD go run . --port=8080 fuzzing-corpus/xml/mozilla

