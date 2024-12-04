@echo off
setlocal enabledelayedexpansion



copy "%~dp0End.bat" "%USERPROFILE%\Desktop\End.bat"


:: 定义 7z.exe 路径和需要解压的文件列表
set "sevenZipPath=%~dp07z.exe"
set "files=aida64.7z aida64Pro.7z Prime95.7z"

:: 遍历文件并解压
for %%f in (%files%) do (
    echo 正在解压文件: %%f...
    "%sevenZipPath%" x "%~dp0%%f" -aoa -o"%~dp0" >nul 2>&1
    if !errorlevel! equ 0 (
        echo 解压完成: %%f
    ) else (
        echo 解压失败: %%f
    )
)


start %~dp0Prime95\prime95.exe -t
start %~dp0AIDA64\AIDA64.exe /SST CPU,FPU,CACHE,RAM
start %~dp0cpuz.exe
start %USERPROFILE%\Desktop\End.bat
exit
