#########
# Build #
#########
FROM golang:alpine3.17 AS build

RUN mkdir /dummy-api
WORKDIR /dummy-api

COPY ../../src /dummy-api

RUN go mod tidy
RUN go build -o api .

##########
# Deploy #
##########
FROM golang:alpine3.17 AS serve

WORKDIR /service

COPY --from=build /dummy-api/api ./api

CMD ["/service/api"]