package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/jaypipes/ghw"
)

func main() {
	getCPUInfo()
	getMemoryInfo()
	getDiskInfo()

	for i, a := range os.Args {
		if i == 0 {
			continue
		}
		if a == "nowait" {
			return
		}
	}
	fmt.Printf("\n\n\n")
	fmt.Println("按回车结束程序")
	fmt.Scanln()

}

func getCPUInfo() {
	// 获取 CPU 型号
	cpuInfo, err := ghw.CPU()
	if err != nil {
		fmt.Printf("Error getting CPU info: %v\n", err)
	} else {
		fmt.Printf("CPU: %s\n", cpuInfo.String())
		for i := range cpuInfo.Processors {
			fmt.Printf("CPU%d: %s\n", i+1, cpuInfo.Processors[i].Model)
		}
		// fmt.Printf("CPU: %s\n", cpuInfo.Processors[0].Model)
	}
}

func getMemoryInfo() {
	// 获取内存信息
	fmt.Println("")
	memInfo, err := ghw.Memory()
	if err != nil {
		fmt.Printf("Error getting memory info: %v\n", err)
		return
	}

	// 打印内存模块的信息
	if len(memInfo.Modules) == 0 {
		fmt.Println("No memory module information available.")
		return
	}

	var allSize float64
	for i, module := range memInfo.Modules {
		fmt.Printf("内存%d:", i+1)
		fmt.Printf("  制造商: %s\t", module.Vendor)
		sizeGB := float64(module.SizeBytes) / (1024 * 1024 * 1024)
		fmt.Printf("  Size: %0.2f GB\n", sizeGB)
		allSize += sizeGB
	}

	fmt.Printf("内存总大小：%0.2f GB\n", allSize)
}

func getDiskInfo() {
	// 获取硬盘信息
	fmt.Println("")
	cmd := exec.Command("wmic", "diskdrive", "get", "model,size")
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	// 解析 wmic 输出
	lines := strings.Split(string(output), "\n")
	var diskInfos []string

	// 将标题行移除并保留硬盘信息行
	for _, line := range lines[1:] {
		if strings.TrimSpace(line) != "" {
			diskInfos = append(diskInfos, strings.TrimSpace(line))
		}
	}

	// 如果两个路径下都没有smartctl则退出
	smartctlPath := "./smartctl.exe"
	_, err = os.Stat(smartctlPath)
	if os.IsNotExist(err) {
		smartctlPath = "./kktool/smartctl.exe"
		_, err = os.Stat(smartctlPath)
		if os.IsNotExist(err) {
			fmt.Println("\n\n无法获取硬盘通电时间,请确保smartctl.exe在目录下!")
			fmt.Scanln()
			return
		}
	}

	fmt.Printf("%-36s %-13s %s\n", "硬盘型号", "容量", "通电时间")
	nvmeNumber := 0
	var powerOnHours string
	// 循环查询每个硬盘的 Power_On_Hours
	for i := 0; i < len(diskInfos); i++ {
		// 处理AHCI、SCSI协议硬盘
		device := fmt.Sprintf("/dev/sd%c", 'a'+i)
		cmd := exec.Command(smartctlPath, "-a", device)
		output, err := cmd.Output()
		if err != nil {
			//处理Nvme协议硬盘
			device := fmt.Sprintf("/dev/nvme%d", nvmeNumber)
			cmd := exec.Command(smartctlPath, "-a", device)
			output, err := cmd.Output()
			if err != nil {
				powerOnHours = "未知"
				fmt.Println("无法识别硬盘smart信息，请通过CrystalDiskInfo获取")
				//continue
			}
			nvmeNumber++
			scanner := bufio.NewScanner(bytes.NewReader(output))
			for scanner.Scan() {
				line := scanner.Text()
				if strings.Contains(line, "Power On Hours:") {
					parts := strings.Split(line, ":")
					if len(parts) > 1 {
						powerOnHours = strings.TrimSpace(parts[1])
						powerOnHours = strings.ReplaceAll(powerOnHours, ",", "")
					}
					break
				}
			}
		} else {
			scanner := bufio.NewScanner(bytes.NewReader(output))
			for scanner.Scan() {
				line := scanner.Text()
				if strings.Contains(line, "Power_On_Hours") {
					// 解析 RAW_VALUE
					parts := strings.Fields(line)
					for _, part := range parts {
						if part == "Power_On_Hours" {
							// RAW_VALUE 通常是最后一列
							powerOnHours = parts[9]
							break
						}
					}
					break
				}
			}
		}

		// 打印硬盘信息和 Power_On_Hours
		if i < len(diskInfos) {
			a := strings.Fields(diskInfos[i])
			diskname := strings.Join(a[:len(a)-1], " ")
			disksize := a[len(a)-1]
			fmt.Printf("%-40s %-15s %s\n", diskname, disksize[0:len(disksize)-9]+"GB", powerOnHours+"h")
			// 转为实际容量大小
			//disksizeForFloat, _ := strconv.ParseFloat(disksize, 64)
			//fmt.Printf("%-35s%sGB/%0.2fGB%15s\n", diskname, disksize[0:len(disksize)-9], (disksizeForFloat / 1024 / 1024 / 1024), powerOnHours+"h")
			powerOnHours = ""
		}
	}
}
