FROM alpine

WORKDIR /app
COPY build /app/
CMD ["/app/order_api-linux"]