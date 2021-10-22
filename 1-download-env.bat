call 0-settings.bat

echo off
call lib :color Installing go...
echo on

if "%SKIP_GO_INSTALL%"=="true" ( goto :skip )
if "%INSTALL_LOCALLY%"=="false" ( goto :skip )
if "%INSTALL_LOCALLY%"=="" ( goto :skip )

call lib :install golang %ARCH_URL% %ARCH_FOLDER%
call lib :print_color %GO% version

call lib :ask

goto :eof

:skip
	echo off
	call lib :color Installation skipped
	call lib :color INSTALL_LOCALLY=%INSTALL_LOCALLY%
	call lib :color SKIP_GO_INSTALL=%SKIP_GO_INSTALL%
	echo on
	call lib :ask
    goto :eof