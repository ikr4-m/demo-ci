FROM docker.io/library/golang:1.25-alpine AS backend
WORKDIR /build
RUN apk add --no-cache gcc musl-dev;

COPY ./go.* .
RUN go mod download -x;
COPY . .
RUN go build -C . -ldflags "-w -s -linkmode external -extldflags -static" -o app;


FROM scratch
COPY --from=backend /build/app /app
CMD ["/app"]
