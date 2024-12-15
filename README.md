# showcase

Stupid simple tabbed homepage for sites with lots of services

## Docker Run

Run the following command to start the container:

```
docker run -d \
  -v <path to config>:/app/config \
  -p 8080:8080 \
  xanderstrike/showcase:latest
```

Or use the provided [docker-compose.yml](https://raw.githubusercontent.com/xanderstrike/showcase/master/docker-compose.yml) file.
