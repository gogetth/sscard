
package main

import (
	"fmt"

	"github.com/ebfe/scard"
	"github.com/Napat/thaiidcard_pcscapdu/apduthaiidcard"
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

	cid, err := apduthaiidcard.CardTxAPDU(card, apduthaiidcard.ApduThaiIdCardCid1, apduthaiidcard.ApduThaiIdCardCid2)
	if err != nil {
		fmt.Println("Error CardTxAPDU: ", err)
		return 
	}
	fmt.Printf("cid: %s\n", string(cid))

	fullnameEN, err := apduthaiidcard.CardTxAPDU(card, apduthaiidcard.ApduThaiIdCardFullnameEn1, apduthaiidcard.ApduThaiIdCardFullnameEn2)
	if err != nil {
		fmt.Println("Error CardTxAPDU: ", err)
		return 
	}
	fmt.Printf("fullnameEN: %s\n", string(fullnameEN))

	fullnameTH, err := apduthaiidcard.CardTxAPDU(card, apduthaiidcard.ApduThaiIdCardFullnameTh1, apduthaiidcard.ApduThaiIdCardFullnameTh2)
	if err != nil {
		fmt.Println("Error CardTxAPDU: ", err)
		return 
	}
	fmt.Printf("fullnameTH: %s\n", string(fullnameTH))

	birth, err := apduthaiidcard.CardTxAPDU(card, apduthaiidcard.ApduThaiIdCardBirth1, apduthaiidcard.ApduThaiIdCardBirth2)
	if err != nil {
		fmt.Println("Error CardTxAPDU: ", err)
		return 
	}
	fmt.Printf("birth: %s\n", string(birth))

	gender, err := apduthaiidcard.CardTxAPDU(card, apduthaiidcard.ApduThaiIdCardGender1, apduthaiidcard.ApduThaiIdCardGender2)
	if err != nil {
		fmt.Println("Error CardTxAPDU: ", err)
		return 
	}
	fmt.Printf("gender: %s\n", string(gender))

	issuer, err := apduthaiidcard.CardTxAPDU(card, apduthaiidcard.ApduThaiIdCardIssuer1, apduthaiidcard.ApduThaiIdCardIssuer2)
	if err != nil {
		fmt.Println("Error CardTxAPDU: ", err)
		return 
	}
	fmt.Printf("issuer: %s\n", string(issuer))

	issueDate, err := apduthaiidcard.CardTxAPDU(card, apduthaiidcard.ApduThaiIdCardIssuedate1, apduthaiidcard.ApduThaiIdCardIssuedate2)
	if err != nil {
		fmt.Println("Error CardTxAPDU: ", err)
		return 
	}
	fmt.Printf("issueDate: %s\n", string(issueDate))

	issueExp, err := apduthaiidcard.CardTxAPDU(card, apduthaiidcard.ApduThaiIdCardExpiredate1, apduthaiidcard.ApduThaiIdCardExpiredate2)
	if err != nil {
		fmt.Println("Error CardTxAPDU: ", err)
		return 
	}
	fmt.Printf("issueExp: %s\n", string(issueExp))

	address, err := apduthaiidcard.CardTxAPDU(card, apduthaiidcard.ApduThaiIdCardAddress1, apduthaiidcard.ApduThaiIdCardAddress2)
	if err != nil {
		fmt.Println("Error CardTxAPDU: ", err)
		return 
	}
	fmt.Printf("address: %s\n", string(address))

	cardPhotoJpg, err := apduthaiidcard.CardPhoto(card)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Image binary: % 2X\n", cardPhotoJpg)

	n2, err := apduthaiidcard.WritePhotoToFile(cardPhotoJpg, "./dat2.jpg")
	fmt.Printf("wrote %d bytes\n", n2)
}

func printbytes(rsp []byte) {
	for i := 0; i < len(rsp)-2; i++ {
		fmt.Printf("%c", rsp[i])
	}
	fmt.Println()
}
