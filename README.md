# binner

binner is a simple Go CLI app that I wrote to deploy my own scripts or CLI apps to `/usr/local/bin`. Toss your tools in the bin.

## Background
I wrote a shell script to sort some files when I was organizing some SNES ROMS before copying them to my homebrewed Wii, and I copied the script to `/usr/local/bin` and assigned execute permissions so that it was in my PATH and I could call it in my shell from any directory. I decided to create a simple tool for my self to do this.

## Installation and Usage
```bash
go install github.com/jetkub/binner
```

Be cheeky and use **binner** to deploy **binner** to `/usr/local/bin` 😉. Call the program with one argument: the path to the binary you want to copy. If the binary already exists, it will ask if you want to overwrite it or not.
```
sudo ~/go/bin/binner ~/go/bin/binner
```
