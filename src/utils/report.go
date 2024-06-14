package utils

import (
    "fmt"
)

func GenerateReport(data map[string]interface{}) {
    for key, value := range data {
        fmt.Printf("%s: %v\n", key, value)
    }
}
