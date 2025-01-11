FROM golang:1.22

LABEL version="1.0"
LABEL description="A forum application where users can post, view their posts, liked posts, and comment on posts."

WORKDIR /forum

COPY . .

RUN go mod download

CMD ["go", "run", "cmd/main.go"]