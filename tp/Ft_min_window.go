package tp

import (
	"math"
)

func Ft_min_window(s string, t string) string {
	if len(t) > len(s) {
		return ""
	}

	// compte des caractères requis de t
	need := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	// variables pour la fenêtre glissante
	count := 0
	left := 0
	start := 0
	minLen := math.MaxInt32
	have := make(map[byte]int)

	// parcours de la chaîne s
	for right := 0; right < len(s); right++ {
		char := s[right]

		// si le caractère fait partie de t, on l'ajoute au compteur "have"
		if need[char] > 0 {
			have[char]++
			// si on a exactement le bon nombre de ce caractère, on augmente le compteur "count"
			if have[char] <= need[char] {
				count++
			}
		}

		// quand on a tous les caractères requis
		for count == len(t) {
			// Mettre à jour la longueur minimale et la position de départ
			if right-left+1 < minLen {
				minLen = right - left + 1
				start = left
			}

			// réduire la fenêtre en enlevant le caractère à gauche
			leftChar := s[left]
			if need[leftChar] > 0 {
				have[leftChar]--
				// si on n'a plus assez de ce caractère dans la fenêtre, on réduit "count"
				if have[leftChar] < need[leftChar] {
					count--
				}
			}
			left++
		}
	}

	// si aucune fenêtre valide n'a été trouvée, retourner une chaîne vide
	if minLen == math.MaxInt32 {
		return ""
	}

	// retourner la plus petite fenêtre
	return s[start : start+minLen]
}
