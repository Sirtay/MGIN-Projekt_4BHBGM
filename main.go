package main

import "fmt"

func main() {
	
	fmt.Println("Kalorienumsatz Rechner von Mandic & Singh, Lab 7 Finales Projekt")

	//Benutzereingaben
	
	fmt.Print("Geben Sie Ihr Gewicht ein: ")

	var gewicht float32


	fmt.Scanln(&gewicht)
	fmt.Print("Geben Sie Ihr Alter ein: ")

	var alter float32


	fmt.Scanln(&alter)
	fmt.Print("Geben Sie Ihre Körpergröße ein (in cm): ")

	var groesse float32


	fmt.Scanln(&groesse)
	fmt.Print("Wie Aktiv sind sie von 1-10: ")

	var aktivitätsgrad float32
	var aktivität float32
	fmt.Scanln(&aktivitätsgrad)

	//Berechnungen

	if aktivitätsgrad<= 10 && aktivitätsgrad>= 1 && alter >= 16 && alter <= 100 && first >= 30 && groeße <= 250 && groeße >= 100 {

		aktivität = aktivitätsgrad

		var formel1 float32 = 655 + ((9.6 * gewicht + (1.8 * groeße) - (4.7 * alter))
		var formel2 float32 = aktivität * 200
		var formel3 float32 = formel1 + formel2
		fmt.Println("")

		fmt.Print("Ihr Grundumsatz beträgt: ")
		fmt.Println(formel1)
		fmt.Println("")

		fmt.Print("Ihr Zusatzumsatz durch Bewegung beträgt: ")
		fmt.Println(formel2)
		fmt.Println("")

		fmt.Print("Ihr Gesamtumsatz beträgt: ")
		fmt.Println(formel3)

	} else {

		fmt.Println("Bitte Eingabe beachten! ")



	}

}
