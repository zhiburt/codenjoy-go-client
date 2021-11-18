#!/usr/bin/env bash

. 0-settings.sh

color $COLOR1 "Building go client..."
echo

eval_echo_color_output "$GO version"

ask
