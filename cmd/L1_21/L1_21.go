package L1_21

import "fmt"

type transport interface {
	navigateToDestination()
}

type client struct {
}

func (t *client) startNavigation(trans transport) {
	fmt.Println("Client start navigate")
	trans.navigateToDestination()
}

type boat struct {
}

func (b *boat) navigateToDestination() {
	fmt.Println("boat navigate to island")
}

type car struct {
}

func (c *car) driveToDestination() {
	fmt.Println("car going to destination")
}

type carAdapter struct {
	carTransport *car
}

func (c *carAdapter) navigateToDestination() {
	fmt.Println("adapting car to navigate to destination")
	c.carTransport.driveToDestination()
}

func TwentyOne() {
	client := &client{}
	boat := &boat{}
	client.startNavigation(boat)
	car := &car{}
	carAdapter := &carAdapter{carTransport: car}
	client.startNavigation(carAdapter)
}
