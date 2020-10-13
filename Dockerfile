FROM golang:alpine

WORKDIR /var/app/

COPY . /var/app/

RUN go build -v cmd/back/main.go

CMD ./main