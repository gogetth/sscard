
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

	apduthaiidcard.CardTxAPDU(card, apduthaiidcard.ApduThaiIdCardCid1, apduthaiidcard.ApduThaiIdCardCid2)
	apduthaiidcard.CardTxAPDU(card, apduthaiidcard.ApduThaiIdCardFullnameEn1, apduthaiidcard.ApduThaiIdCardFullnameEn2)
	apduthaiidcard.CardTxAPDU(card, apduthaiidcard.ApduThaiIdCardFullnameTh1, apduthaiidcard.ApduThaiIdCardFullnameTh2)
	apduthaiidcard.CardTxAPDU(card, apduthaiidcard.ApduThaiIdCardBirth1, apduthaiidcard.ApduThaiIdCardBirth2)
	apduthaiidcard.CardTxAPDU(card, apduthaiidcard.ApduThaiIdCardGender1, apduthaiidcard.ApduThaiIdCardGender2)
	apduthaiidcard.CardTxAPDU(card, apduthaiidcard.ApduThaiIdCardIssuer1, apduthaiidcard.ApduThaiIdCardIssuer2)
	apduthaiidcard.CardTxAPDU(card, apduthaiidcard.ApduThaiIdCardIssuedate1, apduthaiidcard.ApduThaiIdCardIssuedate2)
	apduthaiidcard.CardTxAPDU(card, apduthaiidcard.ApduThaiIdCardExpiredate1, apduthaiidcard.ApduThaiIdCardExpiredate2)
	apduthaiidcard.CardTxAPDU(card, apduthaiidcard.ApduThaiIdCardAddress1, apduthaiidcard.ApduThaiIdCardAddress2)

	cardPhotoJpg, err := apduthaiidcard.CardPhoto(card)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Image binary: % 2X\n", cardPhotoJpg)

	// write jpeg to file
	f, err := os.Create("./dat2.jpg")
    check(err)

	n2, err := f.Write(cardPhotoJpg)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

