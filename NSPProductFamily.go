package main

import (
	 "fmt"
"slices")

type Family map[string]*[]string

func makeSlice() *[]string {
    s := make([]string, 0)
    return &s
}

func removeElement(v *[]string, index int) {
		if index != -1 {
			(*v)[index] = (*v)[len(*v) - 1]
			*v = (*v)[: len(*v) - 1 ]
		}
}

func print(family Family) {
	for k, v := range family {
		fmt.Printf("%s ----> %v \n", k , *v)
		
	}
}

func main() {

	var input bool = true;
	family := make(Family)
	

	for(input) {


		fmt.Println("Enter option \n",
		          "1. Add Family \n",
				  "2. Deprecate a Node \n",
				  "3. Print Family \n",
				  "4. Exit ")
		var option int

		fmt.Scanf("%d", &option)

		fmt.Println("Entered : ", option)
		switch option {

		case 1: 
		
		var familyName string
		var productName string
		fmt.Println("Enter Family Name : ")
		fmt.Scanf("\n%s", &familyName)
		var value *[]string
		var exists bool

		value , exists = family[familyName]

		fmt.Println("Enter Product Name :")
		fmt.Scanf("\n%s", &productName)

		if exists {

			productExists := slices.Contains(*value, productName)

			if productExists {
				fmt.Printf("%s already exists in %s \n", productName, familyName)
				} else {
				*value = append(*value, productName)	
			}
		} else {

			fmt.Printf("Map does not exist with product %s in %s  \n", productName, familyName)
			
			family[familyName] = makeSlice()
			*family[familyName] = append(*family[familyName], productName)
			

			//fmt.Println("Map: " , family ,  *family[familyName])
		}

		


	case 2 : 	

	var productName string
	fmt.Println("Deprecate a node by specifying the chassis :");
	fmt.Scanf("\n%s" , &productName)

	for _, v := range family {

		index := slices.Index(*v, productName)
		fmt.Printf("Index of %s is at %d \n", productName , index)
		removeElement(v, index)
		fmt.Printf("Family after removing %v \n",  *v)
	}
	case 4 : 
	    input = false

	case 3 : 
	    fmt.Println("Printing Family Map :")
		print(family)	
		}

		fmt.Scanf("\n")
	} 

}
