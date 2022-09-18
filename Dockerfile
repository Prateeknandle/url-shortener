FROM golang:latest

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

#COPY *.go ./
COPY . ./

RUN ls

RUN go build -o /docker-gs-ping .

EXPOSE 3000

CMD [ "/docker-gs-ping" ]