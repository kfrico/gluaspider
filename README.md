## go-nsq

[![GoDoc](https://godoc.org/github.com/kfrico/gluaspider?status.svg)](https://godoc.org/github.com/kfrico/gluaspider) [![GitHub release](https://img.shields.io/github/release/kfrico/gluaspider.svg)](https://github.com/kfrico/gluaspider/releases/latest)

simple spider module for [gopher-lua](https://github.com/yuin/gopher-lua)

### Docs

See [godoc](https://godoc.org/github.com/kfrico/gluaspider)

### Installation

```
go get github.com/kfrico/gluaspider
```


### Usage

```
import (
    "github.com/kfrico/gluaspider"
)

L := lua.NewState()
defer L.Close()

L.PreloadModule("spider", gluaspider.NewSpider().Loader)
```

### Tests