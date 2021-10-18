if "%GO_CLIENT_HOME%"=="" (
    call 0-settings.bat
)

echo off
echo [44;93m
echo        +-------------------------------------------------------------------------+
echo        !                   Now we are building go client...                      !
echo        +-------------------------------------------------------------------------+
echo [0m
echo on

call %GO% version

call :ask

goto :eof

:ask
    echo Press any key to continue
    pause >nul
goto :eof