#!/usr/bin/env bash

. 0-settings.sh

color $COLOR1 "Installing go..."
echo

eval_echo "[[ \"$SKIP_GO_INSTALL\" == \"true\" ]] && skip"
eval_echo "[[ \"$INSTALL_LOCALLY\" == \"false\" ]] && skip"
eval_echo "[[ \"$INSTALL_LOCALLY\" == \"\" ]] && skip"

ask_message $COLOR4 "There is a need to update the system and install gcc. Should we update (y/n)?"
if [[ "$ask_result" == "y" ]]; then
   eval_echo "sudo apt -y update"
   eval_echo "sudo apt -y install build-essential"
else
   color $COLOR4 "Skipped"
fi

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