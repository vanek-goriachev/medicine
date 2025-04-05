FROM golang:1.24.1 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /medicine ./cmd/medicine/main.go


FROM scratch AS release

COPY --from=builder /medicine /medicine

CMD ["/medicine"]
