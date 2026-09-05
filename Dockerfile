FROM debian:bookworm-slim

RUN apt-get update && \
    apt-get install -y --no-install-recommends ca-certificates tzdata && \
    rm -rf /var/lib/apt/lists/*

ARG APP=layout
ENV TZ=Asia/Shanghai

COPY ./build/release/${APP} /app/server
COPY ./config/config.yaml /app/config/config.yaml

WORKDIR /app

USER 65532:65532

ENTRYPOINT ["/app/server", "app"]
CMD ["--config", "/app/config/config.yaml"]
