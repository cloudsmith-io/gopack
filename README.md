# gopack

`gopack` is a small utility designed to pack a local Go module into a Zip file which can then be distributed using a private Go module registry (like Cloudsmith).

## Usage

`gopack` has a single required argument, the version number. Assuming you have a module checked out in the current directory, you can run `gopack` as follows:

```bash
$ gopack v1.0.1
```

If all goes well you should then see the file `v1.0.1.zip` in the current directory. If your module lives in a different directory, you can pass that as the second argument:

```bash
$ gopack v1.2.3 ../module/lives/here/
```
