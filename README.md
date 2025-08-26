# ZipCheck

![tests workflow](https://onlyhavecans.works/onlyhavecans/zipcheck/badges/workflows/ci.yml/badge.svg)

## Overview

This is a simple command line application for deep testing your zip files.

While you could do something like this with `zip -T` and some bash I wanted to have something threaded to save time on massive nested directories like with ROM files.

## Usage

Download the version for you platform from [releases](https://github.com/onlyhavecans/zipcheck/releases)

Run this by passing in all the directories you want to check

```shell
./zipcheck <directory> <directory> ...
```
