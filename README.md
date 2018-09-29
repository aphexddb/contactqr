# ContactQR

This project creates QR Codes for your contact information that can be read with a modern smartphone. Most phones expect information in a vCard format. For example, iPhones will scan a QR code automatically when the camera is open.

The project was inspired by a friend who had one of these QR codes on thier lock screen. This enabled people to simply scan the code instead of sending a text message to exchange contact info.

Try it out via Docker:

```sh
docker run aphexddb/contactqr
```

And then visit [http://localhost:8080](http://localhost:8080).

## Building

To build a release update the VERSION file (if needed) and run:

```sh
make docker_build
```

To push a release

```sh
make docker_push
```

## Development

For local development you can change the container port as follows:

```sh
docker run --rm -it -p 8080:8080/tcp -e PORT=8080 contactqr:latest
```

Otherwise, clone this repo and run it:

```sh
make dev
```

### Notes

Built using Go 1.11, probably runs on older versions.

The `vcard` package treats VCard as immutable by using the functional option pattern to build the vCard and getters to read data. Please feel free to use this package for your project! The tests are also in a seperate `tests` package to validate that private variables are treated appropriately.

If you find this useful, please report any bugs or add a feature!

[MIT License](./LICENSE)