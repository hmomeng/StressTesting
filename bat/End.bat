@echo off
setlocal enabledelayedexpansion
echo Aida64如果出现未激活提示，请注意系统时间是否正确
echo.

echo ---------------------------
%~dp0\kktool\ghw.exe nowait
echo ---------------------------

echo.
echo 0.回车键关闭窗口  
echo 1.输入1  关闭测压程序，打开其他测试软件
echo 2.输入2  关闭测压程序，并删除桌面压测文件
set /p userInput=
if "%userInput%"=="1" (
    call :EndProcessesForDir
    cd kktool
    start test.bat
    goto End
)
if "%userInput%"=="2" (
    call :EndProcessesForDir
    call :DeleteFilesAndEndProcesses
    goto DelEnd
)
goto End

:EndProcessesForDir
for /f "tokens=2" %%i in ('tasklist /FI "IMAGENAME eq aida_bench64.dll" /FO TABLE /NH') do taskkill /F /PID %%i
rem 设置目标文件夹路径
set "target_folder=%~dp0kktool"
rem 遍历指定文件夹下的所有 .exe 文件，并结束其对应的进程
for /r "%target_folder%" %%i in (*.exe) do (
    rem 获取文件名（不含路径）
    set "filename=%%~nxi"
    rem 尝试结束与该文件名匹配的进程
    taskkill /f /im "%%~nxi" 2>nul
)
echo 完成进程结束操作
exit /b


:DeleteFilesAndEndProcesses
echo %~dp0
rem 设置要删除的目标文件夹路径
set "target_folder=%~dp0kktool"
rd /s /q "%target_folder%"
del startTest.exe
echo 已成功删除文件夹及其所有内容。
exit /b

REM 仅结束程序
:End
echo 所有操作完成
exit

REM 结束程序并删除End.bat
:DelEnd
set "target_folder=%~dp0kktool"
rd /s /q "%target_folder%"
echo 所有操作完成
del %0
pause


