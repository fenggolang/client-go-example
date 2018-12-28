FROM debian
COPY ./app /app
USER 1001
ENTRYPOINT /app