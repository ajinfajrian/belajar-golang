# syntax=docker/dockerfile:1

FROM golang:1.17-alpine
FROM busybox

FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /web-container

EXPOSE 8080

CMD [ "/web-container" ]

#WORKDIR /app
#COPY go.mod go.sum ./
#RUN go mod download
#COPY *.go ./
#RUN go install github.com/gofiber/fiber/v2
#RUN go build -o /fiber.go
#COPY --from=build /fiber /app/fiber
#EXPOSE 8080
#ENTRYPOINT [ "/fiber" ]
