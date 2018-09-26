# manual-upload

A profoundly simple approach to adding torrents to seedboxes.

While I might use CouchPotato, [GazelleUI](https://github.com/XanderStrike/GazelleUI), and SickRage for most of my torrents, I'll occasionally want to manually add specific ones. My previous workflow of scping them into the correct folder was getting tired, and I wanted to learn Go, so here we are.

This tool makes the assumption that you have a watch folder with the following directories: `movies`, `music`, `tv`, and `misc`. That's how I organize it, and that ought to be good enough for everybody ;)

### installation

If you've got it, [Docker](https://www.docker.com/) is the best way to run this project.

    docker create \
      --name=manual-upload \
      --restart always \
      -v <path to watchfolder>:/watch \
      -e PGID=1000 -e PUID=1000  \
      -e TZ=America/Los_Angeles \
      -e DISCORD_URL=<discord webhook url> \
      -e USERNAME=<username> \
      -e PASSWORD=<password> \
      -p 8080:8080 \
      xanderstrike/manual-upload

* Set the watchfolder to a directory watched by your torrent client
* PGID and PUID can be found by running `id` in a terminal
* TZ is your timezone
* DISCORD_URL is an optional parameter to enable discord notifications of new uploads
* USERNAME and PASSWORD are optional authentication params
* Configure the port by setting `8080:8080` to `<your port>:8080`

Run with:

    docker start manual-upload

### license

MIT license
