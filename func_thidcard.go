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
func ThIDCardCID(card Transmiter, opt ...OptThIDCard) (string, error) {
	cfg := thidcardOpts{}
	for _, o := range opt {
		o(&cfg)
	}

	cid, err := APDUGetRsp(card, APDUThaiIDCardCID)
	if err != nil {
		return "", err
	}

	return mutateString(string(cid), cfg), nil
}

//ThIDCardFullnameEn get full name(English) from Thai national ID smart card.
func ThIDCardFullnameEn(card Transmiter, opt ...OptThIDCard) (string, error) {
	cfg := thidcardOpts{}
	for _, o := range opt {
		o(&cfg)
	}

	fullnameEN, err := APDUGetRsp(card, APDUThaiIDCardFullnameEn)
	if err != nil {
		return "", err
	}

	return mutateString(string(fullnameEN), cfg), nil
}

//ThIDCardFullnameTh get full name(English) from Thai national ID smart card.
func ThIDCardFullnameTh(card Transmiter, opt ...OptThIDCard) (string, error) {
	cfg := thidcardOpts{}
	for _, o := range opt {
		o(&cfg)
	}

	fullnameTH, err := APDUGetRsp(card, APDUThaiIDCardFullnameTh)
	if err != nil {
		return "", err
	}

	return mutateString(string(fullnameTH), cfg), nil
}

//ThIDCardBirth get birth date from Thai national ID smart card.
func ThIDCardBirth(card Transmiter, opt ...OptThIDCard) (string, error) {
	cfg := thidcardOpts{}
	for _, o := range opt {
		o(&cfg)
	}

	birth, err := APDUGetRsp(card, APDUThaiIDCardBirth)
	if err != nil {
		return "", err
	}

	return mutateString(string(birth), cfg), nil
}

//ThIDCardGender get gender from Thai national ID smart card.
func ThIDCardGender(card Transmiter, opt ...OptThIDCard) (string, error) {
	cfg := thidcardOpts{}
	for _, o := range opt {
		o(&cfg)
	}

	gender, err := APDUGetRsp(card, APDUThaiIDCardGender)
	if err != nil {
		return "", err
	}

	return mutateString(string(gender), cfg), nil
}

//ThIDCardIssuer get issuer from Thai national ID smart card.
func ThIDCardIssuer(card Transmiter, opt ...OptThIDCard) (string, error) {
	cfg := thidcardOpts{}
	for _, o := range opt {
		o(&cfg)
	}

	issuer, err := APDUGetRsp(card, APDUThaiIDCardIssuer)
	if err != nil {
		return "", err
	}

	return mutateString(string(issuer), cfg), nil
}

//ThIDCardIssueDate get issue date from Thai national ID smart card.
func ThIDCardIssueDate(card Transmiter, opt ...OptThIDCard) (string, error) {
	cfg := thidcardOpts{}
	for _, o := range opt {
		o(&cfg)
	}

	issuer, err := APDUGetRsp(card, APDUThaiIDCardIssuedate)
	if err != nil {
		return "", err
	}

	return mutateString(string(issuer), cfg), nil
}

//ThIDCardExpireDate get expire date from Thai national ID smart card.
func ThIDCardExpireDate(card Transmiter, opt ...OptThIDCard) (string, error) {
	cfg := thidcardOpts{}
	for _, o := range opt {
		o(&cfg)
	}

	expire, err := APDUGetRsp(card, APDUThaiIDCardExpiredate)
	if err != nil {
		return "", err
	}

	return mutateString(string(expire), cfg), nil
}

//ThIDCardAddress get address from Thai national ID smart card.
func ThIDCardAddress(card Transmiter, opt ...OptThIDCard) (string, error) {
	cfg := thidcardOpts{}
	for _, o := range opt {
		o(&cfg)
	}

	address, err := APDUGetRsp(card, APDUThaiIDCardAddress)
	if err != nil {
		return "", err
	}

	return mutateString(string(address), cfg), nil
}

func mutateString(s string, cfg thidcardOpts) string {
	if cfg.tis620ToUtf8 {
		s = tis620ToUtf8(s)
	}
	if cfg.sharpToSpace {
		s = sharpToSpace(s)
	}
	return s
}

func sharpToSpace(s string) string {
	s = strings.Replace(s, "#", " ", -1)
	return s
}

func tis620ToUtf8(s string) string {
	tis620Reader := bytes.NewBufferString(s)

	reader, err := charset.NewReaderLabel("TIS-620", tis620Reader)

	if err != nil {
		return ""
	}

	b, err := ioutil.ReadAll(reader)
	if err != nil {
		return ""
	}

	return string(b)
}
