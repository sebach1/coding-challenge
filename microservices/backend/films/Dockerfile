FROM golang:alpine AS builder

ARG svc_name

RUN apk add git

LABEL maintainer="Sebastián Chamena <sebachamena@gmail.com>"

WORKDIR /${svc_name}

COPY go.* ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./migrate ./migrations
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./${svc_name} .

FROM scratch

ARG svc_name

WORKDIR /${svc_name}

COPY --from=builder /${svc_name}/${svc_name} ./${svc_name}
COPY --from=builder /${svc_name}/migrate ./migrate

COPY ./migrations/*.sql ./migrations/

CMD ["/films/films"]
