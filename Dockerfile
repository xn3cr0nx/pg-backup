FROM golang:1.15.6 as builder
LABEL maintainer="Patrick Jusic <patrick.jusic@protonmail.com>"

WORKDIR /pg-backup

ENV GOOS=linux \
  GO111MODULE=on

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o ./build/pg-backup -v ./cmd/backup


FROM golang:alpine

WORKDIR /root/
RUN mkdir /root/.pg-backup
RUN mkdir /root/backups
COPY --from=builder /pg-backup/config ./config
COPY --from=builder /pg-backup/build/pg-backup .

CMD ["./pg-backup", "start"]