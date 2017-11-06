
```
   _   _   _   _
  / \ / \ / \ / \
 ( L | i | p | s )
  \_/ \_/ \_/ \_/

     List IPs
```

Lips is basically a Go version of [prips](https://gitlab.com/prips/prips), a
tool that can be used to print all of the IP address on a given range.

## Examples

**CIDR address**

```sh
$ lips 192.168.0.0/24
192.168.0.0
192.168.0.1
...
192.168.0.255
```

**IP range**

```sh
$ lips 192.168.0.0 192.168.0.255
192.168.0.0
192.168.0.1
...
192.168.0.255
```


## Installation

If you have **Go** installed you can clone this repo and run `go install` from
within it. Otherwise releases are available on the
[tags](https://git-mw.uk365office.co.uk/robphoenix/lips/tags) page. If you are
downloading a binary be sure to either place it in your `$PATH` or add the
directory you download it to to your `$PATH`. This will mean you should be able
to use the `lips` command from anywhere on your machine.
