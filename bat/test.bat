echo=1/*>nul&@cls
@echo off
title ѹ�����
color 17
mode con cols=30 lines=15
if exist %~d0Test_Tool (
goto menu
) else (
md "%~d0Test_Tool"
call :http "http://yourdomain.com/test/axel.exe" "%~d0Test_Tool\axel.exe" >nul 2>nul
call :http "http://yourdomain.com/test/cygwin1.dll" "%~d0Test_Tool\cygwin1.dll" >nul 2>nul
call :http "http://yourdomain.com/test/7z.exe" "%~d0Test_Tool\7z.exe" >nul 2>nul
call :http "http://yourdomain.com/test/7z.dll" "%~d0Test_Tool\7z.dll" >nul 2>nul
REM echo ��������ѹ�����......
REM .\Test_Tool\axel.exe -n 10 -o .\Test_Tool "http://yourdomain.com/test/test_tool.7z"
REM .\Test_Tool\7z.exe x .\Test_Tool\test_tool.7z -o.\Test_Tool -aoa
)
:start
:menu
echo �T�T�T�T�T�T�T�T�T�T�T�T�T�T�T�T�T�T�T�T�T�T�T�T�T�T�T�T�T�T
echo     Microsoft ѹ������ϼ�
echo �T�T�T�T�T�T�T�T�T�T�T�T�T�T�T�T�T�T�T�T�T�T�T�T�T�T�T�T�T�T
echo    1.����AIDA64
echo    2.����CPU-Z
echo    3.����Prime95
echo    4.�������ܲ���
echo    5.��װGeekbench
echo    6.����Geekbench
echo    7.΢�������л���
echo    8.Cinebench R23
echo    9.Cinebench R20-08
echo    10.����³��ʦ
echo    11.ʱ��ͬ������
echo �T�T�T�T�T�T�T�T�T�T�T�T�T�T�T�T�T�T�T�T�T�T�T�T�T�T�T�T�T�T
set /p n=ѡ����Ҫ���еĳ���
echo �T�T�T�T�T�T�T�T�T�T�T�T�T�T�T�T�T�T�T�T�T�T�T�T�T�T�T�T�T�T
if "%n%"=="1" call :AIDA64
if "%n%"=="2" call :cpuz
if "%n%"=="3" call :Prime95
if "%n%"=="4" call :disk
if "%n%"=="5" call :geekbench-1
if "%n%"=="6" call :geekbench-2
if "%n%"=="7" call :c++
if "%n%"=="8" call :cinebench-12
if "%n%"=="9" call :Cinebench-08
if "%n%"=="10" call :ludashi
if "%n%"=="11" call :time
if /i "%n%"=="e" exit 
:cuowu
echo �������
pause
goto menu

:Prime95
if exist %~d0Test_Tool\Prime95\prime95.exe (
start %~d0Test_Tool\Prime95\prime95.exe
) else (
echo ��������prime95......
.\Test_Tool\axel.exe -n 10 -o .\Test_Tool "http://yourdomain.com/test/p95v3019b20.win64.zip"
.\Test_Tool\7z.exe x .\Test_Tool\p95v3019b20.win64.zip -o.\Test_Tool\p95v3019b20.win64 -aoa
start %~d0Test_Tool\p95v3019b20.win64\prime95.exe
)
goto menu
REM start %~d0Test_Tool\Prime95\prime95.exe
REM goto menu

:AIDA64
if exist %~d0Test_Tool\aida64\aida64.exe (
start %~d0Test_Tool\aida64\aida64.exe
) else (
echo ��������aida64......
.\Test_Tool\axel.exe -n 10 -o .\Test_Tool "http://yourdomain.com/test/aida64.rar"
.\Test_Tool\7z.exe x .\Test_Tool\aida64.rar -o.\Test_Tool\aida64 -aoa
start %~d0Test_Tool\aida64\aida64.exe
)
goto menu
REM start %~d0Test_Tool\AIDA64\aida64.exe
REM goto menu

:ludashi
if exist %~d0Test_Tool\ludashi (
start %~d0Test_Tool\ludashi\ComputerZ_CN.exe
) else (
echo ��������³��ʦ......
.\Test_Tool\axel.exe -n 10 -o .\Test_Tool "http://yourdomain.com/test/ludashi.7z"
.\Test_Tool\7z.exe x .\Test_Tool\ludashi.7z -o.\Test_Tool -aoa
start %~d0Test_Tool\ludashi\ComputerZ_CN.exe
)
goto menu

:cpuz
if exist %~d0Test_Tool\cpu-z_2.10-cn\cpuz_x64.exe (
start %~d0Test_Tool\cpu-z_2.10-cn\cpuz_x64.exe
) else (
echo ��������cpu-z......
.\Test_Tool\axel.exe -n 10 -o .\Test_Tool "http://yourdomain.com/test/cpu-z_2.10-cn.zip"
.\Test_Tool\7z.exe x .\Test_Tool\cpu-z_2.10-cn.zip -o.\Test_Tool\cpu-z_2.10-cn -aoa
start %~d0Test_Tool\cpu-z_2.10-cn\cpuz_x64.exe
)
goto menu
REM start %CD%\Test_Tool\cpu-z\cpuz_x64.exe
REM goto menu


:disk
if exist %~d0Test_Tool\CrystalDiskMark\DiskMark64.exe (
start %~d0Test_Tool\CrystalDiskMark\DiskMark64.exe
) else (
echo ��������CrystalDiskMark......
.\Test_Tool\axel.exe -n 10 -o .\Test_Tool "http://yourdomain.com/test/CrystalDiskMark.7z"
.\Test_Tool\7z.exe x .\Test_Tool\CrystalDiskMark.7z -o.\Test_Tool -aoa
start %~d0Test_Tool\CrystalDiskMark\DiskMark64.exe
)
goto menu

:time
start %~d0Test_Tool\timeup.exe
goto menu

:geekbench-1
if exist "C:\Program Files (x86)\Geekbench 5\Geekbench 5.exe" (
start "C:\Program Files (x86)\Geekbench 5\Geekbench 5.exe"
) else (
echo ��������Geekbench Pro
.\Test_Tool\axel.exe -n 10 -o .\Test_Tool "http://yourdomain.com/test/Geekbench.7z"
.\Test_Tool\7z.exe x .\Test_Tool\Geekbench.7z -o.\Test_Tool -aoa
echo ���ڰ�װGeekbench Pro
start /wait %~d0Test_Tool\Geekbench-5.4.4-WindowsSetup.exe /S
XCOPY "%~d0Test_Tool\Crack\Patch.exe" "C:\Program Files (x86)\Geekbench 5\" /S /E /Y
start "" "C:\Program Files (x86)\Geekbench 5\Patch.exe"
echo Geekbench Pro��װ���
echo �밴������ص����˵�
pause
)
goto menu

:geekbench-2
start "" "C:\Program Files (x86)\Geekbench 5\Geekbench 5.exe"
goto menu

:Cinebench-12
if exist %~d0Test_Tool\CinebenchR23\Cinebench.exe (
start %~d0Test_Tool\CinebenchR23\Cinebench.exe
) else (
echo ��������CinebenchR23......
.\Test_Tool\axel.exe -n 10 -o .\Test_Tool "http://yourdomain.com/test/CinebenchR23.7z"
.\Test_Tool\7z.exe x .\Test_Tool\CinebenchR23.7z -o.\Test_Tool\CinebenchR23 -aoa
start %~d0Test_Tool\CinebenchR23\Cinebench.exe
)
goto menu

:Cinebench-08
if exist %~d0Test_Tool\CinebenchR20\Cinebench.exe (
start %~d0Test_Tool\CinebenchR20\Cinebench.exe
) else (
echo ��������CinebenchR20......
.\Test_Tool\axel.exe -n 10 -o .\Test_Tool "http://yourdomain.com/test/CinebenchR20.7z"
.\Test_Tool\7z.exe x .\Test_Tool\CinebenchR20.7z -o.\Test_Tool\CinebenchR20 -aoa
start %~d0Test_Tool\CinebenchR20\Cinebench.exe
)
goto menu

:c++
if exist %~d0Test_Tool\vcdaquan.exe (
start %~d0Test_Tool\vcdaquan.exe
) else (
echo ��������vcdaquan......
.\Test_Tool\axel.exe -n 10 -o .\Test_Tool "http://yourdomain.com/test/vcdaquan.exe"
REM start %~d0Test_Tool\vcdaquan.exe
REM start %~d0Test_Tool\vcdaquan.exe /install /silent /norestart
start /wait %~d0Test_Tool\vcdaquan.exe  /VERYSILENT
)
goto menu

::-----------------HTTP���ع��ܺ�������-----------------
:http
echo Source:      "%~1"
echo Destination: "%~f2"
echo Start downloading. . .
cscript -nologo -e:jscript "%~f0" "%~1" "%~2"
echo OK!
goto :eof

*/
var iLocal,iRemote,xPost,sGet;
iLocal =WScript.Arguments(1); 
iRemote = WScript.Arguments(0); 
iLocal=iLocal.toLowerCase();
iRemote=iRemote.toLowerCase();
xPost = new ActiveXObject("Microsoft"+String.fromCharCode(0x2e)+"XMLHTTP");
xPost.Open("GET",iRemote,0);
xPost.Send();
sGet = new ActiveXObject("ADODB"+String.fromCharCode(0x2e)+"Stream");
sGet.Mode = 3;
sGet.Type = 1; 
sGet.Open(); 
sGet.Write(xPost.responseBody);
sGet.SaveToFile(iLocal,2); 
