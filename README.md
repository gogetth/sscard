# Super Smart Card(sscard)

Super Smart Card API on top of pcsc.

## TODO

``` bash
# Linux: install pcsc library
sudo apt-get install pcscd

# goget
go get -u github.com/Napat/sscard/sscard

# go build
go build -o sscard github.com/Napat/sscard/main

./sscard  # ./sscard.exe on windows
```

About requirement and other platform see: `docs/INSTALLATION_xxx.md`

# go run hack
go run $(find ./ | grep ./main/)

## Platforms

- Windows 10
- Linux / Ubuntu

## Reference

- [PCSC in golang](https://ludovicrousseau.blogspot.fr/2016/09/pcsc-sample-in-go.html)
- [APDU command for Thai ID card](https://github.com/Napat/ThaiNationalIDCard/blob/master/APDU.md)
