package main

/*
	Aufgabe 2

	Dieses Programm dient dazu, ein einfaches Schachfeld in der
	Konsole darzustellen. Dabei werden für jede Schachfigur Konstanten
	definiert, die stellvertretend für diese Figuren stehen. Die Konstanten
	sind als uint64 Werte hinterlegt, die eine Art "Code" oder Identifikation
	für jede Figur darstellen:

	* KNIGHT für Springer
	* ROOK   für Turm
	* QUEEN  für Dame
	* KING   für König
	* BISHOP für Läufer
	* PAWN   für Bauer

	Ziel des Programms ist es, basierend auf diesen Konstanten
	ein Schachbrett zu generieren, auf dem jede Figur an ihrer
	Startposition angezeigt wird. Die Ausgabe erfolgt textuell in der Konsole,
	sodass man das Schachfeld mit seinen Figuren visuell erkennen kann.

	-------------------------------------------
	1. Bitboard-Prinzip für Schach

	Ein Schachbrett besteht aus 64 Feldern,
	die wir von 0 bis 63 durchnummerieren können.

	Bit 0 steht für Feld a1 (linke untere Ecke aus Weißersicht)
	Bit 1 steht für Feld b1
	...
	Bit 7 steht für Feld h1
	Bit 8 steht für Feld a2
	...
	Bit 63 steht für Feld h8 (rechte obere Ecke)

	Man zählt also von links nach rechts und von unten nach oben.

	------------------------------------------
	Beispiel der Figurenpositionen:

	KNIGHT uint64 = (1 << 1) | (1 << 6)
	ROOK   uint64 = (1 << 0) | (1 << 7)
	BISHOP uint64 = (1 << 2) | (1 << 5)
	QUEEN  uint64 = (1 << 3)
	KING   uint64 = (1 << 4)
	PAWN   uint64 = 0x000000000000FF00

	Schachbrett (textuelle Darstellung):

	8 · · · · · · · ·
	7 · · · · · · · ·
	6 · · · · · · · ·
	5 · · · · · · · ·
	4 · · · · · · · ·
	3 · · · · · · · ·
	2 ♙ ♙ ♙ ♙ ♙ ♙ ♙ ♙
	1 ♖ ♘ ♗ ♕ ♔ ♗ ♘ ♖
	  a b c d e f g h

	BONUS: Schachfeld im Array speichern
*/

const (
	KNIGHT uint64 = (1 << 1) | (1 << 6)
	ROOK   uint64 = (1 << 0) | (1 << 7)
	BISHOP uint64 = (1 << 2) | (1 << 5)
	QUEEN  uint64 = (1 << 3)
	KING   uint64 = (1 << 4)
	PAWN   uint64 = 0x000000000000FF00
)

func main() {
	// Hier kommt deine Logik zum Anzeigen des Schachfelds
}
