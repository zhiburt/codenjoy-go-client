call 0-settings.bat

echo off
call lib.bat :color Running go client...
echo on

call %GO% run main.go %GAME_TO_RUN% %BOARD_URL%

call lib.bat :ask