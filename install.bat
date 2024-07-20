@echo off
echo Installing GoPanic, please stand by. If you get any errors, make sure you have Golang compiler installed and have it in the PATH directory.
go mod tidy
go build -o configure.exe                        .\cmd\configure\main.go
go build -o gopanic.exe --ldflags -H=windowsgui  .\cmd\gopanic\main.go  
if exist configure.exe (
    echo Programs compiled successfully!
) else (
    echo Something went wrong
    pause 
)
echo - - - - CONFIGURE - - - -
.\configure.exe
echo Configuration successful! Press Enter to exit.
pause