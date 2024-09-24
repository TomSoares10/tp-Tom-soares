package main

import (
	"fmt"
	"tp/tp"
)

func main() {
	// Exercice 1: Ft_coin
	fmt.Println("Exercice 1: Ft_coin")
	fmt.Println(tp.Ft_coin([]int{1, 2, 5}, 11)) // résultat : 3
	fmt.Println(tp.Ft_coin([]int{2}, 3))        // résultat : -1
	fmt.Println(tp.Ft_coin([]int{1}, 0))        // résultat : 0

	// Exercice 2: Ft_missing
	fmt.Println("Exercice 2: Ft_missing")
	fmt.Println(tp.Ft_missing([]int{3, 1, 2}))                   // résultat : 0
	fmt.Println(tp.Ft_missing([]int{0, 1}))                      // résultat : 2
	fmt.Println(tp.Ft_missing([]int{9, 6, 4, 2, 3, 5, 7, 0, 1})) // résultat : 8

	// Exercice 3: Ft_non_overlap
	fmt.Println("Exercice 3: Ft_non_overlap")
	fmt.Println(tp.Ft_non_overlap([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}})) // résultat : 1
	fmt.Println(tp.Ft_non_overlap([][]int{{1, 2}, {2, 3}}))                 // résultat : 0
	fmt.Println(tp.Ft_non_overlap([][]int{{1, 2}, {1, 2}, {1, 2}}))         // résultat : 2

	// Exercice 4: Ft_profit
	fmt.Println("Exercice 4: Ft_profit")
	fmt.Println(tp.Ft_profit([]int{7, 1, 5, 3, 6, 4})) // résultat : 5
	fmt.Println(tp.Ft_profit([]int{7, 6, 4, 3, 1}))    // résultat : 0

	// Exercice 5: Ft_max_substring
	fmt.Println("Exercice 5: Ft_max_substring")
	fmt.Println(tp.Ft_max_substring("abcabcbb")) // résultat : 3
	fmt.Println(tp.Ft_max_substring("bbbbb"))    // résultat : 1

	// Exercice 6: Ft_min_window
	fmt.Println("Exercice 6: Ft_min_window")
	fmt.Println(tp.Ft_min_window("ADOBECODEBANC", "ABC")) // résultat : "BANC"
	fmt.Println(tp.Ft_min_window("a", "aa"))              // résultat : ""
}
