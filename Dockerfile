FROM golang:1.19.3-buster

WORKDIR /app
COPY . /app
RUN go mod tidy
RUN go get github.com/cosmtrek/air
CMD ["air"]