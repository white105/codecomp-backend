FROM golang:latest AS builder
# Copy the code from the host and compile it
RUN go get github.com/golang/dep/cmd/dep
WORKDIR $GOPATH/src/codecomp-backend
COPY . ./
RUN dep ensure -v
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GODEBUG=http2debug=2 go build  -o /codecomp-backend

FROM alpine
COPY --from=builder /codecomp-backend ./
CMD ["./codecomp-backend"]
EXPOSE 8080
