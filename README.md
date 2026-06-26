# jhekasoft-api

`jhekasoft-api` is a backend for https://jhekasoft.github.io.


```
Powered by:
░█▀▀░█▀▄░█▀█░█▀▀░█░█░█▀▀░█▀█░█▀▄
░█▀▀░█▀▄░█▀█░█░░░█▀▄░█▀▀░█░█░█░█
░▀▀▀░▀▀░░▀░▀░▀▀▀░▀░▀░▀▀▀░▀░▀░▀▀░
```

[e-backend](https://github.com/jhekasoft/e-backend)

## Prepare

```bash
cp config.yaml.example config.yaml
```

And then edit `config.yaml` file.

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

## Run with docker

Build image:

```bash
docker build -f Dockerfile -t jhekasoft-api .
```

Run:

```bash
docker run --name jhekasoft-api --rm --network host \
-v "$(pwd)/config.yaml:/app/config.yaml" \
jhekasoft-api
```
