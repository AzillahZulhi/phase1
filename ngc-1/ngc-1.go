package main

import "fmt"

func main() {

	// variable 1
	var mynum int32 = 50
	var mynum2 float32 = 51
	var mynumstr string = "50"
	fmt.Println(mynum)
	fmt.Println(mynum2)
	fmt.Println(mynumstr)

	// variable 2
	var x, y int32
	x = 5
	y = 10
	z := x + y
	fmt.Println(z)

	// CLI
	var name string = "bambang"
	fmt.Println("hello", name)

	// array dan slice 1
	people := []string{"Walt,", "jesse,", "skyler,", "saul,"}
	fmt.Println("panjang slice people ==>", len(people))
	people = append(people, "hank,", "marie")
	fmt.Println("panjang slice people setelah di tambahkan ==>", len(people))
	fmt.Println(people)

	// array dan slice 2
	var people2 []map[string]string
	people2 = append(people2, map[string]string{"name": "hank", "gender": "m"})
	people2 = append(people2, map[string]string{"name": "heisenberg", "gender": "m"})
	people2 = append(people2, map[string]string{"name": "skyler", "gender": "f"})

	fmt.Println("aray sebelum modifikasi")

	for _, person := range people2 {
		fmt.Println(person)
	}
	for i, person := range people2 {
		gender := person["gender"]
		if gender == "f" {
			people2[i]["name"] = "Mrs " + person["name"]
		} else if gender == "m" {
			people2[i]["name"] = "Mr " + person["name"]
		}
	}
	fmt.Println("\naray setelah modifikasi")
	for _, person := range people2 {
		fmt.Println(person)
	}
}
