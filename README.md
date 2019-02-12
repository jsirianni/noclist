# noclist
Simple repository to practice GO and the Cobra CLI

## developing
Noclist is developed with GO 1.10.4 on 64 bit Linux, however, any recent version should work.

### tests
Use GO's built in test suite:
```
go test ./..
```

### building
```
go build
```

## usage

Start the Noc server with docker:
```
docker run --rm -p 8888:8888 adhocteam/noclist
```

Get a list of users
```
noclist get users
```
advanced usage
```
noclist get users --host localhost --port 8888
```

The output will look similar to this:
```
11199617667907522182
17679958362462205107
8161384308545570741
17905933069300437710
...
```
