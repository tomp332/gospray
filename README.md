<div align="center"> 
  <img alt="img" src="https://github.com/tomp332/gobrute/assets/47506972/ea8ea660-867b-46fa-9439-0d7eb7b6c8b3" />

[![Go Report Card](https://goreportcard.com/badge/github.com/tomp332/gobrute?branch=main)](https://goreportcard.com/report/github.com/tomp332/gobrute)
[![Build](https://github.com/tomp332/gobrute/actions/workflows/main.yml/badge.svg)](https://github.com/tomp332/gobrute/actions/workflows/main.yml)
[![License](https://img.shields.io/github/license/tomp332/gobrute.svg)](https://github.com/tomp332/gobrute/blob/main/LICENSE.md)
[![Code Coverage](https://codecov.io/gh/tomp332/gobrute/settings/badge.svg)](https://codecov.io/gh/tomp332/gobrute?branch=main)
</div>

## Description

`goBrut` is a powerful password brute forcer written in Go. It is designed to help security professionals and
penetration testers test the strength of passwords and identify weak passwords in various scenarios.

`goBrut` supports customizable character sets, password length, and hashing algorithms, making it versatile for various
brute force attacks.

## Features

- Customizable character sets and password length.
- Parallel processing for faster brute-forcing.
- Supports different hashing algorithms (MD5, SHA-256, etc.).
- Customizable output format.
- Configurable retry and timeout options for robustness.
- Easy-to-use command-line interface.

## Installation

### Pre-built Binaries

Download pre-built binaries for your platform from the [Releases](https://github.com/tomp332/gobrut/releases) page.
Extract the archive and add the binary to your PATH.

### Build from Source

Make sure you have Go (1.15+) installed and run the following commands:

```bash
$ git clone https://github.com/tomp332/gobrut.git
$ cd goBrut
$ go build
