package tp

import (
	"sort"
)

func Ft_non_overlap(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

	// trier les intervalles par la fin de l'intervalle
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	count := 0
	dernierFin := intervals[0][1]

	//  une boucle pour parcourir les intervalles tries
	for i := 1; i < len(intervals); i++ {
		// si l'intervalle actuel commence avant que le précédent se termine, il se chevauche
		if intervals[i][0] < dernierFin {
			count++ // on doit retirer cet intervalle
		} else {
			// Sinon on met à jour la fin du dernier intervalle non chevauché
			dernierFin = intervals[i][1]
		}
	}

	return count
}
