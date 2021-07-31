FROM --platform=amd64 golang:1.17 as builder
WORKDIR /build
COPY go.mod *.go /build/
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -trimpath -ldflags "-s -w" -o /build/iw-conference-recorder-reciever

FROM --platform=amd64 scratch
COPY --from=builder /build/iw-conference-recorder-reciever /iw-conference-recorder-reciever
EXPOSE 80
ENTRYPOINT [ "/iw-conference-recorder-reciever" ]