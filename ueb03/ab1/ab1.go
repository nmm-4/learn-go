package main

import "fmt"

/*

Aufgabe 1

Schreiben sie ein Programm welches ein Ganzzahl (uint16) 
in ein Binärzahl umwandelt.

Der Fall das der Benutzer ein Zahl eingibt, welches grösser ist als 
uint16 soll nicht überprüft werden.

Beispiel: 

> Zahl: 0
> 0000 0000 0000 0000 

> Zahl: 10
> 0000 0000 0000 1010 

> Zahl: 255
> 0000 0000 1111 1111 
*/
func main() {
	fmt.Println("Welche Zahl wollen sie in Binärdarstellen?")
		

	var x uint64
	fmt.Scan(&x)
	
	fmt.Print("Ihre Zahl lautet: ")
	fmt.Printf("%b \n", x)

	for i:= 0; i < 64; i++ {
		if x & 0x8000000000000000 == 0 {
			fmt.Print("0")
		} else {
			fmt.Print("1")
		}
		x <<= 1
	}
	fmt.Print("\n")

}

