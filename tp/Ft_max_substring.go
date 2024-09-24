package tp

import (
	"fmt"
)

func Ft_max_substring(s string) int {
	// une carte pour stocker l'index des caractères
	dernierIndice := make(map[byte]int)
	longueurMax := 0
	debut := 0

	// parcours de la chaîne de caractères
	for fin := 0; fin < len(s); fin++ {
		caractere := s[fin]

		// si le caractère a déjà été vu et est dans la fenêtre actuelle
		if indice, existe := dernierIndice[caractere]; existe && indice >= debut {
			// déplacer le début juste après l'index du dernier caractère répété
			debut = indice + 1
		}

		// mettre à jour la position du caractère
		dernierIndice[caractere] = fin

		// calculer la longueur de la sous-chaîne actuelle sans répétition
		longueurActuelle := fin - debut + 1
		if longueurActuelle > longueurMax {
			longueurMax = longueurActuelle
		}
	}

	return longueurMax
}

func main() {
	fmt.Println(Ft_max_substring("abcabcbb")) // résultat : 3 ("abc")
	fmt.Println(Ft_max_substring("bbbbb"))    // résultat : 1 ("b")
	fmt.Println(Ft_max_substring("pwwkew"))   // résultat : 3 ("wke")
}
