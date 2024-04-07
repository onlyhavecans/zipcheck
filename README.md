# ZipCheck

![tests workflow](https://github.com/onlyhavecans/zipcheck/actions/workflows/go-test.yml/badge.svg)
![lint workflow](https://github.com/onlyhavecans/zipcheck/actions/workflows/golangci-lint.yml/badge.svg)

## Overview

This is a simple command line application for deep testing your zip files.

While you could do something like this with `zip -T` and some bash I wanted to have something threaded to save time on massive nested directories like with ROM files.

## Usage

Download the version for you platform from [releases](https://github.com/onlyhavecans/zipcheck/releases)

```shell
zipcheck <directory> <directory> ...
```
