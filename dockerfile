FROM golang:1.21.5-bookworm


WORKDIR /app

COPY go.mod .
COPY ./ .

RUN go get
RUN go build -o bin .

ENTRYPOINT [ "/app/bin" ]