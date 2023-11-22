package main


import(
	"fmt"
	"sort"
)


func main(){

	ages := map[string]int{"Kerim":18,"Kakamyrat":16,"Selbi":9}


	var names []string

	for name := range ages{
		names = append(names, name)
	}
	sort.Strings(names)

	for _ , name := range names{
		fmt.Println(name, ages[name])
	}


}