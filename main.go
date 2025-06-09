package main

import (
	"flag"
	"fmt"
	"github.com/stretchr/testify/assert"
	"go-study/mymodule"
	"testing"
)

func printHello() {
	fmt.Println("Hello")
}

func TestPrintHello(t *testing.T) {
	assert := assert.New(t)
	assert.NotNil(printHello, "printHello function should not be nil")
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
