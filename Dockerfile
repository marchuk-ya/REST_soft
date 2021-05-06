FROM golang

WORKDIR /go/src/REST_soft
COPY . .

RUN go mod download
RUN go install -v ./...

RUN go build -o /main.go .

EXPOSE 8080

CMD ["mini-REST"]