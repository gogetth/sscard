[![Go Report Card](https://goreportcard.com/badge/github.com/gogetth/sscard)](https://goreportcard.com/report/github.com/gogetth/sscard)
[![GoDoc](https://godoc.org/github.com/gogetth/sscard?status.svg)](https://godoc.org/github.com/gogetth/sscard)

# Super Smart Card(sscard)

Super Smart Card API on top of scard(pcsc handler) with apdu commands.

## Builtin APDU devices

- Thai ID card (public data)

## TODO

``` bash
# Linux: install pcsc library
sudo apt-get install pcscd

# goget
go get -u golang.org/x/net/html/charset
go get -u github.com/gogetth/sscard

# go build example
go build -o sscard github.com/gogetth/sscard/main

./sscard
```

## References

- [Thank you for original sscard work from Napat](https://github.com/Napat/sscard)
- [PCSC in golang](https://ludovicrousseau.blogspot.fr/2016/09/pcsc-sample-in-go.html)
- [APDU command for Thai ID card](https://github.com/Napat/ThaiNationalIDCard/blob/master/APDU.md)
- [Auto start pcscd using systemd](https://ludovicrousseau.blogspot.com/2011/11/pcscd-auto-start-using-systemd.html)
