package scanner

import (
    "fmt"
    "net"
    "time"
)

func ScanDevices(ipRange string) []string {
    var devices []string

    for _, ip := range getIPRange(ipRange) {
        if isDeviceActive(ip) {
            devices = append(devices, ip)
        }
    }

    return devices
}

func isDeviceActive(ip string) bool {
    _, err := net.DialTimeout("tcp", fmt.Sprintf("%s:80", ip), 1*time.Second)
    if err != nil {
        return false
    }
    return true
}
