@echo off
setlocal enabledelayedexpansion
echo Aida64�������δ������ʾ����ע��ϵͳʱ���Ƿ���ȷ
echo.

echo ---------------------------
%~dp0\kktool\ghw.exe nowait
echo ---------------------------

echo.
echo 0.�س����رմ���  
echo 1.����1  �رղ�ѹ���򣬴������������
echo 2.����2  �رղ�ѹ���򣬲�ɾ������ѹ���ļ�
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
rem ����Ŀ���ļ���·��
set "target_folder=%~dp0kktool"
rem ����ָ���ļ����µ����� .exe �ļ������������Ӧ�Ľ���
for /r "%target_folder%" %%i in (*.exe) do (
    rem ��ȡ�ļ���������·����
    set "filename=%%~nxi"
    rem ���Խ�������ļ���ƥ��Ľ���
    taskkill /f /im "%%~nxi" 2>nul
)
echo ��ɽ��̽�������
exit /b


:DeleteFilesAndEndProcesses
echo %~dp0
rem ����Ҫɾ����Ŀ���ļ���·��
set "target_folder=%~dp0kktool"
rd /s /q "%target_folder%"
del startTest.exe
echo �ѳɹ�ɾ���ļ��м����������ݡ�
exit /b

REM ����������
:End
echo ���в������
exit

REM ��������ɾ��End.bat
:DelEnd
set "target_folder=%~dp0kktool"
rd /s /q "%target_folder%"
echo ���в������
del %0
pause


