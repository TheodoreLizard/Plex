#!/bin/sh

# IP Address of Plex server.
export PLEX_ADDR=10.0.1.8
export PLEX_TOKEN=Vq71cM1Lewqa6RRDB6UV

# export PLEX_ADDR=10.0.1.7
# export PLEX_TOKEN=swzgnJxeztjyes5mJ5EY

./plexerCLI playlist $@
