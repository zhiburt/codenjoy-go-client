call 0-settings.bat

echo off
call lib.bat :color Building go client...
echo on

call lib.bat :print_color %GO% version

call lib.bat :ask