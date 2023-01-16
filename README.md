# Volume tools

## Encoding volumes

The `enc` directory includes simple script to encode volumes.
The `make.sh` script uses `mkfs.ext4` to create volumes of various sizes, and then encodes them into a small efficient format for recreating.

A variety of sizes has been created already and is available for use in `fs_data.go`.

## Example usage

    package main

    func main() {
	      CreateFile(FS_100g, "100g.volume")
    }

The above will create a new file `100g.volume` using the reference data provided in `FS_100g`. It can typically do this in around 5ms.

It should be noted that created volumes have the same UUID etc. They are byte for byte identical to the original volume.

## TODO

It would be nice to have some helper utils for instance to change the UUID of the disk image.