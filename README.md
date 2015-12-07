Spriteful
=========

Spriteful is an API that provides server boot configuration for pixiecore.

See [pixiecore](https://github.com/danderson/pixiecore).

## Building

To build spriteful, [gb](http://getgb.io/) is needed.

``` sh
go get github.com/constabulary/gb/...
```

then, clone the repo:

```shell
git clone https://github.com/AcalephStorage/spriteful.git
cd spriteful
```

then, restore dependencies:

```shell
gb vendor restore
```

and finally, build the binary:

```shell
gb build
```

Binary can be found in `bin/` directory.

## Running

To run spriteful, use the following command:

```shell
spriteful -config /path/to/config/file
```

a sample config file is provided [here](config.json.example).

## pixiecore integration

To integrate with `pixiecore`, point the `-api` argument to this api:

```
$ pixiecore -api http://{bindHost}:{bindPort}/api
```
