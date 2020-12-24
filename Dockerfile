FROM golang:1.15.6 as builder
LABEL maintainer="Patrick Jusic <patrick.jusic@protonmail.com>"

WORKDIR /backup

ENV GOOS=linux \
  GO111MODULE=on

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o ./build/backup -v ./cmd/backup


FROM golang:alpine

WORKDIR /root/
RUN mkdir /root/.pg-backup
COPY --from=builder /backup/config ./config
COPY --from=builder /backup/build/backup .

CMD ["./backup", "export"]