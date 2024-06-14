package scanner

func ScanServices(ipRange string, ports []int) map[string][]string {
    services := make(map[string][]string)

    for _, port := range ports {
        for _, ip := range getIPRange(ipRange) {
            if scanPort(ip, port) {
                services[ip] = append(services[ip], getServiceName(port))
            }
        }
    }

    return services
}

func getServiceName(port int) string {
    serviceMap := map[int]string{
        21:   "FTP",
        22:   "SSH",
        23:   "Telnet",
        25:   "SMTP",
        80:   "HTTP",
        110:  "POP3",
        443:  "HTTPS",
        8080: "HTTP-Alt",
    }

    return serviceMap[port]
}
