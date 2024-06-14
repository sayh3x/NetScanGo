package scanner

import (
    "fmt"
    "net"
    "time"
)

func ScanPorts(ipRange string, ports []int) map[string][]int {
    openPorts := make(map[string][]int)

    for _, port := range ports {
        for _, ip := range getIPRange(ipRange) {
            if scanPort(ip, port) {
                openPorts[ip] = append(openPorts[ip], port)
            }
        }
    }

    return openPorts
}

func scanPort(ip string, port int) bool {
    address := fmt.Sprintf("%s:%d", ip, port)
    conn, err := net.DialTimeout("tcp", address, 1*time.Second)
    if err != nil {
        return false
    }
    conn.Close()
    return true
}
