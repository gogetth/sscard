package main

import (
	"fmt"

	"github.com/Napat/sscard/sscard"
	"github.com/ebfe/scard"
)

// ExampleThaiIDCard ...
func exampleThaiIDCard() {

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
	selectRsp, err := sscard.APDUGetRsp(card, sscard.APDUThaiIDCardSelect)
	if err != nil {
		fmt.Println("Error Transmit:", err)
		return
	}
	fmt.Println("resp sscard.APDUThaiIDCardSelect: ", selectRsp)

	cid, err := sscard.APDUGetRsp(card, sscard.APDUThaiIDCardCID)
	if err != nil {
		fmt.Println("Error APDUGetRsp: ", err)
		return
	}
	fmt.Printf("cid: %s\n", string(cid))

	fullnameEN, err := sscard.APDUGetRsp(card, sscard.APDUThaiIDCardFullnameEn)
	if err != nil {
		fmt.Println("Error APDUGetRsp: ", err)
		return
	}
	fmt.Printf("fullnameEN: %s\n", string(fullnameEN))

	fullnameTH, err := sscard.APDUGetRsp(card, sscard.APDUThaiIDCardFullnameTh)
	if err != nil {
		fmt.Println("Error APDUGetRsp: ", err)
		return
	}
	fmt.Printf("fullnameTH: %s\n", string(fullnameTH))

	birth, err := sscard.APDUGetRsp(card, sscard.APDUThaiIDCardBirth)
	if err != nil {
		fmt.Println("Error APDUGetRsp: ", err)
		return
	}
	fmt.Printf("birth: %s\n", string(birth))

	gender, err := sscard.APDUGetRsp(card, sscard.APDUThaiIDCardGender)
	if err != nil {
		fmt.Println("Error APDUGetRsp: ", err)
		return
	}
	fmt.Printf("gender: %s\n", string(gender))

	issuer, err := sscard.APDUGetRsp(card, sscard.APDUThaiIDCardIssuer)
	if err != nil {
		fmt.Println("Error APDUGetRsp: ", err)
		return
	}
	fmt.Printf("issuer: %s\n", string(issuer))

	issueDate, err := sscard.APDUGetRsp(card, sscard.APDUThaiIDCardIssuedate)
	if err != nil {
		fmt.Println("Error APDUGetRsp: ", err)
		return
	}
	fmt.Printf("issueDate: %s\n", string(issueDate))

	issueExp, err := sscard.APDUGetRsp(card, sscard.APDUThaiIDCardExpiredate)
	if err != nil {
		fmt.Println("Error APDUGetRsp: ", err)
		return
	}
	fmt.Printf("issueExp: %s\n", string(issueExp))

	address, err := sscard.APDUGetRsp(card, sscard.APDUThaiIDCardAddress)
	if err != nil {
		fmt.Println("Error APDUGetRsp: ", err)
		return
	}
	fmt.Printf("address: %s\n", string(address))

	cardPhotoJpg, err := sscard.APDUGetBlockRsp(card, sscard.APDUThaiIDCardPhoto, sscard.APDUThaiIDCardPhotoRsp)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Image binary: % 2X\n", cardPhotoJpg)

	n2, err := sscard.WriteBlockToFile(cardPhotoJpg, "./idcPhoto.jpg")
	fmt.Printf("wrote %d bytes\n", n2)
}

func printbytes(rsp []byte) {
	for i := 0; i < len(rsp)-2; i++ {
		fmt.Printf("%c", rsp[i])
	}
	fmt.Println()
}
