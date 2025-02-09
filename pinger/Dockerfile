FROM golang:1.22

WORKDIR /container_monitoring
ENV TZ="Europe/Minsk"

COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN go build -o pinger

CMD ["/container_monitoring/pinger"]