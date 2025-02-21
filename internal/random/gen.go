package keygen

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

const alphabetUpper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const alphabetLow = "abcdefghijklmnopqrstuvwxyz"
const nums = "0123456789"
const specialChars = "!@#$%&*()_+-=,.^~?<>;:/{}[]|\\'\"`"

func GetRandomPass(size int, useUpper bool, useNum bool, useSpecial bool) string {
	if size < 4 {
		return "Error: password too short"
	}

	if size > 128 {
		size = 128
	}

	pass := []rune{}
	blob := alphabetLow
	if err := addRequiredChar(useUpper, &blob, alphabetUpper); err != nil {
		fmt.Println("Error: ", err)
	}
	if err := addRequiredChar(useNum, &blob, nums); err != nil {
		fmt.Println("Error: ", err)
	}
	if err := addRequiredChar(useSpecial, &blob, specialChars); err != nil {
		fmt.Println("Error: ", err)
	}

	useAtLeastOneChar(useUpper, alphabetUpper, &pass)
	useAtLeastOneChar(useNum, nums, &pass)
	useAtLeastOneChar(useSpecial, specialChars, &pass)

	for len(pass) < size {
		pass = append(pass, rune(blob[rand.Intn(len(blob))]))
	}
	rand.Shuffle(len(pass), func(i, j int) {
		pass[i], pass[j] = pass[j],
			pass[i]
	})

	return string(pass)
}

func addRequiredChar(use bool, char *string, requiredChar string) error {
	if char == nil {
		return errors.New("string pointer is null")
	}
	if use {
		*char += requiredChar
	}
	return nil
}

func useAtLeastOneChar(use bool, requiredChar string, pass *[]rune) {
	seed := time.Now().UnixNano()
	rand.NewSource(seed)
	if use {
		*pass = append(*pass, rune(requiredChar[rand.Intn(len(requiredChar))]))
	}
}
