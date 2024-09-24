package tp

func Ft_missing(nums []int) int {
	n := len(nums)                   // longueur de la liste
	sommeAttendue := n * (n + 1) / 2 // somme attendue pour les nombres de 0 à n
	sommeActuelle := 0

	// calcule la somme des nombres dans le tableau
	for _, nombre := range nums {
		sommeActuelle += nombre
	}

	// le nombre qui manque est la différence entre la somme attendue et la somme actuelle
	return sommeAttendue - sommeActuelle
}
