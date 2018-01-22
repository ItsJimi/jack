# Jack
🏗 Amazing web development tool

## Install
```shell
go get -u github.com/itsjimi/jack
```

## Usage
```shell
$ jack
```
```shell
NAME:
   Jack - 🏗 Amazing web development tool

USAGE:
   jack [global options] command [command options] [arguments...]

VERSION:
   0.1.1

COMMANDS:
     serve, s    serve static files
     connect, c  connect to websocket server
     help, h     Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --config value  path of config file (default: ".jack.json")
   --help, -h      show help
   --version, -v   print the version
```

## Commands

### Serve
Start a web server to serve static files

#### Usage
```shell
jack serve
```

#### Options
- `-addr="0.0.0.0"`
- `-port="8080"`
- `-path="."`

### Connect
[WIP]

## Config file
You can create a `.jack.json` in a directory to load a configuration instead of flags
```json
{
  "addr": "127.0.0.1",
  "port": 1234,
  "path": "./"
}
```

## Contribute
Feel free to fork and make pull requests

## License
[MIT](https://github.com/ItsJimi/jack/blob/master/LICENSE)
