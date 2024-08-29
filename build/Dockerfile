FROM golang:1.23 AS build
RUN apt-get update && \
    apt-get upgrade -y
WORKDIR /go/src/tour
COPY . .

RUN GOOS=linux go build -a -installsuffix cgo -ldflags "-s -w" /go/src/tour/server/main.go && \
    mkdir /app && \
	mv main /app/server && \
	mv static /app && \
    mv templates /app

FROM ubuntu
RUN apt update && \
    apt upgrade -y && \
    apt autoremove -y && \
    apt install -y net-tools tzdata ca-certificates
COPY --from=build /app /app
WORKDIR /app
EXPOSE 8080
ENTRYPOINT ["/app/server"]
