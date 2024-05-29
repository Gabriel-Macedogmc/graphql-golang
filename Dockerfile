FROM esolang/sqlite3:2.4.0

WORKDIR /app

VOLUME [ "/cmd/server/" ]

CMD [ "sh" ]