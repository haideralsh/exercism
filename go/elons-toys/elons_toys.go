package elon

import "fmt"

func (c *Car) Drive() {
	if c.battery < c.batteryDrain { return }

	c.battery = c.battery - c.batteryDrain
	c.distance = c.distance + c.speed
}

func (c *Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", c.distance)

}

func (c *Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", c.battery)
}

func (c *Car) CanFinish(trackDistance int) bool {
	mileage := trackDistance / c.speed

	if mileage * c.batteryDrain > c.battery {
		return false
	}

	return true

}
