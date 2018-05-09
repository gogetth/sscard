package sscard

import (
	"encoding/hex"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func swapHex(hexs []byte) {
	for i, num := range hexs {
		num = ((num << 4) & 0xf0) | ((num >> 4) & 0x0f)
		hexs[i] = num
	}
}

// EncodeIMSI ...
func EncodeIMSI(imsi []byte) []byte {
	swapHex(imsi)
	return imsi
}

// DecodeIMSI ...
// swap hex in this format
// Source Hex : 59 02 30 59 20 57 03 30
// Result Hex : 95 20 03 95 02 57 30 03
func DecodeIMSI(ef []byte) []byte {
	if len(ef) < 4 {
		return nil
	}
	// lenght := ef[0]
	// fmt.Println(lenght)
	swapHex(ef)
	return ef
}

// EncodeICCID ...
func EncodeICCID(iccid []byte) []byte {
	if len(iccid) < 4 {
		return nil
	}
	lenght := iccid[0]
	fmt.Println(lenght)
	swapHex(iccid)
	return iccid
}

// DecodeICCID ...
func DecodeICCID(iccid []byte) []byte {
	swapHex(iccid)
	return iccid
}

// LuhnStringValid returns a boolean indicating if the argument was valid according to the Luhn algorithm.
func LuhnStringValid(luhnString string) bool {
	checksumMod := CalculateChecksum(luhnString, false) % 10

	return checksumMod == 0
}

// LuhnByteValid returns a boolean indicating if the argument was valid according to the Luhn algorithm.
func LuhnByteValid(luhnByte []byte, lenght int) bool {
	checksumMod := CalculateChecksumByte(luhnByte, false, lenght) % 10
	fmt.Println("checksumMod :", checksumMod)
	return checksumMod == 0
}

// Generate creates and returns a string of the length of the argument targetSize.
// The returned string is valid according to the Luhn algorithm.
func Generate(size int) string {
	random := randomString(size - 1)
	controlDigit := strconv.Itoa(generateControlDigit(random))

	return random + controlDigit
}

func generateControlDigit(luhnString string) int {
	controlDigit := CalculateChecksum(luhnString, true) % 10

	if controlDigit != 0 {
		controlDigit = 10 - controlDigit
	}

	return controlDigit
}

// CalculateChecksumByte ...
// For calculate to find checksum digit  must not include checksum digit
// For Validate -> include checksumdigit if sum % 10 = 0 ->> Valid
func CalculateChecksumByte(luhnByte []byte, double bool, lenght int) int {
	luhnString := []rune(hex.EncodeToString(luhnByte))
	if lenght > len(luhnString) {
		fmt.Errorf("Invalid Len")
		return -1
	}
	source := luhnString[:lenght]
	// source := luhnString
	return CalculateChecksum(string(source), double)
}

// CalculateChecksum ...
// For calculate to find checksum digit  must not include checksum digit
// For Validate -> include checksumdigit if sum % 10 = 0 ->> Valid
func CalculateChecksum(luhnString string, double bool) int {
	// Trim "f", it's has used as a filler
	luhnString = strings.Trim(luhnString, "f")
	source := strings.Split(luhnString, "")
	checksum := 0

	for i := len(source) - 1; i > -1; i-- {
		t, _ := strconv.ParseInt(source[i], 10, 8)
		n := int(t)

		if double {
			// fmt.Printf("Double %d\r\n", n)
			n = n * 2
		}
		double = !double

		if n >= 10 {
			n = n - 9
		}
		checksum += n
	}

	return checksum
}
func randomString(size int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	source := make([]int, size)

	for i := 0; i < size; i++ {
		source[i] = rand.Intn(9)
	}

	return integersToString(source)
}
func integersToString(integers []int) string {
	result := make([]string, len(integers))

	for i, number := range integers {
		result[i] = strconv.Itoa(number)
	}

	return strings.Join(result, "")
}
