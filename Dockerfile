FROM golang:1.20

WORKDIR /app
ADD . /api
ADD . /docs
RUN touch .env

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY ./api ./api
COPY ./docs ./docs
RUN ls /app
# COPY /docs/ .

RUN go build -v -o ./atla-api ./api


CMD [ "./atla-api" ]