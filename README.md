# gopack

[![Release](https://img.shields.io/github/release/cloudsmith-io/gopack.svg?style=for-the-badge)](https://github.com/cloudsmith-io/gopack/releases/latest)

`gopack` is a small utility designed to pack a local Go module into a Zip file which can then be distributed using a private Go module registry (like Cloudsmith).

## Install

Download the appropriate version for your platform from [`gopack` Releases](https://github.com/cloudsmith-io/gopack/releases). Once downloaded, the binary can be run from anywhere. You don’t need to install it into a global location. This works well for shared hosts and other systems where you don’t have a privileged account.

Ideally, you should install it somewhere in your `PATH` for easy use. `/usr/local/bin` is the most probable location.

## Usage

`gopack` has a single required argument, the version number. Assuming you have a module checked out in the current directory, you can run `gopack` as follows:

```bash
$ gopack v1.0.1
```

If all goes well you should then see the file `v1.0.1.zip` in the current directory. If your module lives in a different directory, you can pass that as the second argument:

```bash
$ gopack v1.2.3 ../module/lives/here/
```
