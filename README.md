Hashify [![CircleCI](https://circleci.com/gh/ahmedkamals/hashify.svg?style=svg)](https://circleci.com/gh/ahmedkamals/hashify "Build Status")
======

[![license](https://img.shields.io/github/license/mashape/apistatus.svg)](LICENSE.md "License")
[![release](https://img.shields.io/github/v/release/ahmedkamals/hashify.svg)](https://github.com/ahmedkamals/hashify/releases/latest "Release")
[![codecov](https://codecov.io/gh/ahmedkamals/hashify/branch/main/graph/badge.svg?token=XPINFB5JYV)](https://codecov.io/gh/ahmedkamals/hashify "Code Coverage")
[![GolangCI](https://golangci.com/badges/github.com/ahmedkamals/hashify.svg?style=flat-square)](https://golangci.com/r/github.com/ahmedkamals/hashify "Code Coverage")
[![Go Report Card](https://goreportcard.com/badge/github.com/ahmedkamals/hashify)](https://goreportcard.com/report/github.com/ahmedkamals/hashify "Go Report Card")
[![Codacy Badge](https://app.codacy.com/project/badge/Grade/65feb277726f4a10895f028d460f9196)](https://www.codacy.com/manual/ahmedkamals/hashify?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=ahmedkamals/hashify&amp;utm_campaign=Badge_Grade "Code Quality")
[![GoDoc](https://godoc.org/github.com/ahmedkamals/hashify?status.svg)](https://godoc.org/github.com/ahmedkamals/hashify "Documentation")
[![DepShield Badge](https://depshield.sonatype.org/badges/ahmedkamals/hashify/depshield.svg)](https://depshield.github.io "DepShield")
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fahmedkamals%2Fhashify.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fahmedkamals%2Fhashify?ref=badge_shield "Dependencies")

```bash
 _   _           _     _  __
| | | |         | |   (_)/ _|
| |_| | __ _ ___| |__  _| |_ _   _
|  _  |/ _` / __| '_ \| |  _| | | |
| | | | (_| \__ \ | | | | | | |_| |
\_| |_/\__,_|___/_| |_|_|_|  \__, |
                              __/ |
                             |___/
```

hashes url response using md5.

Table of Contents
-----------------

*   [🏎️ Getting Started](#-getting-started)

    *   [Prerequisites](#prerequisites)
    *   [Installation](#installation)
    *   [Examples](#examples)

*   [🕸️ Tests](#-tests)

    *   [📈 Benchmarks](#-benchmarks)
    *   [⚓ Git Hooks](#-git-hooks)

*   [👨‍💻 Credits](#-credits)

*   [🆓 License](#-license)

🏎️ Getting Started
------------------

### Prerequisites

*   [Golang 1.15 or later][1].

### Installation

```bash
go get -u github.com/ahmedkamals/hashify
cp .env.sample .env
```

### Examples

```bash
make run args="-concurrency 3 http://google.com http://twitter.com http://yahoo.com invalid-url"
```

🕸️ Tests
--------

```bash
make unit
```

### 📈 Benchmarks

```bash
make bench
```

### ⚓ Git Hooks

In order to set up tests running on each commit do the following steps:

```bash
git config --local core.hooksPath .githooks
```

![Benchmarks](https://github.com/ahmedkamals/hashify/raw/main/assets/img/bench.png "Benchmarks")

👨‍💻 Credits
----------

*   [ahmedkamals][2]

🆓 LICENSE
----------

Hashify is released under MIT license, please refer to the [`LICENSE.md`](https://github.com/ahmedkamals/hashify/blob/main/LICENSE.md "License") file.

[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fahmedkamals%2Fhashify.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fahmedkamals%2Fhashify?ref=badge_large "Dependencies")

Happy Coding 🙂

[![Analytics](http://www.google-analytics.com/__utm.gif?utmwv=4&utmn=869876874&utmac=UA-136526477-1&utmcs=ISO-8859-1&utmhn=github.com&utmdt=hashify&utmcn=1&utmr=0&utmp=/ahmedkamals/hashify?utm_source=www.github.com&utm_campaign=hashify&utm_term=hashify&utm_content=hashify&utm_medium=repository&utmac=UA-136526477-1)]()

[1]: https://golang.org/dl/ "Download Golang"
[2]: https://github.com/ahmedkamals "Author"
