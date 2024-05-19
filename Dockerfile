FROM golang:1.22

WORKDIR /app
COPY . /app
RUN go build /app

ENV ALLOWED_ORIGIN="*"
ENV PORT="8080"
ENV HEADERS="Content-Type, Authorization"
ENV DSN="root:mysql@1SI18CS096@tcp(localhost:3306)/GOTASKS?charset=utf8mb4&parseTime=True&loc=Local"

EXPOSE 8080
ENTRYPOINT [ "./gotasks" ]