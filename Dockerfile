FROM iron/go
WORKDIR /app
VOLUME /app/config/
ADD showcase-docker /app/
COPY index.html /app/index.html
EXPOSE 8080
ENTRYPOINT ["./showcase-docker"]
