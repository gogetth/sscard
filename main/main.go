
package main

import (
	"fmt"
 	"os"

	"github.com/ebfe/scard"
	"apduthaiidcard"
)

func main() {
	// Establish a PC/SC context
	context, err := scard.EstablishContext()
	if err != nil {
		fmt.Println("Error EstablishContext:", err)
		return
	}

	// Release the PC/SC context (when needed)
	defer context.Release()

	// List available readers
	readers, err := context.ListReaders()
	if err != nil {
		fmt.Println("Error ListReaders:", err)
		return
	}

	// Use the first reader
	reader := readers[0]
	fmt.Println("Using reader:", reader)

	// Connect to the card
	card, err := context.Connect(reader, scard.ShareShared, scard.ProtocolAny)
	if err != nil {
		fmt.Println("Error Connect:", err)
		return
	}

	// Disconnect (when needed)
	defer card.Disconnect(scard.LeaveCard)

	// Send select APDU
	rsp, err := card.Transmit(apduthaiidcard.ApduThaiIdCardSelect)
	if err != nil {
		fmt.Println("Error Transmit:", err)
		return
	}
	fmt.Println("resp apduthaiidcard.ApduThaiIdCardSelect: ", rsp)	

	cardTxAPDU(card, apduthaiidcard.ApduThaiIdCardCid1, apduthaiidcard.ApduThaiIdCardCid2)
	cardTxAPDU(card, apduthaiidcard.ApduThaiIdCardFullnameEn1, apduthaiidcard.ApduThaiIdCardFullnameEn2)
	cardTxAPDU(card, apduthaiidcard.ApduThaiIdCardFullnameTh1, apduthaiidcard.ApduThaiIdCardFullnameTh2)
	cardTxAPDU(card, apduthaiidcard.ApduThaiIdCardBirth1, apduthaiidcard.ApduThaiIdCardBirth2)
	cardTxAPDU(card, apduthaiidcard.ApduThaiIdCardGender1, apduthaiidcard.ApduThaiIdCardGender2)
	cardTxAPDU(card, apduthaiidcard.ApduThaiIdCardIssuer1, apduthaiidcard.ApduThaiIdCardIssuer2)
	cardTxAPDU(card, apduthaiidcard.ApduThaiIdCardIssuedate1, apduthaiidcard.ApduThaiIdCardIssuedate2)
	cardTxAPDU(card, apduthaiidcard.ApduThaiIdCardExpiredate1, apduthaiidcard.ApduThaiIdCardExpiredate2)
	cardTxAPDU(card, apduthaiidcard.ApduThaiIdCardAddress1, apduthaiidcard.ApduThaiIdCardAddress2)

	cardPhoto(card)
}

func cardTxAPDU(card *scard.Card, apducmd1 []byte, apducmd2 []byte){		
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

func cardPhoto(card *scard.Card){
	var cardimg = []byte{}
	apducmd2 := apduthaiidcard.ApduThaiIdCardPhotoGetResp
	for _ , apducmd1 := range apduthaiidcard.ApduThaiIdCardPhoto {
		fmt.Println(apducmd1)
		fmt.Println(apducmd2)

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
			cardimg = append(cardimg, rsp[i])			
		}		 
	}
	fmt.Println("Card image ")
	fmt.Printf("% 2X\n", cardimg)


	////////////////////
	f, err := os.Create("./dat2.jpg")
    check(err)

	n2, err := f.Write(cardimg)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}
