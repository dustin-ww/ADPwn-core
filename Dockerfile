# Build-Stage
FROM golang:1.24.1-alpine AS builder

ARG INSTALL_TOOLS="default"

WORKDIR /app

RUN apk add --no-cache git

COPY config/ /app/config/

RUN if [ -f "/app/config/tools.json" ]; then \
        apk add --no-cache jq && \
        TOOLS_TO_INSTALL=$(jq -r '.tools | join(" ")' /app/config/tools.json 2>/dev/null || echo "${INSTALL_TOOLS}"); \
    else \
        TOOLS_TO_INSTALL="${INSTALL_TOOLS}"; \
    fi && \
    apk add --no-cache ca-certificates && \
    if [ "$TOOLS_TO_INSTALL" != "default" ] && [ -n "$TOOLS_TO_INSTALL" ]; then \
        apk add --no-cache $TOOLS_TO_INSTALL; \
    fi && \
    if [ "$TOOLS_TO_INSTALL" = "default" ] || echo "$TOOLS_TO_INSTALL" | grep -q "nmap"; then \
        apk add --no-cache nmap nmap-scripts; \
    fi

COPY src/ /app/src/

WORKDIR /app/src

RUN go work sync && go mod download

WORKDIR /app
COPY . .

RUN cd src && CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /go/bin/app ./rest

FROM alpine:latest

ARG INSTALL_TOOLS="default"

# Wichtig: Erstelle das /config Verzeichnis und kopiere Dateien dorthin statt nach /root/config
RUN mkdir -p /config
COPY config/ /config/

RUN if [ -f "/config/tools.json" ]; then \
        apk add --no-cache jq && \
        TOOLS_TO_INSTALL=$(jq -r '.tools | join(" ")' /config/tools.json 2>/dev/null || echo "${INSTALL_TOOLS}"); \
    else \
        TOOLS_TO_INSTALL="${INSTALL_TOOLS}"; \
    fi && \
    apk add --no-cache ca-certificates && \
    if [ "$TOOLS_TO_INSTALL" != "default" ] && [ -n "$TOOLS_TO_INSTALL" ]; then \
        apk add --no-cache $TOOLS_TO_INSTALL; \
    fi && \
    if [ "$TOOLS_TO_INSTALL" = "default" ] || echo "$TOOLS_TO_INSTALL" | grep -q "nmap"; then \
        apk add --no-cache nmap nmap-scripts; \
    fi

WORKDIR /root/

COPY --from=builder /go/bin/app .

EXPOSE 8081
EXPOSE 8082

CMD ["./app"]