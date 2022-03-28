# mvd2player
Helper program to play Quake 2 multi-view demos when double-clicked. This ONLY works with Q2Pro

## Config
Before you can use this program, add the config file `mvd2player.json` to your home directory.
- Windows: `c:\users\youraccount`
- Linux: `/home/youraccount`
- Mac: `/Users/youraccount`

Example config:
```
{
    "baseq2folder": "c:/q2/baseq2",
    "q2binary": "c:/q2/q2pro.exe
}
```

Simply edit the paths to match your `baseq2` folder and your q2pro binary

## Usage
In windows, set mvd2player.exe as default application for .mvd2 files.

For all OS, if binary is in your path, you can launch by running `mvd2player <path to demo file>`


## Compiling
- Native: `go build .`
- Cross: `GOOS=windows GOARCH=386 CGO_ENABLED=0 go build .`
