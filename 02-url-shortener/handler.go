package shortener

import (
	"math/rand/v2"
)

func Generate_Random_Int() (generated int){
	return rand.IntN(100)
}