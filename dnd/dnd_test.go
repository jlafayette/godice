package dnd

import (
	"fmt"
	"math/big"
	"reflect"
	"testing"
)

func TestAdvantage(t *testing.T) {
	a := Advantage()
	t.Run("DMap", func(t *testing.T) {
		expected := map[int]int{1: 1, 2: 3, 3: 5, 4: 7, 5: 9, 6: 11, 7: 13, 8: 15, 9: 17, 10: 19, 11: 21, 12: 23, 13: 25, 14: 27, 15: 29, 16: 31, 17: 33, 18: 35, 19: 37, 20: 39}
		actual := a.DMap()
		eq := reflect.DeepEqual(actual, expected)
		if !eq {
			t.Errorf("Got %d, expected %d", actual, expected)
		}
	})
	fmt.Println(a.Average())
	fmt.Println(a.AverageRat())
	fmt.Println(a.Rand(10))
}

// Useful for generating test data for maps
func generateDMapData(a DnDRoll) {
	m := a.DMap()
	fmt.Print("\nmap[int]int{")
	for i := 1; i < a.dice[0]+1; i++ {
		fmt.Printf("%d:%d, ", i, m[i])
	}
	fmt.Print("}\n")
}

// Useful for validating map against percentage output
func printPercentages() {
	a := Advantage()
	m := a.DMap()
	fmt.Print("\n #   %")
	for i := 1; i < 21; i++ {
		r := big.NewRat(int64(m[i]*100), int64(400))
		fmt.Printf("\n%2d %s", i, r.FloatString(2))
	}
}
