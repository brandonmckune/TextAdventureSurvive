
@echo off

SET GOPATH=%~dp0.

REM Check to make sure that our bin folder is in path for install
echo.%PATH%|findstr %~dp0./bin >nul 2>&1
if errorlevel 1 (
    SET PATH=%PATH%;%~dp0./bin
)

go install survive

for /R src %%f in (*.wad) do copy "%%f" bin

START /max bin/survive.exe