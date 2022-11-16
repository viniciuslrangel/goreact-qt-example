@echo on
set GO111MODULE=off
go install -v -tags=no_env github.com/therecipe/qt/cmd/...
set GO111MODULE=
REM go mod vendor
SET GOWORK=off
REM git clone --depth 1 --single-branch https://github.com/therecipe/env_windows_amd64_513.git vendor/github.com/therecipe/env_windows_amd64_513
set QT_VERSION=5.13.0
for /f %%v in ('go env GOPATH') do %%v\bin\qtsetup