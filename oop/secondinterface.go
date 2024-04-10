package main

import (
	"fmt"
	"math"
)


type Sircle struct{
	Radius int
}

type Rect struct{
	Height int
	Width int
}

type CalcArea interface{
	Area()float64
}

func (s *Sircle)Area()float64{
	return float64(s.Radius) * math.Pi
}

func (r *Rect)Area()float64{
	return float64(r.Width * r.Height) 
}


func GetArea(c CalcArea)float64{
	
	return c.Area()
}


func main(){
	rec := Rect{23,32}
	c:= Sircle{32}
	GetArea(c)
}