# syntax=docker/dockerfile:1
FROM golang:1.22-alpine
WORKDIR /go/src/github.com/wheresalice/invidious2newpipe/
COPY . .
#RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o invidious2newpipe .

FROM scratch
COPY --from=0 /go/src/github.com/wheresalice/invidious2newpipe/invidious2newpipe /
COPY Procfile /
CMD ["/invidious2newpipe"]
