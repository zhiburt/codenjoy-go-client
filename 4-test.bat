call 0-settings.bat

echo off
call lib.bat :color Starting go tests...
echo on

call %GO% test ./tests/...

call lib.bat :ask