if "%JAVA_CLIENT_HOME%"=="" (
    call 0-settings.bat
)

echo off
echo        [44;93m+-------------------------------------+[0m
echo        [44;93m!            Installing Go            ![0m
echo        [44;93m+-------------------------------------+[0m
echo on

if "%SKIP_GO_INSTALL%"=="true" ( goto :skip )
if "%INSTALL_LOCALLY%"=="false" ( goto :skip )
if "%INSTALL_LOCALLY%"=="" ( goto :skip )

cd %ROOT%
powershell -command "& { set-executionpolicy remotesigned -s currentuser; [System.Net.ServicePointManager]::SecurityProtocol = 3072 -bor 768 -bor 192 -bor 48; $client=New-Object System.Net.WebClient; $client.Headers.Add([System.Net.HttpRequestHeader]::Cookie, 'oraclelicense=accept-securebackup-cookie'); $client.DownloadFile('%ARCH_GO%','%TOOLS%\go.zip') }"
rd /S /Q %TOOLS%\..\.golang
%ARCH% x -y -o%TOOLS%\.. %TOOLS%\go.zip
rename %TOOLS%\..\%ARCH_GO_FOLDER% .golang
cd %ROOT%

call :ask

goto :eof

:skip
	echo off
	echo        [44;93m  Installation skipped:       [0m
	echo        [44;93m      INSTALL_LOCALLY=%INSTALL_LOCALLY%     [0m
	echo        [44;93m      SKIP_GO_INSTALL=%SKIP_GO_INSTALL%          [0m
	echo on
	goto :ask
goto :eof

:ask
    echo off
    echo        [44;93m+---------------------------------+[0m
    echo        [44;93m!    Press any key to continue    ![0m
    echo        [44;93m+---------------------------------+[0m
    echo on
    pause >nul
goto :eof