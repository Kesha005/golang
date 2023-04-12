package main


import("fmt")

type Plane struct {
	distance int
}

type Car struct {
	distance int
}


func (car *Car) Drive(dist int){
	car.distance += dist
}

func (plane *Plane) Fly(dist int){
	plane.distance += dist
}


func (p *Plane) getDistance()int{
	return p.distance
}

func (c *Car) getDistance() int {
	return c.distance
  }

func getObjectDistance(i interface{})int{
	switch i.(type) {
	case *Car:
	  c := i.(*Car) // Type assertion
	  return c.getDistance()
	case *Plane:
	  p := i.(*Plane) // Type assertion
	  return p.getDistance()
	}
	return 0

}

func main(){
	tesla := &Car{}
	airbus := &Plane{}
	tesla.Drive(10)
	airbus.Fly(1000)
	fmt.Println("Tesla drived", tesla.getDistance())
	fmt.Println("Airbus flued", airbus.getDistance())
}