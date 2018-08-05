# Super Smart Card(sscard)

Super Smart Card API on top of scard(pcsc handler) with apdu commands.

## Builtin APDU devices

- Thai ID card (public data)
- Simcard (public data)

## Installation and requirements

See `docs/INSTALLATION_xxx.md`

## TODO

``` bash
# Linux: install pcsc library
sudo apt-get install pcscd

# goget
go get -u github.com/Napat/sscard

# go build example
go build -o sscard github.com/Napat/sscard/main

./sscard  # ./main.exe on windows(if no -o sscard)
```

## Platforms

- Windows 10
- Linux: Ubuntu, Raspbian stretch(RPi3B)

## References

- [PCSC in golang](https://ludovicrousseau.blogspot.fr/2016/09/pcsc-sample-in-go.html)
- [APDU command for Thai ID card](https://github.com/Napat/ThaiNationalIDCard/blob/master/APDU.md)
- [Auto start pcscd using systemd](https://ludovicrousseau.blogspot.com/2011/11/pcscd-auto-start-using-systemd.html)
