call 0-settings.bat

echo off
call lib.bat :color Installing go...
echo on

if "%SKIP_GO_INSTALL%"=="true" ( goto :skip )
if "%INSTALL_LOCALLY%"=="false" ( goto :skip )
if "%INSTALL_LOCALLY%"=="" ( goto :skip )

call lib.bat :install golang
call lib.bat :print_color %GO% version

call lib.bat :ask

goto :eof

:skip
	echo off
	call lib.bat :color Installation skipped
	call lib.bat :color INSTALL_LOCALLY=%INSTALL_LOCALLY%
	call lib.bat :color SKIP_GO_INSTALL=%SKIP_GO_INSTALL%
	echo on
	call lib.bat :ask
    goto :eof