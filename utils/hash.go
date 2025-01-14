package utils

import (
	"fmt"
	"hash/fnv"
)

func HashStringToInt32(input string) int32 {
	hasher := fnv.New32a()
	_, err := hasher.Write([]byte(input))
	if err != nil {
		panic(fmt.Sprintf("Failed to hash string: %v", err))
	}
	return int32(hasher.Sum32())
}
