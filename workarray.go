package main

import(
	"fmt"
)

func main(){
	x := []int{
		48,96,86,68,
		57,82,63,70,
		37,34,83,27,
		19,97, 9,17,
		21,4,4,45,4343,
		423,432,45,65,667,
		557,84,365,2,3,3,
		676,8787,34,5698,
		884532,6680,45675,
		321,
	}

	
	min := x[0]
	for _ , element := range x {
		if(element < min){
			min = element
		}
	}
	fmt.Println("Minimum value is", min)

}

