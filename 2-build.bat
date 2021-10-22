call 0-settings.bat

echo off
call lib :color Building go client...
echo on

call lib :print_color %GO% version

call lib :ask