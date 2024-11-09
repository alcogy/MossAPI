# MossAPI

container-based dynamic API assembling system.
The architecture is not microservice. similar to modular monolithic. Image is like a modular synth(musical equipment), you can reassemble services, without stop the API server.

## Gateway

This program is mainly as a reverse proxy server. All services (containers) running in private network.
Therefore, need to access to service through gateway.

### How to run

1. Setting connect to your MySQL conf to /gateway/docker/.env

2. Create networks for public and private

```shell
docker network create --internal mossapi-nw-private
docker network create mossapi-nw-public
```

3. Build image and create container.

```shell
docker build -t gateway <path to Dockerfile>
docker run --name gateway -p 9000:9000 --network mossapi-nw-private --network mossapi-nw-public -d gateway
```

## Manager

Management conteiner and database tables.
You can simply make Dockerfile, docker container and database tables.
I provide web-based UI admin.

### How to use.

On dvelopment.

```shell
./go run . admin
```

On product.

```shell
./manager run . admin
```
