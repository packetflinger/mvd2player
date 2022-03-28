# q2demoplayer
Helper program to play Quake 2 demos and multi-view demos when double-clicked. Playing MVD only works with Q2Pro however.

## Config
Before you can use this program, add the config file `q2demoplayer.json` to your home directory.
- Windows: `c:\users\youraccount`
- Linux: `/home/youraccount`
- Mac: `/Users/youraccount`

Example config:
```
{
    "baseq2folder": "c:/q2",
    "q2binary": "q2pro.exe"
}
```

Simply edit the paths to match your Quake 2 folder and your q2pro binary name

## Usage
In windows, set q2demoplayer.exe as default application for .mvd2 and .dm2 files.

For all OS, if binary is in your path, you can launch by running `q2demoplayer <path to demo file>`


## Compiling
- Native: `go build .`
- Cross: `GOOS=windows GOARCH=386 CGO_ENABLED=0 go build .`
