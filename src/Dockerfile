FROM golang:1.19 as builder
RUN mkdir /go/src/mopi
COPY . /go/src/mopi/
WORKDIR /go/src/mopi
RUN go mod tidy
RUN CGO_ENABLED=0 go build -o /mopi cmd/main.go

FROM scratch
COPY --from=builder /mopi /
COPY configurations/ /configurations/
ENTRYPOINT ["./mopi"]
