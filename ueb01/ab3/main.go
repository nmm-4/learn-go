package main

import "fmt"

/*
 * Aufgabe 3
 *
 * Schreiben sie ein Go Programm,
 * welches ein Schaltjahr überprüfen soll,
 *
 * Beispielausgabe:
 * 	   > Welches Jahr wollen sie überprüfen? 2020
 *     Das Jahr 2020 ist ein Schaltjahr.
 *
 *     > Welches Jahr wollen sie überprüfen? 2000
 *     Das Jahr 2000 ist ein Schaltjahr.
 *
 * 	   > Welches Jahr wollen sie überprüfen? 2025
 *     Das Jahr 2025 ist kein Schaltjahr.
 *
 *     > Welches Jahr wollen sie überprüfen? 2100
 *     Das Jahr 2100 ist kein Schaltjahr.
 *
 * Was ein Schaltjahr ist? https://de.wikipedia.org/wiki/Schaltjahr
 */
func main() {

	input := -1
	fmt.Print("Welches Jahr soll überprüft werden? ")
	fmt.Scan(&input)

	fmt.Print("Das Jahr ", input, " ist ")
	if input%400 == 0 {
		fmt.Print("ein")
	} else if input%100 == 0 {
		fmt.Print("kein")
	} else if input%4 == 0 {
		fmt.Print("ein")
	} else {
		fmt.Print("kein")
	}
	fmt.Println(" Schaltjahr")

}
