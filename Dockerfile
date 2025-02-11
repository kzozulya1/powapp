# env for build
FROM golang:1.22.7 as builder

# setup working dir
WORKDIR /opt/pow

# cope app code into container fs
COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -mod=readonly -o bin/pow-server ./server

# env for run
FROM alpine:3.10

# Используем такую же рабочую директорию
WORKDIR /opt/pow

# Скопируем собранный бинарный код из первой ступени
COPY --from=builder /opt/pow/bin/pow-server .

CMD [ "/opt/pow/pow-server" ]