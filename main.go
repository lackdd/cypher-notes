package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func encrypt_pairs(s string) string {
	var pairs string
	CharMap := map[rune]rune{
		'a': 'b',
		'b': 'a',
		'c': 'd',
		'd': 'c',
		'e': 'f',
		'f': 'e',
		'g': 'h',
		'h': 'g',
		'i': 'j',
		'j': 'i',
		'k': 'l',
		'l': 'k',
		'm': 'n',
		'n': 'm',
		'o': 'p',
		'p': 'o',
		'q': 'r',
		'r': 'q',
		's': 't',
		't': 's',
		'u': 'v',
		'v': 'u',
		'w': 'x',
		'x': 'w',
		'y': 'z',
		'z': 'y',
		'A': 'B',
		'B': 'A',
		'C': 'D',
		'D': 'C',
		'E': 'F',
		'F': 'E',
		'G': 'H',
		'H': 'G',
		'I': 'J',
		'J': 'I',
		'K': 'L',
		'L': 'K',
		'M': 'N',
		'N': 'M',
		'O': 'P',
		'P': 'O',
		'Q': 'R',
		'R': 'Q',
		'S': 'T',
		'T': 'S',
		'U': 'V',
		'V': 'U',
		'W': 'X',
		'X': 'W',
		'Y': 'Z',
		'Z': 'Y',
	}

	for _, char := range s {
		if char >= 'A' && char <= 'Z' {
			pairs += string(CharMap[char])
		} else if char >= 'a' && char <= 'z' {
			pairs += string(CharMap[char])
		} else {
			pairs += string(char)
		}
	}
	return ("Encrypted your message using pairs:\n") + pairs
}

func decrypt_pairs(s string) string {
	var pairs string
	CharMap := map[rune]rune{
		'a': 'b',
		'b': 'a',
		'c': 'd',
		'd': 'c',
		'e': 'f',
		'f': 'e',
		'g': 'h',
		'h': 'g',
		'i': 'j',
		'j': 'i',
		'k': 'l',
		'l': 'k',
		'm': 'n',
		'n': 'm',
		'o': 'p',
		'p': 'o',
		'q': 'r',
		'r': 'q',
		's': 't',
		't': 's',
		'u': 'v',
		'v': 'u',
		'w': 'x',
		'x': 'w',
		'y': 'z',
		'z': 'y',
		'A': 'B',
		'B': 'A',
		'C': 'D',
		'D': 'C',
		'E': 'F',
		'F': 'E',
		'G': 'H',
		'H': 'G',
		'I': 'J',
		'J': 'I',
		'K': 'L',
		'L': 'K',
		'M': 'N',
		'N': 'M',
		'O': 'P',
		'P': 'O',
		'Q': 'R',
		'R': 'Q',
		'S': 'T',
		'T': 'S',
		'U': 'V',
		'V': 'U',
		'W': 'X',
		'X': 'W',
		'Y': 'Z',
		'Z': 'Y',
	}

	for _, char := range s {
		if char >= 'A' && char <= 'Z' {
			pairs += string(CharMap[char])
		} else if char >= 'a' && char <= 'z' {
			pairs += string(CharMap[char])
		} else {
			pairs += string(char)
		}
	}
	return ("Decrypted your message using Pairs:\n") + pairs
}

func encrypt_rot13(s string) string {
	var sonum string
	//makes slice of runes, aitab neid runnida, runid oleks see input, pikkusega string s

	for _, r := range s {
		//while loop, kus r on rune, pikkusega s
		switch {
		//erinevad olukorrad mis toimuda saavad (väike a, suur A või suva hyrogliif)
		case r >= 'a' && r <= 'z': //Kui ta on väike täht
			sonum += string((r-'a'+13)%26 + 'a')
		//update (sonum), rune r - 97 + (a-n=13) ja modula 26, et kokku loopiks neid tähti
		// + rune 'a', et teha jääk + 97, andes talle uue tähe Please enter 'e' to encrypt or 'd' to decrypt.n-ist edasi.
		case r >= 'A' && r <= 'Z': //Kui ta on suur täht
			sonum += string((r-'A'+13)%26 + 'A')
		default: //Kui ei ole miski??
			sonum = (sonum + string(r))
			//append, et panna uksteise kulge panema
		}

	}
	return ("Encrypted your message using ROT13:\n") + sonum //printed sõnum
}

