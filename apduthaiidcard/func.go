package apduthaiidcard

import(
	"fmt"

	"github.com/ebfe/scard"	
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func CardTxAPDU(scardCard *scard.Card, apducmd1 []byte, apducmd2 []byte){
	card := scardCard

	// Send command APDU: apducmd1
	rsp, err := card.Transmit(apducmd1)
	if err != nil {
		fmt.Println("Error Transmit:", err)
		return
	}
	fmt.Println("resp1: ", rsp)
	for i := 0; i < len(rsp)-2; i++ {
		fmt.Printf("%c", rsp[i])
	}
	fmt.Println() 

	// Send command APDU: apducmd2
	rsp, err = card.Transmit(apducmd2)
	if err != nil {
		fmt.Println("Error Transmit:", err)
		return
	}
	fmt.Println("resp2: ", rsp)
	for i := 0; i < len(rsp)-2; i++ {
		fmt.Printf("%c", rsp[i])
	}
	fmt.Println() 
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
	fmt.Println("Card image ")
	fmt.Printf("% 2X\n", cardPhotoJpg)

	return cardPhotoJpg, nil
}

