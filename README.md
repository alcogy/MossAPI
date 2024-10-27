# Modular Synthesis API

Docker Container based dynamic API system.
The architecture is not microservice and monolithic, similar to modular monolithic. Image is like a modular synth(musical equipment), you can reassemble services, without stop the API server.

## Gateway

This program is reverse proxy server.
Check exist service and redirect to service container.

## Manager

Management conteiner.
Make Dockerfile, docker image and run container.
