FROM golang:1.16-alpine

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsufix -o bin/app cmd/main.go

FROM golang:1.16-alpine
COPY --from=build /go/src/app/bin/app .
CMD [ "./app", "-env=false" ]
