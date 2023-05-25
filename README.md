# mvx-create_wallet

## Description

This tool allows you to retrieve a wallet containing the specified name.

## Pre-requisites

- [Go](https://golang.org/doc/install) 1.18 or higher

## Installation

```bash
go get github.com/multivers3x/mvx-create_wallet
```

## Commands

```go
 -log string
        Set the log level: debug, info, warn, error, fatal, panic (default "info")
 -wallet string
        Allows you to retrieve a wallet containing the specified name (default "multiversx")
```

### Examples

#### Build

```bash
go build .
```

#### Get wallet by name (default: multiversx)

```bash
./mvx-create_wallet -wallet multiversx
```

#### Set log level

```bash
./mvx-create_wallet -log debug
```
