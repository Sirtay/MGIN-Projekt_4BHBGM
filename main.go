package main

import "fmt"

func main() {
	
	fmt.Println("Kalorienumsatz Rechner von Mandic & Singh, Lab 7 Finales Projekt")
	fmt.Print("Geben Sie Ihr Gewicht ein: ")

	var first float32


	fmt.Scanln(&first)
	fmt.Print("Geben Sie Ihr Alter ein: ")

	var alter float32


	fmt.Scanln(&alter)
	fmt.Print("Geben Sie Ihre Körpergröße ein (in cm): ")

	var groeße float32


	fmt.Scanln(&groeße)
	fmt.Print("Wie Aktiv sind sie von 1-10: ")

	var second float32
	var aktivität float32
	fmt.Scanln(&second)

	if second <= 10 && second >= 1 && alter >= 16 && alter <= 100 && first >= 30 && groeße <= 250 && groeße >= 100 {

		aktivität = second

		var formel float32 = 655 + ((9.6 * first) + (1.8 * groeße) - (4.7 * alter))
		var formel2 float32 = aktivität * 200
		var formel3 float32 = formel + formel2
		fmt.Println("")

		fmt.Print("Ihr Grundumsatz beträgt: ")
		fmt.Println(formel)
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
