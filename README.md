# Gohn

A simple CLI for fetching Hacker News stories.

## Prerequisites

[Download and install Go](https://golang.org/doc/install)

## Quick start

Clone the repo, compile the code, and run it on your computer.

```bash
git clone https://github.com/neil-berg/gohn.git
cd gohn
go build gohn.go
./gohn --count=5 --type="top"
```

## Flag options

**--count (integer)**

The number of stories to fetch and display. Valid values between 1 and 10.

**--type (string)**

The type of stories to query. Valid values are:

- `"top"`
- `"ask"`
- `"show"`
- `"job"`
