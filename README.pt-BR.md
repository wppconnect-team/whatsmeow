[English](./README.md) | **Português (Brasil)**

<p align="center">
  <a href="https://wppconnect.io/">
    <picture>
      <source media="(prefers-color-scheme: dark)" srcset="https://wppconnect-team.github.io/brand/wppconnect-dark.png" />
      <img src="https://wppconnect-team.github.io/brand/wppconnect-light.png" alt="WPPConnect" width="220" height="121" />
    </picture>
  </a>
</p>

# whatsmeow

Use o protocolo multidispositivo do WhatsApp em aplicações Go.

Um fork da WPPConnect Team no ecossistema WPPConnect. Os créditos e a licença do projeto original são preservados.

[Primeiros passos](#primeiros-passos) | [Documentação](#documentação) | [Como contribuir](#como-contribuir) | [Apoie a WPPConnect](#apoie-o-projeto)

[![License: MPL-2.0](https://img.shields.io/badge/license-MPL--2.0-blue)](./LICENSE) [![Documentation](https://img.shields.io/badge/docs-WPPConnect-blue)](https://wppconnect.io/) [![Open Collective](https://img.shields.io/badge/Open%20Collective-WPPConnect-blue)](https://opencollective.com/wppconnect)

## Visão geral

Para desenvolvedores Go que criam integrações por socket. Este fork da WPPConnect Team de [tulir/whatsmeow](https://github.com/tulir/whatsmeow) atende ao ecossistema do Server Go. O caminho de módulo continua sendo `go.mau.fi/whatsmeow`; um `go get` comum do upstream não seleciona este fork.

## Funcionalidades

- Mensagens de texto/mídia, administração de grupos, convites e recibos de entrega/leitura.
- Presença, sincronização de app-state e tratamento de tentativas.
- O envio de status é experimental. O guia original não afirma suporte a envio para listas de transmissão nem chamadas.

## Primeiros passos

Use as ferramentas definidas em [go.mod](go.mod): Go **1.26.0** com toolchain **go1.27.2**. Para inspecionar/compilar este fork:

```sh
git clone https://github.com/wppconnect-team/whatsmeow.git
cd whatsmeow
go mod download
go build ./...
```

Para integrar a outro módulo Go, selecione o fork explicitamente:

```sh
go mod edit -replace=go.mau.fi/whatsmeow=github.com/wppconnect-team/whatsmeow@c9a93790f963bf469f1d322138480590c4cc8f0f
go get go.mau.fi/whatsmeow
```

Siga o [exemplo do pacote](https://pkg.go.dev/go.mau.fi/whatsmeow#example-package) para criar cliente, store persistente e pareamento por QR. A documentação de API vinculada é do upstream; compare com os fontes deste fork fixado.

## Configuração

Escolha um store durável e guarde as credenciais de sessão em local privado. Fixe a revisão do replacement no `go.mod` da aplicação; o comportamento da dependência segue essa revisão. O repositório Server Go oferece um exemplo de replacement versionado e avisos de dependências.

## Documentação

- [Referência Go do upstream](https://pkg.go.dev/go.mau.fi/whatsmeow)
- [Fontes e módulo](./go.mod)
- [Detalhes técnicos e agradecimentos preservados](documentation/readme-reference.md)

## Ecossistema WPPConnect

| Projeto | Finalidade | Quando usar |
| --- | --- | --- |
| [WPPConnect Server Go](https://github.com/wppconnect-team/wppconnect-server-go) | Execute uma API HTTP do WhatsApp sem navegador, feita com Go e whatsmeow. | Avaliar uma API HTTP Go e a cobertura documentada de migração. |

## Apoie o projeto

A contribuição à WPPConnect abaixo apoia a manutenção deste fork no ecossistema; ela não implica financiamento ou endosso dos mantenedores do upstream.

A WPPConnect é mantida por uma comunidade de código aberto. Seu apoio ajuda a financiar a manutenção, a infraestrutura, a documentação e o trabalho das pessoas que desenvolvem o ecossistema.

Se a WPPConnect ajuda seu projeto ou sua empresa, considere fazer uma contribuição pontual ou apoiar de forma recorrente.

**[Apoie a WPPConnect no Open Collective](https://opencollective.com/wppconnect)**

As contribuições são voluntárias. Consulte a página oficial para acompanhar as informações financeiras atuais. Você também pode ajudar com código, testes, documentação, traduções e relatos de problemas bem descritos.

## Como contribuir

Contribuições de código, testes, exemplos, documentação e traduções são bem-vindas. Para mudanças maiores, discuta a proposta em uma issue primeiro. Ao relatar um problema, informe as versões do projeto/runtime, uma reprodução mínima e logs sem dados sensíveis. Atualize `README.md` e `README.pt-BR.md` na mesma alteração quando mudar instalação, configuração, funcionalidades, exemplos, licença ou arrecadação.

Leia [CONTRIBUTING.md](CONTRIBUTING.md).

Após instalar as dependências de desenvolvimento, use as verificações pertinentes à sua alteração:

```sh
go build ./...
go test ./...
```

Exemplos e testes que dependem de conta exigem validação separada com uma conta de teste autorizada. Nunca anexe credenciais ou arquivos de sessão a uma issue.

## Comunidade e suporte

Use as [issues do repositório](https://github.com/wppconnect-team/whatsmeow/issues) para bugs reproduzíveis e propostas de funcionalidades. Para dúvidas de uso, participe pelos canais oficiais:

- [Discord](https://discord.gg/JU5JGGKGNG)
- [Telegram](https://t.me/wppconnect)
- [YouTube](https://www.youtube.com/c/wppconnect)

Discussões do upstream: [#whatsmeow:maunium.net](https://matrix.to/#/#whatsmeow:maunium.net) e [perguntas sobre o protocolo WhatsApp](https://github.com/tulir/whatsmeow/discussions/categories/whatsapp-protocol-q-a). Use as issues do fork para diferenças específicas da WPPConnect.

Este projeto é independente e não é afiliado nem endossado pelo WhatsApp ou pela Meta. Siga os termos aplicáveis e obtenha autorização para operações de conta e mensagens.

## Licença

[Mozilla Public License 2.0 (MPL-2.0)](./LICENSE).

Original whatsmeow authors and contributors; retain copyright headers in the source files.

O documento jurídico original e os avisos de autoria existentes são preservados. Componentes de terceiros mantêm seus próprios termos.
