# Jack
🍽 Serve static files by one simple command

## Install
```shell
go get -u github.com/itsjimi/jack
```

## Usage
```shell
jack serve -addr="0.0.0.0" -port="8080" -path="."
```

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
[MIT](https://github.com/ItsJimi/serve/blob/master/LICENSE)
