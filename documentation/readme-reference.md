# Retained README reference / Referência preservada do README

This document preserves the previous `README.md` for technical details, acknowledgments and history. Start with [README.md](../README.md) or [README.pt-BR.md](../README.pt-BR.md) for the current verified entry path. Older commands, release claims, third-party examples and API signatures here may need migration; consult the current source and linked guides.

Este documento preserva o `README.md` anterior com detalhes técnicos, agradecimentos e histórico. Comece pelo [README em português](../README.pt-BR.md) para as instruções iniciais conferidas. Comandos antigos, afirmações sobre releases, exemplos de terceiros e assinaturas de API podem exigir migração; consulte os fontes e guias atuais.

---

<!-- wppconnect-official-brand -->
<p align="center">
  <a href="https://wppconnect-team.github.io/">
    <picture>
      <source media="(prefers-color-scheme: dark)" srcset="https://wppconnect-team.github.io/brand/wppconnect-dark.png" />
      <img src="https://wppconnect-team.github.io/brand/wppconnect-light.png" alt="WPPConnect" width="220" height="121" />
    </picture>
  </a>
</p>
<!-- /wppconnect-official-brand -->

# whatsmeow
[![Go Reference](https://pkg.go.dev/badge/go.mau.fi/whatsmeow.svg)](https://pkg.go.dev/go.mau.fi/whatsmeow)

whatsmeow is a Go library for the WhatsApp web multidevice API.

## Discussion
Matrix room: [#whatsmeow:maunium.net](https://matrix.to/#/#whatsmeow:maunium.net)

For questions about the WhatsApp protocol (like how to send a specific type of
message), you can also use the [WhatsApp protocol Q&A] section on GitHub
discussions.

[WhatsApp protocol Q&A]: https://github.com/tulir/whatsmeow/discussions/categories/whatsapp-protocol-q-a

## Usage
The [godoc](https://pkg.go.dev/go.mau.fi/whatsmeow) includes docs for all methods and event types.
There's also a [simple example](https://pkg.go.dev/go.mau.fi/whatsmeow#example-package) at the top.

## Features
Most core features are already present:

* Sending messages to private chats and groups (both text and media)
* Receiving all messages
* Managing groups and receiving group change events
* Joining via invite messages, using and creating invite links
* Sending and receiving typing notifications
* Sending and receiving delivery and read receipts
* Reading and writing app state (contact list, chat pin/mute status, etc)
* Sending and handling retry receipts if message decryption fails
* Sending status messages (experimental, may not work for large contact lists)

Things that are not yet implemented:

* Sending broadcast list messages (this is not supported on WhatsApp web either)
* Calls
