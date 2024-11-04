# MossAPI

container-based dynamic API assembling system.
The architecture is not microservice. similar to modular monolithic. Image is like a modular synth(musical equipment), you can reassemble services, without stop the API server.

## Gateway

This program is reverse proxy server.
Check exist service and redirect to service container.

## Manager

Management conteiner and database tables.
You can make Dockerfile, docker container and database tables by command and web-based UI.
