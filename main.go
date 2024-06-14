package main

import (
    "NetScanGo/src/scanner"
    "NetScanGo/src/utils"
    "flag"
    "fmt"
    "strings"
)

func main() {

    ipRange := flag.String("iprange", "", "Range of IP addresses to scan (CIDR format)")
    ports := flag.String("ports", "", "Comma-separated list of ports to scan")
    doScan := flag.Bool("d", false, "Perform detailed scan (scan ports and services)")
    flag.Parse()


    printLogo()


    if *ipRange == "" && *ports == "" && !*doScan {
        printUsageExamples()
        return
    }


    if *ipRange != "" && !*doScan {
        ipList := utils.ParseCIDR(*ipRange)
        fmt.Println("IP Range:")
        for _, ip := range ipList {
            fmt.Println(ip)
        }
        return
    }


    if *doScan {
        if *ipRange == "" {
            fmt.Println("Error: IP range is required for detailed scan")
            return
        }

        
        portList := []int{21, 22, 23, 25, 80, 110, 443, 8080}
        if *ports != "" {
            portList = parsePorts(*ports)
        }

        fmt.Println("Starting detailed scan...")

        
        openPorts := scanner.ScanPorts(*ipRange, portList)
        fmt.Printf("Open Ports: %v\n", openPorts)


        activeDevices := scanner.ScanDevices(*ipRange)
        fmt.Printf("Active Devices: %v\n", activeDevices)

        
        activeServices := scanner.ScanServices(*ipRange, portList)
        fmt.Printf("Active Services: %v\n", activeServices)


        vulnerabilities := scanner.ScanVulnerabilities(*ipRange)
        fmt.Printf("Vulnerabilities: %v\n", vulnerabilities)

        
        trafficAnalysis := scanner.AnalyzeTraffic(*ipRange)
        fmt.Printf("Traffic Analysis: %v\n", trafficAnalysis)

        
        reportData := map[string]interface{}{
            "Open Ports":         openPorts,
            "Active Devices":     activeDevices,
            "Active Services":    activeServices,
            "Vulnerabilities":    vulnerabilities,
            "Traffic Analysis":   trafficAnalysis,
        }
        utils.GenerateReport(reportData)

        fmt.Println("NetScanGo completed.")
    }
}


func parsePorts(ports string) []int {
    var portList []int
    for _, port := range strings.Split(ports, ",") {
        var p int
        fmt.Sscanf(port, "%d", &p)
        portList = append(portList, p)
    }
    return portList
}


func printLogo() {
    logo := `
 _   _      _   _____                                      
| \ | |    | | / ____|                                     
|  \| | ___| |_| |  __  ___ _ __ ___  ___ _ __   __ _ _ __  
| .  \ |/ _ \ __| | |_ |/ _ \ '__/ __|/ _ \ '_ \ / _  | '_ \ 
| |\  |  __/ |_| |__| |  __/ |  \__ \  __/ | | | (_| | | | |
|_| \_|\___|\__|\_____|\___|_|  |___/\___|_| |_|\__,_|_| |_|
`
    fmt.Println(logo)
}


func printUsageExamples() {
    fmt.Println("Usage Examples:")
    fmt.Println("  go run main.go -iprange=192.168.1.0/24")
    fmt.Println("  go run main.go -iprange=192.168.1.0/24 -d")
    fmt.Println("  go run main.go -iprange=192.168.1.0/24 -d -ports=21,22,80,443")
}
