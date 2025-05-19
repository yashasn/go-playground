FROM golang:1.24-alpine AS build
RUN apk update && \
    apk add make 

WORKDIR /app

# copy module files 
COPY go.* ./
RUN go mod download
RUN go mod verify

# Copy source code
COPY . .
RUN make build


FROM alpine:latest

WORKDIR /app

COPY --from=build /app/gameserver .

RUN ls -la

ENV PORT 8080 
EXPOSE 8080

CMD ["./gameserver"]