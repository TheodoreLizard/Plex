#!/bin/sh

# IP Address of Plex server.
export PLEX_ADDR=10.0.1.8

./plexerCLI playlist $@
