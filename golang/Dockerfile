FROM golang:1.17.1 as build
RUN mkdir /app
WORKDIR /app
COPY go.mod .
COPY . .
RUN go build -o /app/main.out

FROM scratch as runtime
COPY --from=build /app/main.out main.out 
CMD ['./main.out']