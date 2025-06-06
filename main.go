package main

import (
	"flag"
	"fmt"
	"go-study/mymodule"
)

func printHello() {
	fmt.Println("Hello")
}

func main() {
	tcaAddressFlag := flag.String("tca-address", "0x70", "I2C address of the TCA9548A multiplexer (default: 0x70)")
	withoutMultiplexerFlag := flag.Bool("without-multiplexer", false, "Set to true if INA260 is connected directly without TCA9548A multiplexer (default: false)")

	flag.Parse()
	if *withoutMultiplexerFlag {
		fmt.Println("Running without TCA9548A multiplexer, using INA260 directly.")
		tcaAddressFlag = nil // Set to nil to skip TCA address usage
	} else {
		fmt.Println("Running with TCA9548A multiplexer.")
	}

	tcaAddressStr := *tcaAddressFlag
	fmt.Printf("Using TCA address: %s\n", tcaAddressStr)
	mymodule.Run()
}
