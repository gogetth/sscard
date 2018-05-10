# Raspberry Pi 3 model B

``` bash
$ uname -a
Linux raspberrypi 4.14.34-v7+ #1110 SMP Mon Apr 16 15:18:51 BST 2018 armv7l GNU/Linux
$
$ lsb_release -a
No LSB modules are available.
Distributor ID:	Raspbian
Description:	Raspbian GNU/Linux 9.4 (stretch)
Release:	9.4
Codename:	stretch
$
$ cat /proc/cpuinfo 
processor	: 0
model name	: ARMv7 Processor rev 4 (v7l)
BogoMIPS	: 76.80
Features	: half thumb fastmult vfp edsp neon vfpv3 tls vfpv4 idiva idivt vfpd32 lpae evtstrm crc32 
CPU implementer	: 0x41
CPU architecture: 7
CPU variant	: 0x0
CPU part	: 0xd03
CPU revision	: 4
processor	: 1
model name	: ARMv7 Processor rev 4 (v7l)
BogoMIPS	: 76.80
Features	: half thumb fastmult vfp edsp neon vfpv3 tls vfpv4 idiva idivt vfpd32 lpae evtstrm crc32 
CPU implementer	: 0x41
CPU architecture: 7
CPU variant	: 0x0
CPU part	: 0xd03
CPU revision	: 4
processor	: 2
model name	: ARMv7 Processor rev 4 (v7l)
BogoMIPS	: 76.80
Features	: half thumb fastmult vfp edsp neon vfpv3 tls vfpv4 idiva idivt vfpd32 lpae evtstrm crc32 
CPU implementer	: 0x41
CPU architecture: 7
CPU variant	: 0x0
CPU part	: 0xd03
CPU revision	: 4
processor	: 3
model name	: ARMv7 Processor rev 4 (v7l)
BogoMIPS	: 76.80
Features	: half thumb fastmult vfp edsp neon vfpv3 tls vfpv4 idiva idivt vfpd32 lpae evtstrm crc32 
CPU implementer	: 0x41
CPU architecture: 7
CPU variant	: 0x0
CPU part	: 0xd03
CPU revision	: 4
Hardware	: BCM2835
Revision	: a22082
Serial		: 000000003eabfaf8
```

## Install golang on rpi

``` bash
$ cd ~/Downloads
$ wget https://dl.google.com/go/go1.10.2.linux-armv6l.tar.gz
$ sudo tar -C /usr/local -xvzf go1.10.2.linux-armv6l.tar.gz
$ cd ~
$ mkdir -p ~/workspace/go/{bin,pkg,src}
$ sudo nano /etc/profile.d/gopath.sh
export PATH=$PATH:/usr/local/go/bin
^C

$ sudo nano ~/.bash_profile
export GOBIN="$HOME/workspace/go/bin"
export GOPATH="$HOME/workspace/go/src"
export GOROOT="/usr/local/go"
^C

$ source /etc/profile && source ~/.bash_profile
$ go version
$ go env
$ mkdir -p $HOME/workspace/go/src/helloworld
$ nano $HOME/workspace/go/src/helloworld/helloworld.go
package main
import "fmt"
func main() {
	fmt.Println("Hello world")
}
^C

$ go install $GOPATH/src/helloworld/helloworld.go
$ $GOBIN/helloworld

```

## Install pcsc-lite

``` bash
$ cd ~/Downloads
$ sudo apt install libsystemd-dev
$ sudo apt install libudev-dev
$ wget https://pcsclite.apdu.fr/files/pcsc-lite-1.8.23.tar.bz2
$ tar -C /tmp/ -xvjf pcsc-lite-1.8.23.tar.bz2
$ cd /tmp/pcsc-lite-1.8.23
$ ./configure
$ make
$ sudo make install
$ pcscd -v
pcsc-lite version 1.8.23.
...
```

## Install ccid

``` bash
$ cd ~/Downloads
$ sudo apt install libusb-1.0-0-dev
$ wget https://ccid.apdu.fr/files/ccid-1.4.28.tar.bz2
$ tar -C /tmp/ -xvjf ccid-1.4.28.tar.bz2
$ cd /tmp/ccid-1.4.28/
$ ./configure 
$ make
$ sudo make install
```

## Run pcscd service

Testminal 01

``` bash
$ sudo killall -9 pcscd
$ sudo LIBCCID_ifdLogLevel=0x000F pcscd --foreground --debug --apdu --color | tee log.txt
$ sudo pcscd --foreground --debug --apdu --color | tee log1.txt
```

Plug smardcard device and run go application in testminal 02

``` bash
$ go get -u github.com/Napat/sscard/sscard
$ go build -o sscard github.com/Napat/sscard/main
$ ./sscard
```

## Reference

- [pcsc](https://pcsclite.apdu.fr/files/)
- [ccid](https://ccid.apdu.fr/)
