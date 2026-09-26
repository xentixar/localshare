# Localshare - temporary file sharing over your LAN

A cli tool that hosts files over the LAN and so that other users can access them.

### How it works
- This commands starts a http server with default port 8092 and creates a `GET /` route with all the files linked which are passed to the `--files` flag.

```bash
localshare share --files photo.png
```

### Available Commands
1. help
2. share

### Available flags
share
  - --files <paths...>      Files to share.Accepts multiple files.
  - --password <password>   Protect the shared files with a password.
  - --port <port>           Port to host the sharing server on. Default: 8092

### Installation
1. Clone the repo
```bash
git clone https://github.com/xentixar/localshare
```

2. Navigate to the directory
```bash
cd localshare
```

3. Build
```bash
CGO_ENABLED=0 go build -o localshare
```

4. Install
```bash
# go env -w GOBIN=$HOME/.local/bin
go install
```
5. Use
```bash
localshare help
```
