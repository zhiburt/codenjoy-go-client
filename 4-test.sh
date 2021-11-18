#!/usr/bin/env bash

. 0-settings.sh

color $COLOR1 "Starting go tests..."
echo

eval_echo "$GO test ./..."

ask