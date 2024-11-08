package eaglesong

import (
	"encoding/hex"
	"testing"
)

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i+1 < j-1; i, j = i+2, j-2 {
		temp1, temp2 := runes[i], runes[j]
		runes[i], runes[j] = runes[j-1], runes[i+1]
		runes[i+1], runes[j-1] = temp2, temp1
	}
	return string(runes)
}

func TestHashForEx(t *testing.T) {
	powHash := "f06644f799ef77ba8a8668bf73f71bc097ee849ab8d5f907df9d55b93b88755e"
	nonce := "3CC6717EBA981CFD0000000000000100"
	bz, _ := hex.DecodeString(powHash + reverseString(nonce))
	hash := EaglesongHash(bz)
	println(hex.EncodeToString(hash))
}

func TestHash(t *testing.T) {
	powHash := "555c3f633a011298f2636d3ed6d16a459069dd5c4bbf2e8689bc33940ac1f70b"
	nonce_str := "28AFB3BD5414C7CE0000000800000100"

	bz, _ := hex.DecodeString(powHash + nonce_str)
	hash := EaglesongHash(bz)
	println(hex.EncodeToString(hash))
}
