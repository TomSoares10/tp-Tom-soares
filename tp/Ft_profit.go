package tp

import (
	"math"
)

func Ft_profit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	profitMax := 0
	prixMin := math.MaxInt32

	// parcours le tableaudes prix
	for _, prix := range prices {
		// mettre à jour le prix minimum (si nécessaire)
		if prix < prixMin {
			prixMin = prix
		}
		// calculer le profit potentiel si on vend maintenant
		profit := prix - prixMin
		// mettre à jour le profit maximal
		if profit > profitMax {
			profitMax = profit
		}
	}

	return profitMax
}
