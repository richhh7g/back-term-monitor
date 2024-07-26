# Term Alarms | Backend

<p align="center">
  <a href="https://github.com/richhh7g/term-alarms" target="_blank">
    <img src="https://term-monitor.s3.amazonaws.com/monitor.png" alt="Term Alarms" height="150" width=325"/>
  </a>
</p>

## Descrição

Este projeto e o backend da API do Term Alarms para monitorar concorrentes que usam termos de marca em resultados patrocinados do Google.

## Requisitos

| Ferramenta | Versão | Descrição
| - | - | -
| `Golang` | `>= v1.19.0` | Versão [superior](https://golang.org/dl/) ou [igual](https://golang.org/dl/#go1.19) a `1.19.0`.
| `Docker CLI` | `>= v20.10.0` | Versão [superior](https://github.com/docker/cli/tags) ou [igual](https://github.com/docker/cli/releases/tag/v20.10.0).
| `Docker Compose` | `>= v2.0.0` | Versão [superior](https://github.com/docker/compose/releases) ou [igual](https://github.com/docker/compose/releases/tag/v2.0.0) a `2.0.0`.

## Como usar

### 1. Configurar as Variáveis de Ambiente

Crie um arquivo `.env` com as configurações necessárias. Você pode usar o arquivo de exemplo `.env.example` como referência.

```bash
cp .env.example .env
```

### 2. Rodar o Projeto

Para rodar o backend do projeto, execute o comando abaixo a partir do diretório raiz do projeto:

```bash
go run cmd/server/main.go
```
