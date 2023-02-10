package main

import (
	"errors"
	"fmt"
)

func sayHello(name string) {
	fmt.Printf("Bonjour à tous, je m'appelle %v. \n", name)
}
func calculatePerimRect(lo, la int) int {
	return 2 * (lo + la)
	//calcule le périmètre d'un rectangle.
}
func sayBye(name string) (string, error) {
	if name == "" {
		return "", errors.New("Erreur : L'utilisateur n'a pas spécifié de nom !")
		// renvoie msg d'erreur si on spécifie pas de nom
	}
	byeMessage := fmt.Sprintf("%v s'en va ! Bonne soirée.\n", name)
	return byeMessage, nil //nil pour return le msg d'erreur au cas où
	//Message d'au revoir,
}

func main() {
	sayHello("Rony")
	r1 := calculatePerimRect(5, 2) // stockage du résultat du calcul ds la var r1
	fmt.Printf("Le périmètre de mon rectangle est de : %v.\n", r1)
	fmt.Printf(sayBye("")) // tu appelles sayBye, et tu lui dis que le name = Rony
	//(tu mets saybye entre () avec un printf pcq snion ça s'affiche pas dans la console.)
}
