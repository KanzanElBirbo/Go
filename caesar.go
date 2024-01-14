package main

import (
	"fmt"
)

// Funkce caesarCipher provádí šifrování nebo dešifrování zprávy pomocí Caesarovy šifry.
// Přijímá vstupní řetězec a hodnotu posunu a vrací šifrovanou nebo dešifrovanou zprávu.
func caesarCipher(input string, shift int) string {
	result := ""
	for _, char := range input {
		// Ověření, zda znak patří do rozsahu malých písmen
		if char >= 'a' && char <= 'z' {
			// Aplikace posunu a normalizace na rozsah malých písmen
			result += string((char-'a'+rune(shift))%26 + 'a')
		} else if char >= 'A' && char <= 'Z' {
			result += string((char-'A'+rune(shift))%26 + 'A')
		} else {
			// Zahrnutí znaku, který není písmenem
			result += string(char)
		}
	}
	return result
}

func main() {
	// Původní zpráva
	message := "Hello, Golang!"
	// Hodnota posunu
	shift := 3

	// Šifrování zprávy
	encrypted := caesarCipher(message, shift)
	fmt.Println("Encrypted:", encrypted)

	// Dešifrování zprávy
	decrypted := caesarCipher(encrypted, -shift)
	fmt.Println("Decrypted:", decrypted)
}
