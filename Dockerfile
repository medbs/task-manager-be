FROM golang:latest
WORKDIR /app
COPY ./ /app
ENV GO111MODULE=on
RUN CGO_ENABLED=0 GOOS=linux go build -o tm /app/cmd
ENTRYPOINT ["./app/tm"]
