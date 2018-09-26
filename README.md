# showcase

This is garbage please don't look

### installation

If you've got it, [Docker](https://www.docker.com/) is the best way to run this project.

    docker create \
      --name=showcase \
      --restart always \
      -e PGID=1000 -e PUID=1000  \
      -e TZ=America/Los_Angeles \
      -e USERNAME=<username> \
      -e PASSWORD=<password> \
      -p 8080:8080 \
      xanderstrike/showcase

* PGID and PUID can be found by running `id` in a terminal
* TZ is your timezone
* USERNAME and PASSWORD are optional authentication params
* Configure the port by setting `8080:8080` to `<your port>:8080`

Run with:

    docker start showcase

### license

MIT license
