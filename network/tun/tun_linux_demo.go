package main

import (
	"fmt"
	"github.com/songgao/water"
	"github.com/songgao/water/waterutil"
	"log"
	"network/tun/utils"
)

func main() {
	// 定义虚拟网卡的IP和网络掩码
	// 通过 CIDR 方式进行定义, ip=172.10.0.101, netmask=255.255.255.0
	cidr := "172.10.0.101/24"

	// 网卡配置，指定设备类型为 TUN。MacOS系统是无法指定 tun 设备名称的，所以这里都不指定了
	tunConfig := water.Config{
		DeviceType: water.TUN,
	}
	// 创建 TUN 设备
	dev, err := water.New(tunConfig)
	if err != nil {
		log.Fatalln("create tun error", err)
	}
	defer func() {
		_ = dev.Close()
	}()

	// 启动网卡、设置 IP、添加网段路由表
	// 设置网卡 MTU
	utils.ExecCmd("/sbin/ip", "link", "set", "dev", dev.Name(), "mtu", "1500")
	// 设置网卡 IP 和 网络掩码 （CIDR方式）
	utils.ExecCmd("/sbin/ip", "addr", "add", cidr, "dev", dev.Name())
	// 启用网卡
	utils.ExecCmd("/sbin/ip", "link", "set", "dev", dev.Name(), "up")

	// 缓存区大小与 MTU 一致
	buf := make([]byte, 1500)
	for {
		n, err := dev.Read(buf)
		if err != nil {
			panic(err)
		}
		// IP 数据包
		data := buf[:n]
		if !waterutil.IsIPv4(data) {
			// 只处理 ipv4 数据包
			continue
		}
		// 来源IP、端口
		srcIp := waterutil.IPv4Source(data)
		srcPort := waterutil.IPv4SourcePort(data)
		// 目标IP
		destIp := waterutil.IPv4Destination(data)
		destPort := waterutil.IPv4DestinationPort(data)
		fmt.Printf("source: %s:%d, dest: %s:%d\n", srcIp.String(), srcPort, destIp.String(), destPort)

		// 如果目标IP与本地 Tun 设备的IP相同，则将数据包写回到 Tun 设备
		if srcIp.Equal(destIp) {
			_, _ = dev.Write(data)
		} else {
			// TODO 将数据通过公网发送给服务端进行转发，并将公网响应数据包写入 Tun 设备
			fmt.Println("公网转发")
		}
	}
}
