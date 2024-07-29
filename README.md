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
| `MongoDB` | `>= v4.4.0` | Versão [superior](https://www.mongodb.com/docs/manual/release-notes/4.4/) ou [igual](https://www.mongodb.com/docs/manual/release-notes/4.4/#mongodb-4-4-0) a `4.4.0`.
| `GolangMigrate` | `>= v4.15.0` | Versão [superior](https://github.com/golang-migrate/migrate/releases) ou [igual](https://github.com/golang-migrate/migrate/releases/tag/v4.15.0) a `4.15.0`.
| `Docker CLI` | `>= v20.10.0` | Versão [superior](https://github.com/docker/cli/tags) ou [igual](https://github.com/docker/cli/releases/tag/v20.10.0) a `20.10.0`.
| `Docker Compose` | `>= v2.0.0` | Versão [superior](https://github.com/docker/compose/releases) ou [igual](https://github.com/docker/compose/releases/tag/v2.0.0) a `2.0.0`.

## Como usar e configurar o projeto

### Variáveis de Ambiente

Crie um arquivo `.env` com as configurações necessárias. Você pode usar o arquivo de exemplo `.env.example` como referência.

```bash
cp cmd/server/.env.example cmd/server/.env
```

### Migrações

Para gerenciar migrações do banco de dados `MongoDB`, você pode usar a ferramenta `migrate`. A seguir estão as instruções para criar, subir e voltar migrações.

#### 1. Criar nova migração

Para criar uma nova migração, execute o comando abaixo a partir do diretório raiz do projeto:

```bash
migrate create --ext json --dir "internal/infra/data/client/mongo/migration" --tz UTC nome_da_migração
```

#### 2. Subir Migrações

Para aplicar as migrações, execute o comando abaixo a partir do diretório raiz do projeto:

```bash
migrate --path="internal/infra/data/client/mongo/migration" --database "mongodb://user:password@host:port/dbname?ssl=false&authSource=admin" up
```

#### 3. Voltar Migrações

Para reverter as migrações, execute o comando abaixo a partir do diretório raiz do projeto:

```bash
migrate --path="internal/infra/data/client/mongo/migration" --database "mongodb://user:password@host:port/dbname?ssl=false&authSource=admin" down numeroDeMigraçõesParaReverter
```

### Executando Localmente

#### 1. Subir banco de dados

Para subir o banco de dados, execute o comando abaixo a partir do diretório raiz do projeto:

```bash
docker compose -f cmd/server/docker-compose.yml up mongo mongo-express
```

#### 2. Executar o backend

Para rodar o backend do projeto, execute o comando abaixo a partir do diretório raiz do projeto:

```bash
go run cmd/server/main.go
```

### Executando com Docker Compose

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

## Acessando a Documentação

A documentação completa da API pode ser acessada através do seguinte link:
[http://localhost:3000/docs/index.html](http://localhost:3000/docs/index.html)

### Gerando a Documentação com Swagger

Para gerar a documentação da API utilizando Swagger, siga os passos abaixo:

#### 1. Instale o Swagger CLI

Se ainda não tiver o Swagger CLI instalado, você pode instalá-lo globalmente com o seguinte comando:

```sh
go install github.com/swaggo/swag/cmd/swag@latest
```

#### 2. Gere a Documentação

A partir do diretório raiz do projeto execute o comando

```sh
swag init -g cmd/server/main.go -ot go,yaml
```
