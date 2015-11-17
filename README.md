Spriteful
=========

Spriteful is an API that provides server boot configuration for pixiecore.

See [pixiecore](https://github.com/danderson/pixiecore).

## Building

To build spriteful, go v1.5+ is needed and the `GO15VENDOREXPERIMENT` should be enabled.

```
$ export GO15VENDOREXPERIMENT=1
```

then, clone the repo:

```
$ git clone https://github.com/AcalephStorage/spriteful.git
$ cd spriteful
```

then, load the submodules (dependencies):

```
$ git submodule init
$ git submodule update
```

and finally, build the binary:

```
$ go build .
```

## Running

To run spriteful, use the following command:

```
$ spriteful -config /path/to/config/file
```

a sample config file is provided [here](config.json.example).

## pixiecore integration

To integrate with `pixiecore`, point the `-api` argument to this api:

```
$ pixiecore -api http://{bindHost}:{bindPort}/api
```
