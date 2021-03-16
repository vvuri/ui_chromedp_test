FROM golang:1.14-alpine AS build

WORKDIR /src/
COPY main.go go.* /src/
RUN apk add --no-cache git make musl-dev go
RUN go get github.com/chromedp/chromedp
RUN CGO_ENABLED=0 go build -o /bin/uitest

#FROM scratch
FROM alpine
COPY --from=build /bin/uitest /bin/uitest
RUN apk add chromium
ENTRYPOINT ["/bin/uitest"]
