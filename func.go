package sscard

import (
	"fmt"
)

// Transmitter is an interface that wrap the command to communicate with smart card via application protocol data unit (APDU) according to ISO/IEC 7816.
type Transmitter interface {
	Transmit([]byte) ([]byte, error)
}

// APDUGetRsp Send list of APDU and get last command response
// ispadzeroOptional is optional(default = true) to replace adpu tail section
func APDUGetRsp(card Transmitter, apducmds [][]byte, ispadzeroOptional ...bool) ([]byte, error) {
	var resp []byte

	ispadzero := true
	if len(ispadzeroOptional) > 0 {
		ispadzero = ispadzeroOptional[0]
	}

	// Send command APDU: apducmds
	for _, apducmd := range apducmds {
		rsp, err := card.Transmit(apducmd)
		if err != nil {
			fmt.Println("Error Transmit:", err)
			return nil, err
		}
		//printRsp(rsp)
		resp = rsp
	}

	// pad zero
	if ispadzero == true {
		dlen := len(resp)
		resp[dlen-2] = 0
	}

	return resp, nil
}

// APDUGetBlockRsp Send list of APDU and append all response
func APDUGetBlockRsp(scardCard Transmitter, apducmds [][]byte, apducmdRsp []byte) ([]byte, error) {
	var respBlock []byte
	card := scardCard

	for _, apducmd1 := range apducmds {
		// Send command APDU: apducmd1
		rsp, err := card.Transmit(apducmd1)
		if err != nil {
			fmt.Println("Error Transmit:", err)
			return nil, err
		}
		// printRsp(rsp)

		// Send command APDU: apducmdRsp
		rsp, err = card.Transmit(apducmdRsp)
		if err != nil {
			fmt.Println("Error Transmit:", err)
			return nil, err
		}
		//printRsp(rsp)

		respBlock = append(respBlock, rsp[:len(rsp)-2]...)

	}
	// fmt.Printf("% 2X\n", respBlock)

	return respBlock, nil
}

func printRsp(rsp []byte) {
	fmt.Println("resp: ", rsp)
	for i := 0; i < len(rsp)-2; i++ {
		fmt.Printf("%c", rsp[i])
	}
	fmt.Println()
}
