# srcds-watchdog

This tool queries a given srcds gameserver and returns:

* status 0 if the server is online
* status 1 if the server is offline
* status 2 if the program malfunctions
* status 3 if the parameters passed are wrong

Additionally for servers that are online, a summary is printed.

Example:

```
INFO[0000] 151.80.111.130:27015 is online               
INFO[0000] Hostname: Covid-19 Mod (V1) by ElitE HunterZ |Voice, Rank, 2020 event| 
INFO[0000] Players: 53/64                                
INFO[0000] Game: ████► Zombie Hunting                   
INFO[0000] Version: 5394425   
```


#building
