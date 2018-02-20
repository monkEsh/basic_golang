// in go there are no classes
package main

import (
	"fmt"
)

const uSixTeenBitMax float64 = 65535
const kmhMultiple float64 = 1.60934

type car struct {
	gasPedal      uint16 //min 0 max 65535
	breakPedal    uint16
	steeringWheel int16 // -32k - +32k
	topSpeedKmh   float64
}

func (c car) kmh() float64 {
	return float64(c.gasPedal) * (c.topSpeedKmh / uSixTeenBitMax)
}

func (c car) mph() float64 {
	return float64(c.gasPedal) * (c.topSpeedKmh / uSixTeenBitMax / kmhMultiple)
}

// below function will update value at pointer
func (c *car) newTopSpeed(newSpeed float64) {
	c.topSpeedKmh = newSpeed
}

// below function will return struct with updated topspeed value
func carNewTopSpeed(c car, speed float64) car {
	c.topSpeedKmh = speed
	return c
}

func main() {
	aCar := car{gasPedal: 65000, breakPedal: 0, steeringWheel: 12561, topSpeedKmh: 225.0}
	bCar := car{12312, 2222, 24312, 222.0}

	// To print all element in struct we can use following
	fmt.Println(aCar)
	fmt.Println(bCar)

	// to print specific value in struct we can use following
	fmt.Println(aCar.steeringWheel)

	// use function to print speed in kmh
	fmt.Println(aCar.kmh())

	// use function to print speed in mph
	fmt.Println(aCar.mph())

	// Change top speed
	aCar.newTopSpeed(500)

	// use function to print speed in kmh
	fmt.Println(aCar.kmh())

	// use function to print speed in mph
	fmt.Println(aCar.mph())

	// change speed without updating aCar
	aCar = carNewTopSpeed(aCar, 100)

	// use function to print speed in kmh
	fmt.Println(aCar.kmh())

	// use function to print speed in mph
	fmt.Println(aCar.mph())
}
