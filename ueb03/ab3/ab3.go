package main

/*
    Aufgabe 3
    ----------
    In dieser Aufgabe sollen Arrays und Slices geübt werden.
    Es werden zu jeder Unteraufgabe eine Zahlenfolge vom Benutzer
    gelesen und dementsprechend erfolgt die Verarbeitung + Ausgabe.

    TIPP:
        Es kann beim aktiven Entwickeln auch eine Konstante definiert
        werden, die den Benutzer simuliert, um das Testen zu vereinfachen.

    In jeder Unteraufgabe wird das Sortieren gebraucht. Dabei soll 
    für jede Unteraufgabe ein anderer Algorithmus verwendet werden, z.B.:

        3.1: Bubble Sort
        3.2: Selection Sort
        3.3: Insertion Sort

    Hier sind vier Sortieralgorithmen, die verwendet werden sollen:
    (Um den größtmöglichen Lernerfolg zu haben, wird ausdrücklich empfohlen,
    keinen Python-Code in Go-Code zu übersetzen, sondern die Algorithmen
    vom Grund auf zu verstehen und den Code selbst zu verfassen.)

    - (leicht)   Bubble Sort
      https://www.w3schools.com/dsa/dsa_algo_bubblesort.php

    - (mittel)   Selection Sort
      https://www.w3schools.com/dsa/dsa_algo_selectionsort.php

    - (mittel)   Insertion Sort
      https://www.w3schools.com/dsa/dsa_algo_insertionsort.php

    - (schwer)   Shell Sort (bester Lernerfolg und mein Liebling)
	
    Aufgaben:
    ---------

    3.1)
        Die Zahlenfolge vom Benutzer soll sortiert werden.
        Beispiel:
            [0, 17, 256, 7, 15] -> [0, 7, 15, 17, 256]

    3.2)
        Nach Quersummen sortieren.
        
        TIPP: 
            Um die Quersumme einer Zahl zu berechnen:
            - Die Ziffern einzeln extrahieren:
              Zahl % 10 = letzte Ziffer
            - Beispiel 93:
                93 % 10 = 3
                9  % 10 = 9
              Quersumme = 9 + 3 = 12

            - Beispiel 123:
                123 % 10 = 3
                12  % 10 = 2
                1   % 10 = 1
              Quersumme = 1 + 2 + 3 = 6

        Beispielsortierung nach Quersummen:
            [0, 17, 256, 7, 15] -> [0, 15, 7, 17, 256]

            Quersummen:
            0   = 0
            17  = 8  (1 + 7)
            256 = 13 (2 + 5 + 6)
            7   = 7
            15  = 6  (1 + 5)

    3.3)
        Sortierung nach der Anzahl an Einsen in der Binärdarstellung.
        
        Beispiel:
            [0, 17, 256, 7, 15] -> [0, 256, 17, 7, 15]

            Binärdarstellungen (vereinfacht):
            0   = 000000000 (0 Einsen)
            256 = 100000000 (1 Eins)
            17  = 000010001 (2 Einsen)
            7   = 000000111 (3 Einsen)
            15  = 000001111 (4 Einsen)
*/

func main() {
    const aufgabe = 1

    switch aufgabe {
    case 1:
        {
            // TODO Aufgabe 3.1: Bubble Sort implementieren
        }
    case 2:
        {
            // TODO Aufgabe 3.2: Selection Sort mit Quersummen als Sortierkriterium implementieren
        }
    case 3:
        {
            // TODO Aufgabe 3.3: Shell Sort oder Insertion Sort nach Anzahl der Einsen in Binärdarstellung implementieren
        }
    }
}
