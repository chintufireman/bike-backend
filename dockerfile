FROM golang:1.21 AS build

WORKDIR /src

COPY . .

RUN CGO_ENABLED=0 go build -o bin .

#second stage
FROM scratch

WORKDIR /src

COPY --from=build /src/bin .
ENTRYPOINT ["/src/bin"]