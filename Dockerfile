FROM golang:alpine
ADD . /go/src/github.com/afnarqui
WORKDIR /go/src/github.com/afnarqui
RUN apk add git
WORKDIR $GOPATH/src/github.com/afnarqui
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
ENV PORT=8080
ENV PORT=8001
ENV PORT=26257


CMD ["go", "run", "main.go"]