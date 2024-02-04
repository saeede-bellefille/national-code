FROM golang:1.21.6 
WORKDIR /app 
RUN go install github.com/swaggo/swag/cmd/swag@latest 
COPY go.mod go.sum ./ 
RUN go mod download 
COPY . . 
RUN go build -o server *.go
ENTRYPOINT /app/server