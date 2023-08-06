# GoLang Brute Forcer

[![Go Report Card](https://goreportcard.com/badge/github.com/tomp332/golang-brute-forcer)](https://goreportcard.com/report/github.com/tomp332/brute-forcer)
[![Build](https://github.com/your-username/golang-brute-forcer/actions/workflows/build.yml/badge.svg)](https://github.com/tomp332/brute-forcer/actions/workflows/build.yml)
[![License](https://img.shields.io/github/license/your-username/golang-brute-forcer.svg)](https://github.com/tomp332/brute-forcer/LICENSE.md)
[![Code Coverage](https://codecov.io/gh/your-username/golang-brute-forcer/branch/main/graph/badge.svg)](https://codecov.io/gh/tomp332/brute-forcer)

![GoLang Brute Forcer](banner.png)

## Description

GoLang Brute Forcer is a command-line tool that performs brute force attacks to test the security of passwords or keys.
It is designed to help identify weak passwords and improve overall system security.

## Features

- Customizable character sets and password length.
- Parallel processing for faster brute-forcing.
- Supports different hashing algorithms (MD5, SHA-256, etc.).
- Customizable output format.

## Installation

### Pre-built binaries

You can download pre-built binaries for your platform from
the [Releases](https://github.com/your-username/golang-brute-forcer/releases) page. Extract the archive and add the
binary to your PATH.

### Build from source

To build from source, make sure you have Go (1.15+) installed and run the following commands:

```bash
$ git clone https://github.com/your-username/golang-brute-forcer.git
$ cd golang-brute-forcer
$ go build