@echo off
setlocal enabledelayedexpansion



copy "%~dp0End.bat" "%USERPROFILE%\Desktop\End.bat"


:: ���� 7z.exe ·������Ҫ��ѹ���ļ��б�
set "sevenZipPath=%~dp07z.exe"
set "files=aida64.7z aida64Pro.7z Prime95.7z"

:: �����ļ�����ѹ
for %%f in (%files%) do (
    echo ���ڽ�ѹ�ļ�: %%f...
    "%sevenZipPath%" x "%~dp0%%f" -aoa -o"%~dp0" >nul 2>&1
    if !errorlevel! equ 0 (
        echo ��ѹ���: %%f
    ) else (
        echo ��ѹʧ��: %%f
    )
)


start %~dp0Prime95\prime95.exe -t
start %~dp0AIDA64\AIDA64.exe /SST CPU,FPU,CACHE,RAM
start %~dp0cpuz.exe
start %USERPROFILE%\Desktop\End.bat
exit
