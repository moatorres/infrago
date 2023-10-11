# stage 1
FROM golang:1.21.3-alpine3.18 AS base

RUN apk add --no-cache tini

WORKDIR /app

ENTRYPOINT [ "/sbin/tini", "--" ]

# stage 2
FROM base AS builder

COPY go.mod .

RUN go mod download

COPY . .

RUN go build -o main

# stage 3
FROM base AS runner

WORKDIR /

RUN adduser -D -s /bin/ash nonroot

WORKDIR /home/nonroot/app

COPY --chown=nonroot:nonroot --from=builder /app/static ./static
COPY --chown=nonroot:nonroot --from=builder /app/main ./

USER nonroot

EXPOSE 3000

CMD [ "/home/nonroot/app/main" ]
