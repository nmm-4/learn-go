package main

import "fmt"

/*
 * Aufgabe 2
 *
 * Schreiben Sie ein Go-Programm,
 * welches prüft, ob drei Zahlen ein gültiges Dreieck bilden.
 *
 * Ein Dreieck ist gültig, wenn jede Seite kleiner als die Summe der anderen beiden ist.
 *
 * Beispielausgabe:
 *		> Wie lautet ihr Dreieck
 *		> 1. Seite: 3
 *		> 2. Seite: 4
 *		> 3. Seite: 5
 *		Die Seiten 3, 4, 5 bilden ein gültiges Dreieck.
 *
 *		> Wie lautet ihr Dreieck
 *		> 1. Seite: 1
 *		> 2. Seite: 2
 *		> 3. Seite: 3
 *		Die Seiten 1, 2, 3 bilden kein gültiges Dreieck.
 */

func main() {
	fmt.Println("Wie lautet ihr Dreieck?")
	x, y, z := 0, 0, 0
	fmt.Print("1. Seite ")
	fmt.Scan(&x)
	fmt.Print("2. Seite ")
	fmt.Scan(&y)
	fmt.Print("3. Seite ")
	fmt.Scan(&z)

	if x < y+z && y < x+z && z < x+y {
		fmt.Println("Die Seiten", x, y, z, "bilden ein gültiges Dreieck.")
	} else {
		fmt.Println("Die Seiten", x, y, z, "bilden kein gültiges Dreieck.")
	}
}
