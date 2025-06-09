package main

import (
	"flag"
	"fmt"
	"github.com/stretchr/testify/assert"
	"go-study/mymodule"
	"periph.io/x/conn/v3/i2c"
	"periph.io/x/conn/v3/i2c/i2creg"
	"periph.io/x/host/v3"
	"testing"
)

func printHello() {
	fmt.Println("Hello")
}

func TestPrintHello(t *testing.T) {
	assert := assert.New(t)
	assert.NotNil(printHello, "printHello function should not be nil")
}

func initializeI2C() (i2c.BusCloser, error) {
	if _, err := host.Init(); err != nil {
		return nil, fmt.Errorf("failed to initialize host: %w", err)
	}
	bus, err := i2creg.Open("") // Opens the default I2C bus
	if err != nil {
		return nil, fmt.Errorf("failed to open I2C bus: %w", err)
	}
	return bus, nil
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
	bus, _ := initializeI2C()
	dev := i2c.Dev{Bus: bus, Addr: 0x40} // Example device initialization, replace with actual device address
	fmt.Println(dev)
}
