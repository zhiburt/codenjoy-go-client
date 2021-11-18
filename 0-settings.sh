#!/usr/bin/env bash

. lib.sh

color $COLOR1 "Setup variables..."
echo

eval_echo "[[ \"$GAME_TO_RUN\" == \"\" ]] && GAME_TO_RUN=mollymage"
eval_echo "[[ \"$BOARD_URL\" == \"\" ]]   && BOARD_URL=http://127.0.0.1:8080/codenjoy-contest/board/player/0?code=000000000000"

eval_echo "ROOT=$PWD"

eval_echo "[[ \"$SKIP_TESTS\" == \"\" ]]  && SKIP_TESTS=true"

eval_echo "TOOLS=$ROOT/.tools"
eval_echo "ARCH=tar"

# Set to true if you want to ignore jdk and maven installation on the system
eval_echo "[[ \"$INSTALL_LOCALLY\" == \"\" ]] && INSTALL_LOCALLY=true"

eval_echo "[[ \"$INSTALL_LOCALLY\" == "true" ]] && GOPATH="

eval_echo "[[ \"$GOPATH\" == \"\" ]]  && GOPATH=$ROOT/.golang"

eval_echo "GO=$GOPATH/bin/go"
eval_echo "export PATH=\"$GOPATH/bin:$PATH\""

color $COLOR4 "GOPATH=$GOPATH"
echo

eval_echo "ARCH_URL=https://dl.google.com/go/go1.17.3.linux-amd64.tar.gz"
eval_echo "ARCH_FOLDER=go"

eval_echo "GO_CLIENT_HOME=$ROOT"