# Localshare - temporary file sharing over your LAN

A cli tool that hosts a file over that LAN and user can get the file.

### How it works
- User starts a daemon using the following command

```bash
localshare share ./photo.png
```

- It starts a http server in some port and then shows the url and a qr from where other devices in the LAN access it.

### Commands
1. password
2. help
3. list
4. share
5. view
