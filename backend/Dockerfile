FROM golang:latest
ENV GOPATH /backend
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY *.go ./
RUN go build -o /backend
EXPOSE 5000
CMD [ "/backend" ]