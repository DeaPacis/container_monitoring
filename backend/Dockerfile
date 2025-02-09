FROM golang:1.22

WORKDIR /container_monitoring
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN go build -o backend

CMD ["/container_monitoring/backend"]