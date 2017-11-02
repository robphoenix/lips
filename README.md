
```
   _   _   _   _
  / \ / \ / \ / \
 ( L | i | p | s )
  \_/ \_/ \_/ \_/

     List IPs
```

Lips is basically a Go version of [prips](https://gitlab.com/prips/prips), a
tool that can be used to print all of the IP address on a given range.

It currently only works with a slash notation CIDR address (`192.168.0.0/24`).

## Examples

```sh
$ lips 192.168.0.0/24
192.168.0.0
192.168.0.1
...
192.168.0.255
```


## Binaries

Releases are available on the
[tags](https://git-mw.uk365office.co.uk/robphoenix/lips/tags) page.
