# TEFAS Scraper

Simple go package to retrieve fund information from [TEFAS](https://www.tefas.gov.tr/).

## Installation

To install this package, you can use the `go get` command:

```shell
go get github.com/artuncolak/tefas
```

## Example

Here's a simple example of how to use the `tefas` package to retrieve fund information:

```go
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/artuncolak/tefas"
)

func main() {
    // Create a new instance of TefasClient
    client := tefas.New()

    // Define the fund information request
    fundInfoRequest := tefas.FundInfoRequest{
        Type:      tefas.YAT, // Fund type
        Code:      "HMS", // Fund code
        StartDate: time.Date(2024, 8, 24, 0, 0, 0, 0, time.UTC), // Start date
        EndDate:   time.Now(), // End date
    }

    // Retrieve fund information
    funds, err := client.GetFundInfo(fundInfoRequest)
    if err != nil {
        log.Fatalf("Error retrieving fund information: %v", err)
    }

    // Print the retrieved fund information
    for _, fund := range funds {
        fmt.Printf("Date: %s, Code: %s, Name: %s, Price: %.2f\n", fund.Date, fund.Code, fund.Name, fund.Price)
    }
}
```