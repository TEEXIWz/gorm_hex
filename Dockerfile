FROM golang

WORKDIR /app

ARG CACHEBUST=1

RUN git clone -b develop https://<TEEXIWz>:<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>@github.com/TEEXIWz/gorm_hex.git .

RUN go mod download

ENV GIN_MODE=release

COPY ./ ./

RUN go install github.com/cosmtrek/air@latest

# Don't forget to add .air.toml .gitignore
RUN air init

EXPOSE 8080

CMD ["air"]