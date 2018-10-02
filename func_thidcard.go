package sscard

import (
	"bytes"
	"io/ioutil"
	"strings"

	"golang.org/x/net/html/charset"
)

type thidcardOpts struct {
	sharpToSpace bool
	tis620ToUtf8 bool
}

// OptThIDCard provide options for each operation.
type OptThIDCard func(*thidcardOpts)

// OptSharpToSpace add option to convert all # to space before return a string
func OptSharpToSpace() OptThIDCard {
	return func(cfg *thidcardOpts) {
		cfg.sharpToSpace = true
	}
}

// OptTis620ToUtf8 add option to convert TIS620 string back to UTF-8 before return a string
func OptTis620ToUtf8() OptThIDCard {
	return func(cfg *thidcardOpts) {
		cfg.tis620ToUtf8 = true
	}
}

//ThIDCardCID get cid from Thai national ID smart card.
func ThIDCardCID(card Transmitter, opt ...OptThIDCard) (string, error) {
	cfg := thidcardOpts{}
	for _, o := range opt {
		o(&cfg)
	}

	cid, err := APDUGetRsp(card, APDUThaiIDCardCID)
	if err != nil {
		return "", err
	}

	return thIDCardMutateString(string(cid), cfg)
}

//ThIDCardFullnameEn get full name(English) from Thai national ID smart card.
func ThIDCardFullnameEn(card Transmitter, opt ...OptThIDCard) (string, error) {
	cfg := thidcardOpts{}
	for _, o := range opt {
		o(&cfg)
	}

	fullnameEN, err := APDUGetRsp(card, APDUThaiIDCardFullnameEn)
	if err != nil {
		return "", err
	}

	return thIDCardMutateString(string(fullnameEN), cfg)
}

//ThIDCardFullnameTh get full name(Thai) from Thai national ID smart card.
func ThIDCardFullnameTh(card Transmitter, opt ...OptThIDCard) (string, error) {
	cfg := thidcardOpts{}
	for _, o := range opt {
		o(&cfg)
	}

	fullnameTH, err := APDUGetRsp(card, APDUThaiIDCardFullnameTh)
	if err != nil {
		return "", err
	}

	return thIDCardMutateString(string(fullnameTH), cfg)
}

//ThIDCardBirth get birth date from Thai national ID smart card.
func ThIDCardBirth(card Transmitter, opt ...OptThIDCard) (string, error) {
	cfg := thidcardOpts{}
	for _, o := range opt {
		o(&cfg)
	}

	birth, err := APDUGetRsp(card, APDUThaiIDCardBirth)
	if err != nil {
		return "", err
	}

	return thIDCardMutateString(string(birth), cfg)
}

//ThIDCardGender get gender from Thai national ID smart card.
func ThIDCardGender(card Transmitter, opt ...OptThIDCard) (string, error) {
	cfg := thidcardOpts{}
	for _, o := range opt {
		o(&cfg)
	}

	gender, err := APDUGetRsp(card, APDUThaiIDCardGender)
	if err != nil {
		return "", err
	}

	return thIDCardMutateString(string(gender), cfg)
}

//ThIDCardIssuer get issuer from Thai national ID smart card.
func ThIDCardIssuer(card Transmitter, opt ...OptThIDCard) (string, error) {
	cfg := thidcardOpts{}
	for _, o := range opt {
		o(&cfg)
	}

	resp, err := APDUGetRsp(card, APDUThaiIDCardIssuer)
	resp = []byte(strings.Replace(string(resp), " ", "", -1))
	resp = bytes.Replace(resp, []byte("\000"), nil, -1)
	if err != nil {
		return "", err
	}
	issuer := resp
	issuer = rmBackslashBytes(issuer)

	return thIDCardMutateString(string(issuer), cfg)
}

//ThIDCardIssueDate get issue date from Thai national ID smart card.
func ThIDCardIssueDate(card Transmitter, opt ...OptThIDCard) (string, error) {
	cfg := thidcardOpts{}
	for _, o := range opt {
		o(&cfg)
	}

	issuer, err := APDUGetRsp(card, APDUThaiIDCardIssuedate)
	if err != nil {
		return "", err
	}

	return thIDCardMutateString(string(issuer), cfg)
}

//ThIDCardExpireDate get expire date from Thai national ID smart card.
func ThIDCardExpireDate(card Transmitter, opt ...OptThIDCard) (string, error) {
	cfg := thidcardOpts{}
	for _, o := range opt {
		o(&cfg)
	}

	expire, err := APDUGetRsp(card, APDUThaiIDCardExpiredate)
	if err != nil {
		return "", err
	}

	return thIDCardMutateString(string(expire), cfg)
}

//ThIDCardAddress get address from Thai national ID smart card.
func ThIDCardAddress(card Transmitter, opt ...OptThIDCard) (string, error) {
	cfg := thidcardOpts{}
	for _, o := range opt {
		o(&cfg)
	}

	address, err := APDUGetRsp(card, APDUThaiIDCardAddress)
	if err != nil {
		return "", err
	}

	return thIDCardMutateString(string(address), cfg)
}

func thIDCardMutateString(s string, cfg thidcardOpts) (str string, err error) {
	str = s
	if cfg.tis620ToUtf8 {
		str, err = tis620ToUtf8(str)
		if err != nil {
			return str, err
		}
	}
	if cfg.sharpToSpace {
		str, err = sharpToSpace(str)
		if err != nil {
			return str, err
		}
	}
	return str, nil
}

// rmBackslashBytes remove backslash from []byte
func rmBackslashBytes(s []byte) []byte {
	return bytes.Replace(s, []byte("\\"), nil, -1)
}

func sharpToSpace(s string) (string, error) {
	s = strings.Replace(s, "#", " ", -1)
	return s, nil
}

func tis620ToUtf8(s string) (string, error) {
	tis620Reader := bytes.NewBufferString(s)

	reader, err := charset.NewReaderLabel("TIS-620", tis620Reader)

	if err != nil {
		return "", err
	}

	b, err := ioutil.ReadAll(reader)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
