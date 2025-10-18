FROM alpine:3.20

WORKDIR /app

COPY data-plane                                     .
RUN mkdir services
COPY services/killarney_sectors_v4_annotated.json   services/killarney_sectors_v4_annotated.json

EXPOSE 8080
CMD ["./data-plane"]