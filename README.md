# ContactQR

This project creates QR Codes for your contact information that can be read with a modern smartphone. Most phones expect information in a vCard format. For example, iPhones will scan a QR code automatically when the camera is open.

The project was inspired by a friend who had one of these QR codes on thier lock screen. This enabled people to simply scan the code instead of sending a text message to exchange contact info.

Try it out via Docker:

```sh
docker run aphexddb/contactqr
```

And then visit [http://localhost:8080](http://localhost:8080).

## Releasing

To build a release update the VERSION file (if needed) and run:

```sh
make docker_build
```

To push a release

```sh
make docker_push
```

## Development

To run the server for development:

```bash
make dev
```

Or run the container:

```sh
make docker_build
make run
```

To run the UI for development:

```bash
cd ui
make dev
```

### Requirements

* UI built with [Gatsby](https://www.gatsbyjs.org/docs/) using npm `5.6.0` and node `v8.11.2`.
* Server built using Go `1.11`.

### Notes

The `vcard` package treats VCard as immutable by using the functional option pattern to build the vCard and getters to read data. Please feel free to use this package for your project! The tests are also in a seperate `tests` package to validate that private variables are treated appropriately.

If you find this useful, please report any bugs or add a feature!

[MIT License](./LICENSE)