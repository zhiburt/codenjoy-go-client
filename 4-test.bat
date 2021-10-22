call 0-settings.bat

echo off
call lib :color Starting go tests...
echo on

call %GO% test ./tests/...

call lib :ask