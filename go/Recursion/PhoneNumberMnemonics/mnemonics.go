package main

var Letters = map[byte][]byte{
	'0': {'0'},
	'1': {'1'},
	'2': {'a', 'b', 'c'},
	'3': {'d', 'e', 'f'},
	'4': {'g', 'h', 'i'},
	'5': {'j', 'k', 'l'},
	'6': {'m', 'n', 'o'},
	'7': {'p', 'q', 'r', 's'},
	'8': {'t', 'u', 'v'},
	'9': {'w', 'x', 'y', 'z'},
}

func PhoneNumberMnemonics(phoneNumber string) []string {
	cMnemonic := make([]byte, len(phoneNumber))
	for i := range cMnemonic {
		cMnemonic[i] = '0'
	}

	// Result array which will be passed by reference.
	mnemonicsFound := make([]string, 0)

	// Helper which recursively runs
	phoneNumberMnemonicsHelper(0, phoneNumber, cMnemonic, &mnemonicsFound)

	return mnemonicsFound
}

func phoneNumberMnemonicsHelper(idx int, phoneNumber string, current []byte, found *[]string) {
	if idx == len(phoneNumber) {
		mnemonic := string(current)
		*found = append(*found, mnemonic)
	} else {
		// Everything happens here.
		digit := phoneNumber[idx]
		letters := Letters[digit]

		for _, letter := range letters {
			current[idx] = letter
			phoneNumberMnemonicsHelper(idx+1, phoneNumber, current, found)
		}
	}
}
