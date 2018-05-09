package main

import (
	"fmt"

	"github.com/Napat/sscard/sscard"
	"github.com/ebfe/scard"
)

// ExampleSimCard ...

func exampleSimCard() {
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

	// APDUReadICCID
	iccid, err := sscard.APDUGetRsp(card, sscard.APDUReadICCID)
	if err != nil {
		fmt.Println("Error Transmit:", err)
		return
	}
	fmt.Printf("resp sscard.APDUReadICCID: % 02x\r\n", iccid)
	iccidDecode := sscard.DecodeICCID(iccid)
	fmt.Printf("ICCID Decode : % 02x Valid : %v\r\n\n", iccidDecode, sscard.LuhnByteValid(iccidDecode[:10], 20))

	imsi, err := sscard.APDUGetRsp(card, sscard.APDUReadIMSI)
	if err != nil {
		fmt.Println("Error Transmit:", err)
		return
	}
	fmt.Printf("resp sscard.APDUReadIMSI: % 02x\r\n", imsi)
	imsiDecode := sscard.DecodeIMSI(imsi)
	fmt.Printf("IMSI Decode : % 02x Valid : %v\r\n\n", imsiDecode, sscard.LuhnByteValid(imsiDecode[:9], 18))

	// APDUReadSMS
	sms, err := sscard.APDUGetRsp(card, sscard.APDUReadSMS)
	if err != nil {
		fmt.Println("Error Transmit:", err)
		return
	}
	fmt.Printf("resp sscard.APDUReadSMS: % 02x\r\n\n", sms)

	// APDUReadLOCI
	data, err := sscard.APDUGetRsp(card, sscard.APDUReadLOCI)
	if err != nil {
		fmt.Println("Error Transmit:", err)
		return
	}
	fmt.Printf("resp sscard.APDUReadLOCI: % 02x\r\n\n", data)

}
