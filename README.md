**English** | [Português (Brasil)](./README.pt-BR.md)

<p align="center">
  <a href="https://wppconnect.io/">
    <picture>
      <source media="(prefers-color-scheme: dark)" srcset="https://wppconnect-team.github.io/brand/wppconnect-dark.png" />
      <img src="https://wppconnect-team.github.io/brand/wppconnect-light.png" alt="WPPConnect" width="220" height="121" />
    </picture>
  </a>
</p>

# whatsmeow

Use the WhatsApp multi-device protocol from Go applications.

A WPPConnect Team fork in the WPPConnect ecosystem. Original project credits and license are preserved.

[Getting started](#getting-started) | [Documentation](#documentation) | [Contributing](#contributing) | [Support WPPConnect](#support-the-project)

[![License: MPL-2.0](https://img.shields.io/badge/license-MPL--2.0-blue)](./LICENSE) [![Documentation](https://img.shields.io/badge/docs-WPPConnect-blue)](https://wppconnect.io/) [![Open Collective](https://img.shields.io/badge/Open%20Collective-WPPConnect-blue)](https://opencollective.com/wppconnect)

## Overview

For Go developers building socket-based integrations. This WPPConnect Team fork of [tulir/whatsmeow](https://github.com/tulir/whatsmeow) supplies the Go Server ecosystem. Its module path remains `go.mau.fi/whatsmeow`; a plain upstream `go get` does not select this fork.

## Features

- Text/media messages, group management, invite links and delivery/read receipts.
- Presence, app-state synchronization and retry handling.
- Status sending is experimental. The original guide does not claim broadcast-list sending or calls.

## Getting started

Use the toolchain specified by [go.mod](go.mod): Go **1.26.0** with toolchain **go1.27.2**. To inspect/build this fork:

```sh
git clone https://github.com/wppconnect-team/whatsmeow.git
cd whatsmeow
go mod download
go build ./...
```

For integration into another Go module, select the fork explicitly:

```sh
go mod edit -replace=go.mau.fi/whatsmeow=github.com/wppconnect-team/whatsmeow@c9a93790f963bf469f1d322138480590c4cc8f0f
go get go.mau.fi/whatsmeow
```

Follow the [package example](https://pkg.go.dev/go.mau.fi/whatsmeow#example-package) for a client, persistent store and QR pairing. The linked API documentation is upstream; compare it with this pinned fork source.

## Configuration

Choose a durable store and retain session credentials privately. Pin the replacement revision in your application's `go.mod`; dependency behavior follows that revision. The Go Server repository provides an example of a versioned fork replacement and dependency notices.

## Documentation

- [Upstream Go reference](https://pkg.go.dev/go.mau.fi/whatsmeow)
- [Source and module](./go.mod)
- [Retained technical details and acknowledgments](documentation/readme-reference.md)

## WPPConnect ecosystem

| Project | Purpose | When to use it |
| --- | --- | --- |
| [WPPConnect Server Go](https://github.com/wppconnect-team/wppconnect-server-go) | Run a browser-free WhatsApp HTTP API built with Go and whatsmeow. | Evaluate a Go HTTP API and documented migration coverage. |

## Support the project

The WPPConnect contribution below supports maintenance of this fork within the ecosystem; it does not imply funding or endorsement of upstream maintainers.

WPPConnect is maintained by an open-source community. Your support helps fund maintenance, infrastructure, documentation, and the people working on the ecosystem.

If WPPConnect helps your project or business, consider making a one-time contribution or becoming a recurring supporter.

**[Support WPPConnect on Open Collective](https://opencollective.com/wppconnect)**

Contributions are voluntary. See the official page for current financial information. You can also help with code, tests, documentation, translations and well-described issue reports.

## Contributing

Contributions of code, tests, examples, documentation and translations are welcome. For larger changes, discuss the approach in an issue first. Include the project/runtime versions, a minimal reproduction and sanitized logs when reporting a problem. Update both `README.md` and `README.pt-BR.md` in the same change when installation, configuration, features, examples, licensing or funding changes.

Read [CONTRIBUTING.md](CONTRIBUTING.md).

After installing the development dependencies, use the repository checks relevant to your change:

```sh
go build ./...
go test ./...
```

Account-dependent examples and tests require separate validation with an authorized test account. Never attach credentials or session files to an issue.

## Community and support

Use [repository issues](https://github.com/wppconnect-team/whatsmeow/issues) for reproducible bugs and feature proposals. For usage questions, join the community through the official channels:

- [Discord](https://discord.gg/JU5JGGKGNG)
- [Telegram](https://t.me/wppconnect)
- [YouTube](https://www.youtube.com/c/wppconnect)

Upstream discussion: [#whatsmeow:maunium.net](https://matrix.to/#/#whatsmeow:maunium.net) and [WhatsApp protocol Q&A](https://github.com/tulir/whatsmeow/discussions/categories/whatsapp-protocol-q-a). Use fork issues for WPPConnect-specific differences.

This project is independent and is not affiliated with or endorsed by WhatsApp or Meta. Follow applicable terms and obtain authorization for account and messaging operations.

## License

[Mozilla Public License 2.0 (MPL-2.0)](./LICENSE).

Original whatsmeow authors and contributors; retain copyright headers in the source files.

The original legal document and existing copyright notices are preserved. Third-party components retain their own terms.
