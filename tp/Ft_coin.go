package tp

func Ft_coin(coins []int, amount int) int {
	// initialise le nombre de piece minimal
	min := amount + 1

	// fonction pour calculer le nombre de piece necesaire
	var necess func(int, int)
	necess = func(amount int, count int) {
		// Si le montant atteint 0, on a trouvé une combinaison possible
		if amount == 0 {
			if count < min {
				min = count
			}
			return
		}

		// Si le montant est inférieur à 0, ce n'est pas une solution valide
		if amount < 0 {
			return
		}

		// Essayer chaque pièce
		for _, coin := range coins {
			necess(amount-coin, count+1)
		}
	}

	// Appel de la fonction pour chercher le nombre minimal de pièces
	necess(amount, 0)

	// Si aucune solution n'a été trouvée, retourner -1
	if min == amount+1 {
		return -1
	}

	return min
}
