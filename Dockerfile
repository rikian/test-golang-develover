FROM golang

WORKDIR /test-golang-developer

COPY . .

RUN go mod tidy
RUN go build -o ./test-golang-developer .

CMD ["./test-golang-developer"]