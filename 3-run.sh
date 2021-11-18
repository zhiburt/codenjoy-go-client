#!/usr/bin/env bash

. 0-settings.sh

color $COLOR1 "Starting go client..."
echo

eval_echo "$GO run main.go $GAME_TO_RUN $BOARD_URL"

ask
