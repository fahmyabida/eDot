# Start from a golang base image
FROM golang:latest as builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
# RUN make build
RUN CGO_ENABLED=0 make build

FROM alpine:latest  

RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/brick-transfer .
CMD ["./brick-transfer user"]
