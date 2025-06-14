package main

import "fmt"

/*
* Abschlussaufgabe
*
* Schreiben Sie ein Go-Programm,
* das einen Betrag in Euro und Cent in Scheine und Münzen umrechnet.
*
* Dabei soll der Betrag mit möglichst wenigen Scheinen und Münzen ausgegeben werden.
*
* Die Eingabe wird von dem Benutzer gemacht.
*
* Beispiel 1:
* > Wie viel Euro wollen sie? 87
* > Wie viel Cent wollen sie? 36
* 	50€-Schein: 1
*	20€-Schein: 1
*	10€-Schein: 1
*	 5€-Schein: 1
*  ------------------
*	  2€-Münze: 1
*	20ct-Münze: 1
*	10ct-Münze: 1
*	 5ct-Münze: 1
*	 1ct-Münze: 1

* Beispiel 2:
* > Wie viel Euro wollen sie? 888
* > Wie viel Cent wollen sie? 0

*  500€-Schein: 1
*  200€-Schein: 1
*  100€-Schein: 1
*   50€-Schein: 1
*   20€-Schein: 1
*   10€-Schein: 1
*    5€-Schein: 1
*  ------------------
*     2€-Münze: 1
* 	  1€-Münze: 1

* Beispiel 3:
* > Wie viel Euro wollen sie? 0
* > Wie viel Cent wollen sie? 200
*

	(Keine Scheine nötig)

*  ------------------
*     2€-Münze: 1

* Beispiel 4:
* > Wie viel Euro wollen sie? 0
* > Wie viel Cent wollen sie? 0
*

	(Keine Scheine nötig)

*  ------------------
*    (Keine Münzen nötig)
*/
func main() {

	euro, cent := 0, 0

	fmt.Print("Wie viel Euro wollen sie? ")
	fmt.Scan(&euro)
	fmt.Print("Wie viel Cent wollen sie? ")
	fmt.Scan(&cent)

	euro += cent / 100
	cent %= 100

	tmpEuro := euro

	if euro/500 > 0 {
		fmt.Println("500€-Schein :", euro/500)
		euro %= 500
	}

	if euro/200 > 0 {
		fmt.Println("200€-Schein :", euro/200)
		euro %= 200
	}

	if euro/100 > 0 {
		fmt.Println("100€-Schein :", euro/100)
		euro %= 100
	}

	if euro/50 > 0 {
		fmt.Println(" 50€-Schein :", euro/50)
		euro %= 50
	}

	if euro/20 > 0 {
		fmt.Println(" 20€-Schein :", euro/20)
		euro %= 20
	}

	if euro/10 > 0 {
		fmt.Println(" 10€-Schein :", euro/10)
		euro %= 10
	}

	if euro/5 > 0 {
		fmt.Println("  5€-Schein :", euro/5)
		euro %= 5
	}
	if tmpEuro == euro {
		fmt.Println("(Keine Scheine nötig)")
	}
	fmt.Println("------------------")

	tmpEuro = euro
	tmpCent := cent
	if euro/2 > 0 {
		fmt.Println("  2€-Münze  :", euro/2)
		euro %= 2
	}
	if euro > 0 {
		fmt.Println("  1€-Münze  :", euro)
	}

	if cent/50 > 0 {
		fmt.Println("  50 Cent   :", cent/50)
		cent %= 50
	}

	if cent/20 > 0 {
		fmt.Println("  20 Cent   :", cent/20)
		cent %= 20
	}

	if cent/10 > 0 {
		fmt.Println("  10 Cent   :", cent/10)
		cent %= 10
	}

	if cent/5 != 0 {
		fmt.Println("   5 Cent   :", cent/5)
		cent %= 5
	}

	if cent/2 != 0 {
		fmt.Println("   2 Cent   :", cent/2)
		cent %= 2
	}
	if cent > 0 {
		fmt.Println("   1 Cent   :", cent)
	}
	if tmpEuro == 0 && cent == tmpCent {
		fmt.Println("(Keine Scheine nötig)")
	}
}
