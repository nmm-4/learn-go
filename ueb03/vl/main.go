
package main

import "fmt"


func waitUntil() {
	foo := " "
	fmt.Scanln(&foo)
}








func main() {
	{
		a := [5]int{10, 20, 30, 30, 50}
		fmt.Println(cap(a),len(a))

		waitUntil()
		s := a[1:5]     

		fmt.Println(s)
		waitUntil()

		s = append(s, 10)

		fmt.Println(s)
		waitUntil()
		s[0]++

		fmt.Println(a)
	}
	waitUntil()
	fmt.Println("---------------------------------------")
	{
		a := [5]int{10, 20, 30, 30, 50}
		fmt.Println(cap(a),len(a))

		waitUntil()
		s := a[1:4]     

		fmt.Println(s)
		waitUntil()

		s = append(s, 10)

		fmt.Println(s)
		waitUntil()

		fmt.Println(a)
	}
	waitUntil()
	fmt.Println("---------------------------------------")
	{
		a := [5]int{10, 20, 30, 30, 50}
		fmt.Println(cap(a),len(a))

		waitUntil()
		s := a[1:4:4]     

		fmt.Println(s)
		waitUntil()

		s = append(s, 10)

		fmt.Println(s)
		waitUntil()

		fmt.Println(a)
	}
	waitUntil()
	fmt.Println("---------------------------------------")
	{
		words := [6]string{"Go", "ist", "eine", "coole", 
		"Programmiersprache", "!"}

		slice1 := words[1:4] 
		fmt.Println("slice1:", slice1)


		slice2 := append(slice1, "und", "schnell")  
		fmt.Println("slice2:", slice2)

		fmt.Println("words:", words)


		slice3 := slice2[2:5] 
		fmt.Println("slice3:", slice3)

		slice3[0] = "fantastische"
		fmt.Println("slice3 nach Änderung:", slice3)
		fmt.Println("slice2 nach Änderung:", slice2)
		fmt.Println("words nach Änderung:", words)
	}
	waitUntil()
	fmt.Println("---------------------------------------")
	{
		words := [6]string{"Go", "ist", "eine", "coole", "Programmiersprache", "!"}

		slice1 := words[1:4] 
		fmt.Println("slice1:", slice1)


		slice2 := append(slice1, "und", "schnell", "stabil")  
		fmt.Println("slice2:", slice2)

		fmt.Println("words:", words)


		slice3 := slice2[2:5] 
		fmt.Println("slice3:", slice3)

		slice3[0] = "fantastische"
		fmt.Println("slice3 nach Änderung:", slice3)
		fmt.Println("slice2 nach Änderung:", slice2)
		fmt.Println("words nach Änderung:", words)
	}
}

