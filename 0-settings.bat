echo off
call lib :color Setup variables...
echo on

if "%GAME_TO_RUN%"=="" ( set GAME_TO_RUN=mollymage)
if "%BOARD_URL%"==""   ( set BOARD_URL=http://127.0.0.1:8080/codenjoy-contest/board/player/0?code=000000000000)

set ROOT=%CD%

if "%SKIP_TESTS%"=="" ( set SKIP_TESTS=true)

set CODE_PAGE=65001
chcp %CODE_PAGE%

set TOOLS=%ROOT%\.tools
set ARCH=%TOOLS%\7z\7za.exe

rem Set to true if you want to ignore go installation on the system
if "%INSTALL_LOCALLY%"=="" ( set INSTALL_LOCALLY=true)

if "%INSTALL_LOCALLY%"=="true" ( set GOPATH=)
if "%GOPATH%"=="" ( set GOPATH=%ROOT%\.golang)

set GO=%GOPATH%\bin\go

echo off
call lib :color GOPATH=%GOPATH%
echo on

set ARCH_URL=https://golang.org/dl/go1.16.5.windows-amd64.zip
set ARCH_FOLDER=go

set GO_CLIENT_HOME=%ROOT%