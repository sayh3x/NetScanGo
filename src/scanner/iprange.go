package scanner

import (
    "NetScanGo/src/utils"
)

func getIPRange(cidr string) []string {
    return utils.ParseCIDR(cidr)
}
