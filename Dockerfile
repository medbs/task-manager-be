FROM golang:latest
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

COPY . .
ENV GO111MODULE=on
EXPOSE 8099

RUN CGO_ENABLED=0 GOOS=linux go build -o tm /app/cmd

CMD ["./tm"]
