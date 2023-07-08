package dog

import "strings"

func WhenGrownUp(s string) string {
	return "When the puppy grows up it says: " + strings.ToUpper(s)
}

//func returns "human years" in "dog years" (x7)
func Years(y int) int {
	return y * 7
}
