package main

import (
	"fmt"
	"math/rand"
)

/*
 * Aufgabe 4
 *
 * Bevor Sie mit der Aufgabe starten, lesen Sie bitte den Artikel bis Abschnitt 5:
 * https://golangdocs.com/generate-random-numbers-in-golang
 * 
 * HINWEIS:
 * Leider verwendet der Artikel eine veraltete API, daher ist das Nötigste bereits vorgegeben.
 *
 * Aufgabenstellung:
 * - Vom Benutzer soll ein SEED eingelesen werden, welcher für die Zufallszahlengenerierung benötigt wird.
 *   (Dieser Teil ist vorgegeben.)
 * - Es sollen Zufallszahlen im Bereich von 0 bis 100 generiert werden.
 * - Insgesamt sollen N Zufallszahlen erzeugt und in einem Array gespeichert werden.
 * - Anschließend sollen Mittelwert, Median und Summe der generierten Zahlen berechnet werden.
 * - Abschließend erfolgt die Ausgabe der Ergebnisse.
 *
 * Beispiel:
 * N = 6
 * > Seed: 42
 * [5 87 68 50 23 45]
 * Mittelwert: 46,33
 * Median:     47,5
 * Summe:      278
 * --------------------------
 * N = 5
 * > Seed: 42
 * [5 87 68 50 23]
 * Mittelwert: 46,6
 * Median:     50
 * Summe:      233
 */

const N int = 5

func main() {
	var seed int64
	fmt.Print("Seed: ")
	fmt.Scan(&seed)
	generator := rand.New(rand.NewSource(seed))
	// generator mit Seed initialisiert, verwenden Sie generator.Int() usw., um Zufallszahlen zu erzeugen.
}
