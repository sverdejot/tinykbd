# tinykbd: single-key macro keyboard

This repo contains all the needed files to create a macro-pad using an SAMD21-based microcrontroller, which, when pressed, opens an URL in the user browser.

This project also automates the install script distribution by deploying the necessary files to an AWS S3 buckets which, in my case, is exposed behind a CloudFront distribution.

> I have also create an HTTP redirection on my DNS provider so the install script is available under https://tinykbd.sverdejot.dev

In order to use the macro pad, the user first needs to install the binary that will check for the key press by continously reading from serial port. This is done by running:

```bash
curl -sfL https://tinykbd.sverdejot.dev | sh -s
```

Then, just by connecting the SAMD21 to the PC it will automatically connect and be ready to be used.

## Considerations

* The installation script is custom for this specific compilation of the binary, which does nothing. You need to create, compile and configure your own fork of the repository in order to personlize it.
* Currently, the binary only compiles for Apple Silicon ARM architecture. In case you want to compile it for any other arch. you will need to change both for Go to build the binary and GH Action to build it on the right runner.
