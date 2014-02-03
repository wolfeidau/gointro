# go tutorial

To get started with this tutorial open up your terminal.

Navigate to the location you copied the contents of the gostick, in my
case `/Users/markw/gostick`.

```bash
cd gostick
```

## go not installed

If you don't have go installed pick a package for your os and extract it here. The options are:

* go1.2.darwin-386-osx10.6.tar.gz
* go1.2.darwin-amd64-osx10.8.tar.gz
* go1.2.linux-386.tar.gz
* go1.2.linux-amd64.tar.gz

```bash
tar cxvf packages/go1.2.linux-amd64.tar.gz
```

Now `source` the env script.

```bash
source env.sh
```

## go installed

If you already have go installed just `source` the gopath script.

```bash
source gopath.sh
```

Check the go command is in the path in my case this resolves to
`/Users/markw/gostick/go/bin/go`.

```bash
which go
```

Navigate into the hello world example and build it.

```bash
cd helloworld
go build -v
./helloworld
```

This should output.

```
yo
```

# Links

This is based on the talk that [Mark Smith](https://twitter.com/zorkian) did at lca2014.

* https://github.com/xb95/lca2014
* http://www.slideshare.net/dreamwidth/lca2014-introduction-to-go
