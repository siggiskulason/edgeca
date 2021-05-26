# EdgeCA
**EdgeCA** is an ephemeral, in-memory CA providing service mesh machine identities.
It automates the management and issuance of TLS certificates. It can either run with a self-certificated Root CA certificate or use an issuing certificate retrieved using the [Venafi vCert](https://github.com/Venafi/vcert) software.

It solves the many limitations of the embedded service mesh CAs by providing developers a fast, easy, and integrated source of machine identities whilst also providing security teams with the required policy and oversight.  

It also enables ephemeral certificate-based authorization, which reduces the need for permanent access credentials, explicit access revocation or traditional SSH key management. 

EdgeCA is open source, written in Go, and licenced with the Apache 2.0 Licence

For more information read these [instructions](docs) on how to install and run EdgeCA. 

The easiest way to install the application is to use [snaps](./snap)

```
snap install edgeca
```

[![Get it from the Snap Store](https://snapcraft.io/static/images/badges/en/snap-store-white.svg)](https://snapcraft.io/edgeca)

## Contributing to EdgeCA
**EdgeCA** is an open source project currently in early development stages. We welcome and appreciate all contributions from the developer community.
Please read our documentation on [contributing](https://github.com/edgesec-org/edgeca/blob/main/CONTRIBUTING.md) for more information. To report a problem or share an idea, create an [Issue](https://github.com/edgesec-org/edgeca/issues) and then use [Pull Requests](https://github.com/edgesec-org/edgeca/pulls) to contribute bug fixes or proposed enhancements. Got questions? [Join us on Slack](https://join.slack.com/t/edgesec/signup)!

## License
Copyright 2020-2021 © [EdgeSec OÜ](https://edgesec.org). All rights reserved.

EdgeCA is licensed under the Apache License, Version 2.0. See [LICENSE](https://github.com/edgesec-org/edgeca/blob/main/LICENSE) for the full license text.