func decrypt_rot13(s string) string {
	var sonum string
	//makes slice of runes, aitab neid runnida, runid oleks see input, pikkusega string s

	for _, r := range s {
		//while loop, kus r on rune, pikkusega s
		switch {
		//erinevad olukorrad mis toimuda saavad (väike a, suur A või suva hyrogliif)
		case r >= 'a' && r <= 'z': //Kui ta on väike täht
			sonum += string((r-'a'+13)%26 + 'a')
		//update (sonum), rune r - 97 + (a-n=13) ja modula 26, et kokku loopiks neid tähti
		// + rune 'a', et teha jääk + 97, andes talle uue tähe n-ist edasi.
		case r >= 'A' && r <= 'Z': //Kui ta on suur täht
			sonum += string((r-'A'+13)%26 + 'A')
		default: //Kui ei ole miski??
			sonum = (sonum + string(r))
			//append, et panna uksteise kulge panema
		}

	}
	return ("Decrypted your message using ROT13:\n") + sonum //printed sõnum
}

func encrypt_reverse(s string) string {
	var newStr string
	for _, ch := range s {
		switch {
		//erinevad olukorrad mis toimuda saavad (väike a, suur A või suva hyrogliif)
		case ch >= 'a' && ch <= 'z': //Kui ta on väike täht
			newStr += string('z' - (ch-'a')%26)
		//update (sonum), rune r - 97 + (a-n=13) ja modula 26, et kokku loopiks neid tähti
		// + rune 'a', et teha jääk + 97, andes talle uue tähe Please enter 'e' to encrypt or 'd' to decrypt.n-ist edasi.
		case ch >= 'A' && ch <= 'Z': //Kui ta on suur täht
			newStr += string('z' - (ch-'a')%26)
		default: //Kui ei ole miski??
			newStr = (newStr + string(ch))

		}
		//newStr += string('z' - (ch - 'a')%26 )
	}
	return ("Encrypted your message using reverse:\n") + newStr
}
func decrypt_reverse(s string) string {
	var newStr string
	for _, ch := range s {
		switch {
		case ch >= 'a' && ch <= 'z': //Kui ta on väike täht
			newStr += string('z' - (ch-'a')%26)
		case ch >= 'A' && ch <= 'Z': //Kui ta on suur täht
			newStr += string('z' - (ch-'a')%26)
		default: //Kui ei ole miski??
			newStr = (newStr + string(ch))

		}
		//newStr += string('z' - (ch-'a')%26)
	}
	return ("Decrypted your message using reverse:\n") + newStr
}

func getInput() (toEncrypt bool, encoding string, message string) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Hello User, this tool will help you encrypt or decrypt your message.")

	for {
		fmt.Print("Select Operation (1/2)\n 1.Encrypt\n 2.Decrypt\n")

		scanner.Scan()
		operation := strings.TrimSpace(scanner.Text())
		if operation == "1" {
			toEncrypt = true
			break
		} else if operation == "2" {
			toEncrypt = false
			break
		} else {
			fmt.Println("Invalid input:", operation)
		}
	}

	for {
		fmt.Print("Select method: \n 1.Rot13\n 2.Reverse \n 3.Pairs \n")
		scanner.Scan()
		encoding = strings.TrimSpace(scanner.Text())
		if encoding == "1" || encoding == "2" || encoding == "3" {
			break
		} else {
			fmt.Println("Invalid input:", encoding)
		}
	}

	fmt.Print("Enter your message: ")
	scanner.Scan()
	message = strings.TrimSpace(scanner.Text())

	return
}
func main() {
	toEncrypt, encoding, message := getInput()

	var result string
	if toEncrypt {
		switch encoding {
		case "1":
			result = encrypt_rot13(message)
		case "2":
			result = encrypt_reverse(message)
		case "3":
			result = encrypt_pairs(message)
		}
	} else {
		switch encoding {
		case "1":
			result = decrypt_rot13(message)
		case "2":
			result = decrypt_reverse(message)
		case "3":
			result = decrypt_pairs(message)
		}
	}

	fmt.Println(result + ("\nThank you for using our services!"))
}
