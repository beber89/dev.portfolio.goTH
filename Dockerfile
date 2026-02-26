FROM gcr.io/distroless/base-debian12
WORKDIR /app
COPY server /app/server
COPY static /app/static
COPY locales /app/locales
COPY content /app/content
EXPOSE 8080
ENTRYPOINT ["/app/server"]
