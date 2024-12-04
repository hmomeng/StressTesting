package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var ListenPath = "." // 将filePath定义为全局变量
var IPList []net.IP

func main() {
	//exeFilePath, _ := os.Getwd()
	exeName := filepath.Base(os.Args[0])
	ListenPath = filepath.Dir(os.Args[0])
	LinstentPort := "9999"

	//fmt.Println(exeFilePath, exeName, ListenPath, LinstentPort)
	//exeFilePath := os.Args[0]
	//cmdArgs := os.Args[1:]

	// 判断路径是否正确
	if len(os.Args) == 2 {
		//filePath = cmdArgs[0]
		arg := os.Args[1]
		if arg == "?" || arg == "help" || arg == "/help" {
			// exePath := args[0]
			//i := strings.LastIndex(filePath, "\\")

			fmt.Println(exeName + " [端口] [路径]")

			fmt.Printf("默认端口：%s, 默认路径：%s \n", LinstentPort, ".")
			fmt.Println("端口范围: 0-65535")
			return
		}
		// // 参数提示
		// if filePath == "?" || filePath == "help" || filePath == "/help" {
		// 	// exePath := args[0]
		// 	i := strings.LastIndex(filePath, "\\")
		// 	exeName := filePath[i+1:]
		// 	fmt.Println(exeName + " [路径] [端口]")

		// 	fmt.Printf("默认端口：%s, 默认路径：%s \n", LinstentPort, ".")
		// 	fmt.Println("端口范围:1024-65535")
		// 	return
		// }

	}

	if len(os.Args) == 3 {
		ListenPath = os.Args[2]
		_, err := os.Stat(ListenPath)
		if err == nil {
			fmt.Printf("正在加载路径：%s\n", ListenPath)
		} else if os.IsNotExist(err) {
			// 路径不存在
			fmt.Printf("路径 '%s' 不存在。\n", ListenPath)
			return
		} else {
			// 发生其他错误
			fmt.Printf("发生错误：%v\n", err)
			return
		}

		LinstentPort = os.Args[1]
		if !isValidPort(LinstentPort) {
			fmt.Println("端口输入错误")
			return
		}
	}
	//LinstentPort = ":" + LinstentPort

	// 注册文件下载处理函数
	http.HandleFunc("/download", downloadHandler)

	// 文件服务器处理
	http.Handle("/", http.FileServer(http.Dir(ListenPath)))
	//dir, _ := os.Getwd()
	fmt.Print("当前系统所有IP: ")
	//PrintIPList()

	GetIPList()
	//fmt.Println(IPList)
	fmt.Print("\n\n")
	fmt.Println("FilePath : " + ListenPath)
	fmt.Println("Listening: " + string(IPList[0].String()) + ":" + LinstentPort)

	//fmt.Println(LinstentPort)
	http.ListenAndServe("0.0.0.0:"+LinstentPort, nil)
}

func downloadHandler(w http.ResponseWriter, r *http.Request) {
	// 获取请求路径
	path := r.URL.Path

	// 去掉前缀，得到相对路径
	relativePath := strings.TrimPrefix(path, "/download/")

	// 构造完整的文件路径
	fileFullPath := ListenPath + string(os.PathSeparator) + relativePath

	// 打开文件
	file, err := os.Open(fileFullPath)
	if err != nil {
		http.Error(w, "文件不存在", http.StatusNotFound)
		return
	}
	defer file.Close()

	// 获取文件修改时间
	fileInfo, err := file.Stat()
	if err != nil {
		http.Error(w, "无法获取文件信息", http.StatusInternalServerError)
		return
	}
	// // 设置响应头，告诉浏览器这是一个文件下载
	// w.Header().Set("Content-Disposition", "attachment; filename="+file.Name())

	// // 将文件内容写入响应体
	// http.ServeContent(w, r, file.Name(), fileInfo.ModTime(), file)

	// 设置响应头，告诉浏览器这是一个文件下载
	w.Header().Set("Content-Disposition", "attachment; filename="+filepath.Base(file.Name()))

	// 将文件内容写入响应体
	http.ServeContent(w, r, filepath.Base(file.Name()), fileInfo.ModTime(), file)
}

func isValidPort(portStr string) bool {
	// 尝试将输入的字符串转换为整数
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return false // 无法转换为整数，不是有效的端口
	}

	// 确保端口在正常的范围内（1-65535）
	if port < 1 || port > 65535 {
		return false
	}

	// 检查是否是系统占用的端口（一般情况下，系统占用的端口在1-1023范围内）
	if port >= 1 && port <= 1023 {
		return false
	}

	return true

}

// func getLocalIP() string {
// 	// 获取所有网络接口的信息
// 	interfaces, err := net.Interfaces()
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return ""
// 	}

// 	// 遍历每个网络接口
// 	for _, iface := range interfaces {
// 		// 获取每个接口的地址信息
// 		addrs, err := iface.Addrs()
// 		if err != nil {
// 			fmt.Println("Error:", err)
// 			continue
// 		}

// 		// 遍历每个地址
// 		for _, addr := range addrs {
// 			// 将地址转换为 IP 地址类型
// 			ip, _, err := net.ParseCIDR(addr.String())
// 			if err != nil {
// 				fmt.Println("Error:", err)
// 				continue
// 			}

// 			// 过滤掉非IPv4地址和回环地址
// 			if ip.To4() != nil && !ip.IsLoopback() {
// 				return ip.String()
// 			}
// 		}
// 	}
// 	return ""
// }

func PrintIPList() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, addr := range addrs {
		var ip net.IP
		switch v := addr.(type) {
		case *net.IPNet:
			ip = v.IP
		case *net.IPAddr:
			ip = v.IP
		}

		if ip == nil || ip.IsLoopback() {
			continue
		}

		// 过滤掉IPv6地址
		ip = ip.To4()
		if ip == nil {
			continue
		}

		// 过滤掉169开头的IP地址
		if strings.HasPrefix(ip.String(), "169.") {
			continue
		}

		fmt.Print(ip, " ")
	}
}

func GetIPList() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, addr := range addrs {
		var ip net.IP
		switch v := addr.(type) {
		case *net.IPNet:
			ip = v.IP
		case *net.IPAddr:
			ip = v.IP
		}

		if ip == nil || ip.IsLoopback() {
			continue
		}

		// 过滤掉IPv6地址
		ip = ip.To4()
		if ip == nil {
			continue
		}

		// 过滤掉169开头的IP地址
		if strings.HasPrefix(ip.String(), "169.") {
			continue
		}
		IPList = append(IPList, ip)
		fmt.Print(ip, " ")
	}
}
