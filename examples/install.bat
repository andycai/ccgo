@echo off

setlocal

if exist install.bat goto ok
echo install.bat must be run from its folder
goto end

:ok

if "%1" neq "" goto doit
echo install.bat project-name
goto end

:doit

set OLDGOPATH=%GOPATH%
set GOPATH=%~dp0

gofmt -tabs=true -tabwidth=4 -w src

go install %1

:end
echo finished