package apduthaiidcard

import(
	"fmt"
	"os"

	"github.com/ebfe/scard"	
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

// CardTxAPDU ispadzeroOptional is optional(default = true) to replace adpu tail section  
func CardTxAPDU(scardCard *scard.Card, apducmd1 []byte, apducmd2 []byte, ispadzeroOptional ...bool) ([]byte, error) {
	ispadzero := true
	if len(ispadzeroOptional) > 0 {
		ispadzero = ispadzeroOptional[0]
	}
	
	card := scardCard

	// Send command APDU: apducmd1
	rsp, err := card.Transmit(apducmd1)
	if err != nil {
		fmt.Println("Error Transmit:", err)
		return nil, err
	}
	// fmt.Println("resp1: ", rsp)
	// for i := 0; i < len(rsp)-2; i++ {
	// 	fmt.Printf("%c", rsp[i])
	// }
	// fmt.Println() 

	// Send command APDU: apducmd2
	rsp, err = card.Transmit(apducmd2)
	if err != nil {
		fmt.Println("Error Transmit:", err)
		return nil, err
	}
	// fmt.Printf("resp2: %T: %v\n", rsp, rsp)
	// for i := 0; i < len(rsp)-2; i++ {
	// 	fmt.Printf("%c", rsp[i])
	// }
	// fmt.Println() 

	if ispadzero == true {
		dlen := len(rsp)
		rsp[dlen-2] = 0
	}

	return rsp, nil
}

func CardPhoto(scardCard *scard.Card) ([]byte, error){
	var cardPhotoJpg []byte
	card := scardCard

	apducmd2 := ApduThaiIdCardPhotoGetResp
	for _ , apducmd1 := range ApduThaiIdCardPhoto {
		// Send command APDU: apducmd1
		rsp, err := card.Transmit(apducmd1)
		if err != nil {
			fmt.Println("Error Transmit:", err)
			return nil, err
		}
		// fmt.Println("resp1: ", rsp)
		// for i := 0; i < len(rsp)-2; i++ {
		// 	fmt.Printf("%c", rsp[i])
		// }
		// fmt.Println() 

		// Send command APDU: apducmd2
		rsp, err = card.Transmit(apducmd2)
		if err != nil {
			fmt.Println("Error Transmit:", err)
			return nil, err
		}
		//fmt.Println("resp2: ", rsp)
		
		cardPhotoJpg = append(cardPhotoJpg, rsp[:len(rsp)-2]...)
		//for i := 0; i < len(rsp)-2; i++ {			
		//	cardPhotoJpg = append(cardPhotoJpg, rsp[i])			
		//}		 
	}
	// fmt.Println("Card image ")
	// fmt.Printf("% 2X\n", cardPhotoJpg)

	return cardPhotoJpg, nil
}

// WritePhotoToFile write jpeg to file
func WritePhotoToFile(jpgbytes []byte, fullname string) (int, error) {
	f, err := os.Create(fullname)
    if err != nil {
        return -1, err
    }
	
	defer f.Close()

	n, err := f.Write(jpgbytes)
	//fmt.Printf("wrote %d bytes\n", n)
	return n, err
}
