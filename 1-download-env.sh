#!/usr/bin/env bash

. 0-settings.sh

color $COLOR1 "Installing go..."
echo

eval_echo "[[ \"$SKIP_GO_INSTALL\" == \"true\" ]] && skip"
eval_echo "[[ \"$INSTALL_LOCALLY\" == \"false\" ]] && skip"
eval_echo "[[ \"$INSTALL_LOCALLY\" == \"\" ]] && skip"


eval_echo "sudo apt update"
eval_echo "sudo apt install build-essential"

eval_echo "install 'golang' '$ARCH_URL' '$ARCH_FOLDER'"
eval_echo_color_output "$GO version"

ask

skip() {
    color $COLOR3 "Installation skipped"
    color $COLOR3 "INSTALL_LOCALLY=$INSTALL_LOCALLY"
    color $COLOR3 "SKIP_GO_INSTALL=$SKIP_GO_INSTALL"
    ask
    exit
}