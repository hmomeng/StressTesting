package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
)

/*
downloadListPath 指定服务器IP地址，download.txt内保存需要下载的地址，每行一个
downloadDir 	 保存的文件夹名
startfilename	 下载完完毕后启动的脚本
*/

var (
	downloadListPath = "http://127.0.0.1:9999/downloadlist.txt"
	downloadDir      = "./kktool"
	startfilename    = "StartTest.bat"
)

func main() {
	// 获取下载列表
	fileUrls, err := fetchDownloadList(downloadListPath)
	if err != nil {
		fmt.Printf("无法获取下载列表: %v\n", err)
		return
	}

	// 确保下载目录存在
	if err := os.MkdirAll(downloadDir, os.ModePerm); err != nil {
		fmt.Printf("创建下载目录失败: %v\n", err)
		return
	}

	// 使用 WaitGroup 管理多协程
	var wg sync.WaitGroup
	var mu sync.Mutex
	failedUrls := []string{}

	for _, url := range fileUrls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			fmt.Println("开始下载：", url)
			if err := downloadFile(url, downloadDir); err != nil {
				mu.Lock()
				failedUrls = append(failedUrls, url)
				mu.Unlock()
				fmt.Printf("下载失败: %s, 错误: %v\n", url, err)
			}
		}(url)
	}

	wg.Wait()

	// 打印无法访问的网址
	if len(failedUrls) > 0 {
		fmt.Println("以下网址无法访问:")
		for _, url := range failedUrls {
			fmt.Println(url)
		}
	}

	// 执行下载路径下的 start.exe
	startExePath := filepath.Join(downloadDir, startfilename)
	if err := execCommand(startExePath); err != nil {
		fmt.Printf("执行 start.exe 失败: %v\n", err)
	}
}

// fetchDownloadList 从指定 URL 获取下载列表
func fetchDownloadList(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP 请求失败: %s", resp.Status)
	}

	var urls []string
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			urls = append(urls, line)
		}
	}

	return urls, scanner.Err()
}

// downloadFile 下载单个文件到指定目录
func downloadFile(url, dir string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP 请求失败: %s", resp.Status)
	}

	fileName := filepath.Base(url)
	filePath := filepath.Join(dir, fileName)

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	return err
}

// execCommand 执行指定路径的可执行文件
func execCommand(filePath string) error {
	cmd := exec.Command(filePath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
