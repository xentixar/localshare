# Localshare - temporary file sharing over your LAN

A cli tool that hosts a file over that LAN and user can get the file.

### How it works
- User starts a daemon using the following command

```bash
localshare share --files photo.png
```

- It starts a http server in some port and then shows the url and a qr from where other devices in the LAN access it.

### Commands
1. help
2. list
