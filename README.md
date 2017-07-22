# Serve
Serve static files

## Install
```shell
go get -u github.com/itsjimi/serve
```

## Usage
```shell
serve -port="8080" -path="."
```

## Config file
You can create a `.serve.json` in a directory to load a configuration instead of flags.
```json
{
  "port": 1234,
  "path": "./"
}
```