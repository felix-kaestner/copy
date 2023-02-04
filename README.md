# Copy

<p align="center">
    <span>Create deep copies of values in Go.</span>
    <br><br>
    <a href="https://github.com/felix-kaestner/copy/issues">
        <img alt="Issues" src="https://img.shields.io/github/issues/felix-kaestner/copy?color=29b6f6&style=flat-square">
    </a>
    <a href="https://github.com/felix-kaestner/copy/stargazers">
        <img alt="Stars" src="https://img.shields.io/github/stars/felix-kaestner/copy?color=29b6f6&style=flat-square">
    </a>
    <a href="https://github.com/felix-kaestner/copy/blob/main/LICENSE">
        <img alt="License" src="https://img.shields.io/github/license/felix-kaestner/copy?color=29b6f6&style=flat-square">
    </a>
    <a href="https://pkg.go.dev/github.com/felix-kaestner/copy">
        <img alt="Stars" src="https://img.shields.io/badge/go-documentation-blue?color=29b6f6&style=flat-square">
    </a>
    <a href="https://goreportcard.com/report/github.com/felix-kaestner/copy">
        <img alt="Issues" src="https://goreportcard.com/badge/github.com/felix-kaestner/copy?style=flat-square">
    </a>
    <!-- <a href="https://codecov.io/gh/felix-kaestner/copy">
        <img src="https://img.shields.io/codecov/c/github/felix-kaestner/copy?style=flat-square&token=KK7ZG7A90X"/>
    </a> -->
    <a href="https://twitter.com/kaestner_felix">
        <img alt="Twitter" src="https://img.shields.io/badge/twitter-@kaestner_felix-29b6f6?style=flat-square">
    </a>
</p>

## Usage 

```go
package main

import "github.com/felix-kaestner/copy"

type Answer struct {
    Value int
}

func main() {
    meaningOfLife := &Answer{Value: 42}

    quote := copy.Deep(meaningOfLife)
}
```

##  Installation

Install with the `go get` command:

```
$ go get -u github.com/felix-kaestner/copy
```

## Contribute

All contributions in any form are welcome! 🙌🏻  
Just use the [Issue](.github/ISSUE_TEMPLATE) and [Pull Request](.github/PULL_REQUEST_TEMPLATE) templates and I'll be happy to review your suggestions. 👍

## Cheers ✌🏻
