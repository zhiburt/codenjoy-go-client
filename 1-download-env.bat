if "%JAVA_CLIENT_HOME%"=="" (
    call 0-settings.bat
)

echo off
echo [44;93m
echo        +-------------------------------------+
echo        !            Installing Go            !
echo        +-------------------------------------+
echo [0m
echo on

cd %ROOT%
rd /S /Q %TOOLS%\..\.golang
powershell -command "& { set-executionpolicy remotesigned -s currentuser; [System.Net.ServicePointManager]::SecurityProtocol = 3072 -bor 768 -bor 192 -bor 48; $client=New-Object System.Net.WebClient; $client.Headers.Add([System.Net.HttpRequestHeader]::Cookie, 'oraclelicense=accept-securebackup-cookie'); $client.DownloadFile('%ARCH_GO%','%TOOLS%\go.zip') }"
%ARCH% x -y -o%TOOLS%\..\.golang %TOOLS%\go.zip
cd %ROOT%

call :ask

goto :eof

:ask
    echo Press any key to continue
    pause >nul
goto :eof
