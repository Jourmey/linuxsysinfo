package linuxsysinfo

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

const (
	//ProcMeminfoFile     = "/proc/meminfo"
	//ProcCpuinfoFile     = "/proc/cpuinfo"

	ProcMeminfoFile = `C:\gopath\work\linuxsysinfo\example\meminfo`
	ProcCpuinfoFile = `C:\gopath\work\linuxsysinfo\example\cpuinfo`
)

//var (
//	reTwoColumns = regexp.MustCompile("\t+: ")
//	reExtraSpace = regexp.MustCompile(" +")
//	reCacheSize  = regexp.MustCompile(`^(\d+) KB$`)
//)

type info map[string]string

func getInfo0(file string, fu func(text string)) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		fu(s.Text())
	}
	return nil
}

func getInfo1(file string) (info, error) {
	info := make(info)
	m := regexp.MustCompile("\t+: ")

	err := getInfo0(file, func(text string) {
		sl := m.Split(text, 2)
		if sl != nil && len(sl) == 2 {
			info[sl[0]] = sl[1]
		}
	})

	return info, err
}

func getInfo2(file string) (info, error) {
	info := make(info)
	m := regexp.MustCompile(": +")

	err := getInfo0(file, func(text string) {
		sl := m.Split(text, 2)
		if sl != nil && len(sl) == 2 {
			info[sl[0]] = sl[1]
		}
	})

	return info, err
}

func getInfo3(file string) (info, error) {
	//command := exec.Command("/bin/bash", "-c", "ifconfig")
	//txt, err := command.CombinedOutput()
	//if err != nil {
	//	return nil, err
	//}

	txt2 := `enP1p38s0f0: flags=4099<UP,BROADCAST,MULTICAST>  mtu 1500
        inet 192.168.2.254  netmask 255.255.255.0  broadcast 192.168.2.255
        ether 6c:b3:11:41:65:90  txqueuelen 1000  (Ethernet)
        RX packets 0  bytes 0 (0.0 B)
        RX errors 0  dropped 0  overruns 0  frame 0
        TX packets 0  bytes 0 (0.0 B)
        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0
        device memory 0xf0e0300000-f0e031ffff

enP1p38s0f1: flags=4099<UP,BROADCAST,MULTICAST>  mtu 1500
        inet 192.168.3.254  netmask 255.255.255.0  broadcast 192.168.3.255
        ether 6c:b3:11:41:65:91  txqueuelen 1000  (Ethernet)
        RX packets 0  bytes 0 (0.0 B)
        RX errors 0  dropped 0  overruns 0  frame 0
        TX packets 0  bytes 0 (0.0 B)
        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0
        device memory 0xf0e0320000-f0e033ffff

enP1p39s0f0: flags=4099<UP,BROADCAST,MULTICAST>  mtu 1500
        inet 192.168.4.254  netmask 255.255.255.0  broadcast 192.168.4.255
        ether 6c:b3:11:41:65:94  txqueuelen 1000  (Ethernet)
        RX packets 0  bytes 0 (0.0 B)
        RX errors 0  dropped 0  overruns 0  frame 0
        TX packets 0  bytes 0 (0.0 B)
        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0
        device memory 0xf0e0400000-f0e04fffff

enP1p39s0f1: flags=4163<UP,BROADCAST,RUNNING,MULTICAST>  mtu 1500
        inet 192.168.16.134  netmask 255.255.255.0  broadcast 192.168.16.255
        inet6 fe80::6eb3:11ff:fe41:6595  prefixlen 64  scopeid 0x20<link>
        ether 6c:b3:11:41:65:95  txqueuelen 1000  (Ethernet)
        RX packets 23538432  bytes 3831224774 (3.5 GiB)
        RX errors 0  dropped 4  overruns 0  frame 0
        TX packets 4906291  bytes 4291056996 (3.9 GiB)
        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0
        device memory 0xf0e0500000-f0e05fffff

enP1p43s0: flags=4099<UP,BROADCAST,MULTICAST>  mtu 1500
        inet 192.168.1.254  netmask 255.255.255.0  broadcast 192.168.1.255
        inet6 fe80::2e0:3ff:fe50:c  prefixlen 64  scopeid 0x20<link>
        ether 00:e0:03:50:00:0c  txqueuelen 1000  (Ethernet)
        RX packets 1321  bytes 88845 (86.7 KiB)
        RX errors 0  dropped 0  overruns 0  frame 0
        TX packets 132  bytes 16734 (16.3 KiB)
        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0
        device memory 0xf0e0b00000-f0e0b1ffff

lo: flags=73<UP,LOOPBACK,RUNNING>  mtu 65536
        inet 127.0.0.1  netmask 255.0.0.0
        inet6 ::1  prefixlen 128  scopeid 0x10<host>
        loop  txqueuelen 1  (Local Loopback)
        RX packets 282655573  bytes 68290186896 (63.6 GiB)
        RX errors 0  dropped 0  overruns 0  frame 0
        TX packets 282655573  bytes 68290186896 (63.6 GiB)
        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0
`

	tt := strings.Split(txt2, "\n\n")

	i := make([]IfconfigInfo, len(tt))
	for _, s := range tt {
		ifc, err := ifconfigParse(s)
		if err != nil {
			continue
		}
		i = append(i, *ifc)
	}

	fmt.Println(tt)
	return nil, nil
}
