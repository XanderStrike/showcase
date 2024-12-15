# showcase

Stupid simple tabbed homepage for sites with lots of services.

Ideal for selfhosters, seedboxes, etc.

## Docker Run

Create a showcase-config.yml file and fill it in like this:

```yaml
title: My Dashboard
announcement: Welcome to my dashboard!
username: admin  # Optional: remove or leave blank to disable auth
password: secret # Optional: remove or leave blank to disable auth
tabs:
  - name: Home
    url: https://example.com
  - name: Status
    url: https://status.example.com
  - name: Monitoring
    url: https://grafana.example.com
```

Create a docker-compose.yml file and fill it in like this:

```yaml
version: '3.8'

services:
  showcase:
    image: xanderstrike/showcase:latest
    ports:
      - "8080:8080"
    volumes:
      - ./showcase-config.yml:/app/config/config.yml
    restart: unless-stopped
``` 

Run the following command to start the container:

```
docker compose up -d
```

## Configuration

The config.yml file supports the following options:

- `title`: The title shown in the browser tab
- `announcement`: A message shown in the bottom right
- `username`: Optional basic auth username
- `password`: Optional basic auth password
- `tabs`: A list of links to show in the menu
  - `name`: The text shown in the menu
  - `url`: Where the link should go

Authentication is disabled by default. To enable it, add both username and password to your config.
