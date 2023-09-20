# Voltools

Tools for dealing with disk volumes.

[![License: Apache 2.0](https://img.shields.io/badge/License-Apache%202.0-brightgreen.svg)](https://www.apache.org/licenses/LICENSE-2.0)
[![Discord](https://dcbadge.vercel.app/api/server/JYmFhtdPeu?style=flat)](https://loopholelabs.io/discord)
![Go Version](https://img.shields.io/badge/go%20version-%3E=1.18-61CFDD.svg)
[![Go Reference](https://pkg.go.dev/badge/github.com/loopholelabs/voltools.svg)](https://pkg.go.dev/github.com/loopholelabs/voltools)

## Overview

### Encoding Volumes

The `hack` directory includes Makefile to encode volumes.
It uses `mkfs.ext4` to create volumes of various sizes, and then encodes them into a small efficient format for recreating.

A variety of sizes has been created already and is available for use in `fs_data.go`. Attempting to create a filesize not encoded will result in an error.

### Example Usage

```go
package main

func main() {
      err := CreateFile(100, "100g.volume")
    if err != nil {
      panic(err)
    }
}
```

The above will create a new file `100g.volume` using the reference data provided in `fs_data.go`. It can typically do this in around 5ms.

It should be noted that created volumes have the same UUID etc. They are byte for byte identical to the original volume.

### Enhancements

It would be nice to have some helper utils for instance to change the UUID of the disk image.

## Contributing

Bug reports and pull requests are welcome on GitHub at [https://github.com/loopholelabs/voltools][gitrepo]. For more contribution information check out [the contribution guide](https://github.com/loopholelabs/voltools/blob/master/CONTRIBUTING.md).

## License

The Voltools project is available as open source under the terms of the [Apache License, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0).

## Code of Conduct

Everyone interacting in the Voltools project's codebases, issue trackers, chat rooms and mailing lists is expected to follow the [CNCF Code of Conduct](https://github.com/cncf/foundation/blob/master/code-of-conduct.md).

## Project Managed By:

[![https://loopholelabs.io][loopholelabs]](https://loopholelabs.io)

[gitrepo]: https://github.com/loopholelabs/voltools
[loopholelabs]: https://cdn.loopholelabs.io/loopholelabs/LoopholeLabsLogo.svg
[loophomepage]: https://loopholelabs.io
