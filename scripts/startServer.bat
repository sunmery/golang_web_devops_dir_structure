chcp 65001

@echo off
@REM 延迟扩充变量的功能
setlocal enabledelayedexpansion

@REM 选项
choice /c CP /t 5 /d C /m "请选择你当前的终端 C: CMD P: PowerShell"

@REM 1为第一个选项
if "!errorlevel!"=="1" (
    call :ChoiceA
) else (
    call :ChoiceB
)

@REM 运行结束退出运行改任务的终端窗口
:End
exit

@REM CMD的执行
:ChoiceA
echo 您选择了CMD。
START "run frontend server" cmd /k .\cmd\run_frontend_server.bat
START "run backend server" cmd /k .\cmd\run_backend_server.bat
goto :eof

@REM PowerShell的执行
:ChoiceB
echo 您选择了PowerShell。
START "run frontend server" "F:\Microsoft\PowerShell\7\pwsh.exe" -NoExit -Command .\cmd\run_frontend_server.bat
START "run frontend server" "F:\Microsoft\PowerShell\7\pwsh.exe" -NoExit -Command .\cmd\run_backend_server.bat
goto :eof
