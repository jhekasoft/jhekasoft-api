# jhekasoft-api

`jhekasoft-api` is a backend for all the projects.

![cat](./modules/doc/data/public/android-chrome-192x192.png)

```
▗▄▄▄▖▗▄▄▖  ▗▄▖  ▗▄▄▖▗▖ ▗▖▗▄▄▄▖▗▖  ▗▖▗▄▄▄ 
▐▌   ▐▌ ▐▌▐▌ ▐▌▐▌   ▐▌▗▞▘▐▌   ▐▛▚▖▐▌▐▌  █
▐▛▀▀▘▐▛▀▚▖▐▛▀▜▌▐▌   ▐▛▚▖ ▐▛▀▀▘▐▌ ▝▜▌▐▌  █
▐▙▄▄▖▐▙▄▞▘▐▌ ▐▌▝▚▄▄▖▐▌ ▐▌▐▙▄▄▖▐▌  ▐▌▐▙▄▄▀
```

## Create database

```bash
sudo -iu postgres
createdb ebackend
```

## Prepare

```bash
cp .config.example .config
```

And then edit `.config` file.

## Run HTTP-server

```bash
make run
```

## Building

Build binary:

```bash
make build
```

Clean:

```bash
make clean
```

Run binary:

```bash
./build/jhekasoft-api serve
```

## Run as service (POSIX systems with systemd)

```bash
sudo mkdir /opt/jhekasoft-api
sudo cp ./build/* /opt/jhekasoft-api -r
sudo cp /opt/jhekasoft-api/.jhekasoft-api.example /opt/jhekasoft-api/.jhekasoft-api
sudo cp ./systemd/jhekasoft-api.service /etc/systemd/system/
sudo systemctl daemon-reload
sudo systemctl enable --now jhekasoft-api.service
```

## Module generation

See [e-backend-cli](https://github.com/jhekasoft/e-backend-cli).

# Run with docker

Build image:

```bash
docker build -f dockerfiles/Dockerfile -t jhekasoft-api .
```

Run:

```bash
docker run --name jhekasoft-api --rm --network host \
-v "$(pwd)/.jhekasoft-api:/app/.jhekasoft-api" \
jhekasoft-api
```
