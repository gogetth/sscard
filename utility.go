package sscard

import "os"

// WriteBlockToFile i.e: write jpeg to file
func WriteBlockToFile(databytes []byte, fullname string) (int, error) {
	f, err := os.Create(fullname)
	if err != nil {
		return -1, err
	}

	defer f.Close()

	n, err := f.Write(databytes)
	//fmt.Printf("wrote %d bytes\n", n)
	return n, err
}
