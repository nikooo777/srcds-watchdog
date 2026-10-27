# srcds-watchdog

## Usage
This tool queries a given srcds gameserver and exits with:

* code 0 if the server is online
* code 1 if the server is offline (timeout or connection refused)
* code 2 if the program malfunctions
* code 3 if the parameters passed are wrong

Additionally for servers that are online, a summary is printed.

## Example

```
$ ./srcds-watchdog 54.37.245.51:27045

INFO[0000] 54.37.245.51:27045 is online
INFO[0000] Hostname: Zombie Hunting V3.1 by ElitE HunterZ |All Weapons|Rank|Respawn
INFO[0000] Players: 26/64
INFO[0000] Game: ████► Zombie All Weapons
INFO[0000] Version: 11003710
```

## Installing
You will need Go 1.27+ installed. Refer to the [official documentation](https://go.dev/doc/install) for how to install it.

```bash
go install github.com/nikooo777/srcds-watchdog@latest
```

## Building

```bash
git clone https://github.com/nikooo777/srcds-watchdog.git
cd srcds-watchdog
go build
```

## Download
Prebuilt binaries for Linux (amd64, arm64) and Windows (amd64) are attached to each [release](https://github.com/nikooo777/srcds-watchdog/releases).

No guarantees that it will work on your system. Also it's always best to compile it yourself!
