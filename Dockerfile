FROM registry-adapter.tools.cosmic.sky/core-platform/core-community/baseimages/golang:1.21.2-alpine3.18 AS build

WORKDIR /app

COPY . ./

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /myapp


FROM chromedp/headless-shell:latest

RUN apt-get update; apt install dumb-init -y

ENTRYPOINT ["dumb-init", "--"]

COPY --from=build /myapp /tmp

CMD ["/tmp/myapp"]


# FROM chromedp/headless-shell:latest AS run

# WORKDIR /app

# COPY --from=build /myapp /app/myapp

# ENTRYPOINT ["/app/myapp"]
