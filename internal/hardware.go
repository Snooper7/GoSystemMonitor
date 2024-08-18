package hardware

import (
	"fmt"
	"runtime"

	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

func GetSystemSection() (string, error) {
	runTimeOS := runtime.GOOS

	vmStat, err := mem.VirtualMemory()
	if err != nil {
		return "", err
	}

	hostStat, err := host.Info()
	if err != nil {
		return nil, err
	}

	output :=
		fmt.Sprintf("Hostname: %s\nTotal Memory: %d\nUsedMemory: %d\nOS: %s",
			hostStat.Hostname, vmStat.Total, vmStat.Used, runTimeOS)

	return output, nil
}

func GetCpuSection() {

}

func GetDiskSection() {

}
