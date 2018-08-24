package main

import (
	"fmt"

	"github.com/ebfe/scard"
	"github.com/gogetth/sscard"
)

// exampleThaiIDCard ...
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
	fmt.Printf("cid: _%s_\n", string(cid))

	fullnameEN, err := sscard.APDUGetRsp(card, sscard.APDUThaiIDCardFullnameEn)
	if err != nil {
		fmt.Println("Error APDUGetRsp: ", err)
		return
	}
	fmt.Printf("fullnameEN: _%s_\n", string(fullnameEN))

	fullnameTH, err := sscard.APDUGetRsp(card, sscard.APDUThaiIDCardFullnameTh)
	if err != nil {
		fmt.Println("Error APDUGetRsp: ", err)
		return
	}
	fmt.Printf("fullnameTH: _%s_\n", string(fullnameTH))

	birth, err := sscard.APDUGetRsp(card, sscard.APDUThaiIDCardBirth)
	if err != nil {
		fmt.Println("Error APDUGetRsp: ", err)
		return
	}
	fmt.Printf("birth: _%s_\n", string(birth))

	gender, err := sscard.APDUGetRsp(card, sscard.APDUThaiIDCardGender)
	if err != nil {
		fmt.Println("Error APDUGetRsp: ", err)
		return
	}
	fmt.Printf("gender: _%s_\n", string(gender))

	issuer, err := sscard.APDUGetRsp(card, sscard.APDUThaiIDCardIssuer)
	if err != nil {
		fmt.Println("Error APDUGetRsp: ", err)
		return
	}
	fmt.Printf("issuer: _%s_\n", string(issuer))

	issueDate, err := sscard.APDUGetRsp(card, sscard.APDUThaiIDCardIssuedate)
	if err != nil {
		fmt.Println("Error APDUGetRsp: ", err)
		return
	}
	fmt.Printf("issueDate: _%s_\n", string(issueDate))

	issueExp, err := sscard.APDUGetRsp(card, sscard.APDUThaiIDCardExpiredate)
	if err != nil {
		fmt.Println("Error APDUGetRsp: ", err)
		return
	}
	fmt.Printf("issueExp: _%s_\n", string(issueExp))

	address, err := sscard.APDUGetRsp(card, sscard.APDUThaiIDCardAddress)
	if err != nil {
		fmt.Println("Error APDUGetRsp: ", err)
		return
	}
	fmt.Printf("address: _%s_\n", string(address))

	cardPhotoJpg, err := sscard.APDUGetBlockRsp(card, sscard.APDUThaiIDCardPhoto, sscard.APDUThaiIDCardPhotoRsp)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Printf("Image binary: % 2X\n", cardPhotoJpg)

	n2, err := sscard.WriteBlockToFile(cardPhotoJpg, "./idcPhoto.jpg")
	if err != nil {
		fmt.Println("Error WriteBlockToFile: ", err)
		return
	}
	fmt.Printf("wrote %d bytes\n", n2)
}

// exampleThaiIDCard2 ...
func exampleThaiIDCard2() {

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

	cid, err := sscard.ThIDCardCID(card)
	if err != nil {
		fmt.Println("Error APDUGetRsp: ", err)
		return
	}
	fmt.Printf("cid: _%s_\n", string(cid))

	fullnameEN, err := sscard.ThIDCardFullnameEn(card)
	if err != nil {
		fmt.Println("Error APDUGetRsp: ", err)
		return
	}
	fmt.Printf("fullnameEN: _%s_\n", string(fullnameEN))

	fullnameTH, err := sscard.ThIDCardFullnameTh(card, sscard.OptTis620ToUtf8())
	if err != nil {
		fmt.Println("Error APDUGetRsp: ", err)
		return
	}
	fmt.Printf("fullnameTH: _%s_\n", string(fullnameTH))

	birth, err := sscard.ThIDCardBirth(card)
	if err != nil {
		fmt.Println("Error APDUGetRsp: ", err)
		return
	}
	fmt.Printf("birth: _%s_\n", string(birth))

	gender, err := sscard.ThIDCardGender(card)
	if err != nil {
		fmt.Println("Error APDUGetRsp: ", err)
		return
	}
	fmt.Printf("gender: _%s_\n", string(gender))

	issuer, err := sscard.ThIDCardIssuer(card)
	if err != nil {
		fmt.Println("Error APDUGetRsp: ", err)
		return
	}
	fmt.Printf("issuer: _%s_\n", string(issuer))

	issueDate, err := sscard.ThIDCardIssueDate(card)
	if err != nil {
		fmt.Println("Error APDUGetRsp: ", err)
		return
	}
	fmt.Printf("issue date: _%s_\n", string(issueDate))

	expire, err := sscard.ThIDCardExpireDate(card)
	if err != nil {
		fmt.Println("Error APDUGetRsp: ", err)
		return
	}
	fmt.Printf("expire date: _%s_\n", string(expire))

	address, err := sscard.ThIDCardAddress(card, sscard.OptTis620ToUtf8(), sscard.OptSharpToSpace())
	if err != nil {
		fmt.Println("Error APDUGetRsp: ", err)
		return
	}
	fmt.Printf("address: _%s_\n", string(address))

	cardPhotoJpg, err := sscard.APDUGetBlockRsp(card, sscard.APDUThaiIDCardPhoto, sscard.APDUThaiIDCardPhotoRsp)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Printf("Image binary: % 2X\n", cardPhotoJpg)

	n2, err := sscard.WriteBlockToFile(cardPhotoJpg, "./idcPhoto.jpg")
	if err != nil {
		fmt.Println("Error WriteBlockToFile: ", err)
		return
	}
	fmt.Printf("wrote %d bytes\n", n2)

}

func printbytes(rsp []byte) {
	for i := 0; i < len(rsp)-2; i++ {
		fmt.Printf("%c", rsp[i])
	}
	fmt.Println()
}
