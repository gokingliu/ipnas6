package util

import (
	"os/exec"
	"strings"
)

// GetIPv4 获取 IPv4 地址
func GetIPv4() string {
	// 读取 IPv4 命令
	IPv4LinuxCommand := "ifconfig -a | grep 'inet [^f:]' | grep -E -v '10|127|172|192|inet6' | awk '{print $2}'| tr -d 'addr:'"
	IPv4OpenWRTCommand := `ip -o addr show | grep -E 'eth|en' | grep 'inet [^f:]' | sed -nr 's#.+? +inet ([0-9.]+)/[0-9]+ brd [0-9./]+ scope global .*#\1#p'`
	// 执行 Linux 命令
	result, err := exec.Command("/bin/sh", "-c", IPv4LinuxCommand).Output()
	// 执行错误返回空字符串
	if err != nil {
		// 执行 OpenWRT 命令
		result, _ = exec.Command("/bin/sh", "-c", IPv4OpenWRTCommand).Output()
		return ""
	}
	// 读取命令结果中的文本z
	IPv4 := strings.TrimSpace(string(result))

	return IPv4
}

// GetIPv6 获取 IPv6 地址
func GetIPv6() string {
	// 读取 IPv6 命令
	IPv6LinuxCommand := "ifconfig -a | grep 'inet6 [^f:]' | grep 'temporary' | awk '{print $2}'| tr -d 'addr:'"
	IPv6OpenWRTCommand := `ip -o addr show | grep -E 'eth|en' | grep -v deprecated | grep 'inet6 [^f:]' | sed -nr 's#^.+? +inet6 ([a-f0-9:]+)/.+? scope global .*? valid_lft ([0-9]+sec) .*#\2 \1#p' | grep 'ff:fe'| sort -nr | head -n1 | cut -d' ' -f2`
	// 执行命令
	result, err := exec.Command("/bin/sh", "-c", IPv6LinuxCommand).Output()
	// 执行错误返回空字符串
	if err != nil {
		// 执行 OpenWRT 命令
		result, _ = exec.Command("/bin/sh", "-c", IPv6OpenWRTCommand).Output()
		return ""
	}
	// 读取命令结果中的文本
	IPv6 := strings.TrimSpace(string(result))

	return IPv6
}
