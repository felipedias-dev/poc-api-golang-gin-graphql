FROM golang:1.21.3

LABEL maintener="felipedias.dev@gmail.com"

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

ENV PORT=5000

RUN go build
RUN find . -name "*.go" -type f -delete

EXPOSE $PORT

CMD [ "./poc-api-golang-gin-graphql" ]
