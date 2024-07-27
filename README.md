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

## Como usar e configurar o projeto

### Variáveis de Ambiente

Crie um arquivo `.env` com as configurações necessárias. Você pode usar o arquivo de exemplo `.env.example` como referência.

```bash
cp .env.example .env
```

### Executando Localmente:

#### 1. Rodar o projeto

Para rodar o backend do projeto, execute o comando abaixo a partir do diretório raiz do projeto:

```bash
go run cmd/server/main.go
```

### Executando com Docker Compose:

#### 1. Construir e executar os containers

Para rodar o projeto com o Docker Compose, execute o comando abaixo a partir do diretório raiz do projeto:

```bash
docker compose -f cmd/server/docker-compose.yml up
```

#### 2. Visualizar logs

Para visualizar os logs, execute o comando abaixo a partir do diretório raiz do projeto:

```bash
docker compose -f cmd/server/docker-compose.yml logs -f
```

#### 3. Parar os containers

Para parar os containers, execute o comando abaixo a partir do diretório raiz do projeto:

```bash
docker compose -f cmd/server/docker-compose.yml down
```

#### 4. Monitorar o desempenho dos containers

Para monitorar o desempenho dos containers, execute o comando abaixo a partir do diretório raiz do projeto:

```bash
docker compose -f cmd/server/docker-compose.yml stats
```
