package main

func main() {
	w := func() {
		println("Je suis une fonction anonyme sans return")
	}
	w()

	z := func() string {
		println("Je suis une fonction anonyme.")
		return "Rony"
	}() // le () est nécessaire, sinon ça te renvoie un truc bizarre

	println(z)

	println("-----------")

	x, y := func() (int, int) {

		println("Aucun paramètre. Deux returns")
		return 5, 7
	}() // on appelle la fonction, comme pour z

	println(x)
	print(y)

	func(a, b int) {
		println("\nA * A + B * B = ", a*a+b*b)
	}(x, y) // tu utilises les valeurs de la fonction d'avant
}
