FROM golang:alpine
ADD . /go/src/app
WORKDIR /go/src/app
ENV PORT=8080
ENV PORT=8001
CMD ["go", "run", "main.go"]