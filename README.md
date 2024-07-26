# Term Alarms | Backend

<p align="center">
  <a href="https://itrackbrasil.com.br/solucoes/isend-tms/" target="_blank">
    <img src="https://term-monitor.s3.us-east-1.amazonaws.com/monitor.png?response-content-disposition=inline&X-Amz-Security-Token=IQoJb3JpZ2luX2VjEAcaCXVzLWVhc3QtMSJHMEUCIC1RZWnb36mC4FQ92K43wB7m2LWcmV3ke3oON3da8g3ZAiEAy50xgy7qx6nHO11bZWwkzNpzdJ8zxpkp8ixUa8OKFlEq7QII4P%2F%2F%2F%2F%2F%2F%2F%2F%2F%2FARAAGgw0MjA5NDc3NjEyOTkiDHLHyH3wNXciR%2BzTZyrBApnClWfiKfoNNBWvNukTcFuj3PXksai7iHIHWbFZ3oUmJbgLOwPs4kzTYmQw23N63bTevfjFb7TqGaaaB2PsS9Yapq23f0DuY5l0k0eAzXF2HPZHkF0Sda0HmrQBfm%2FnFVRZHQ5lJV1%2BmFrRvkb6nrWZfXdro68O0mSnZWACgQDmB4YdNPAJeaiD%2Fq057voIWlASHuWw7zo8C%2B35cpn0Vxv8k2MY9ii5MExeIOkz%2BOBGi2nopy8l1TqF10ZHxxk19qeIkkTfn6YixHRdTJH52pqbrL6GuM4fiFaAdo0KpX%2BIJ%2FX2985v45Z9lRRB%2Fv5VSnTYmqm3QC53RJ99qr00JMFNJ9wrxU%2BKrTOeXzm%2BmvMbykYs7nZPtjyjFwEj9G2jcto6PsY%2F2ouAf0ARkbUbGpjKdFVakz5C7tP7XJ%2BMYDgH1zD7x5C1BjqzAjOYzphNVeetXOs9Be5fiNBGPuSI59nlF9mTW62UKqSxpnVOaC7LBx8Rjt2arsnv1uyTlE3w1sUoReGW7ofaTydtRLnKcTTFB62%2F6USFopiGMgfvOsHDxnEgFLP2Wryy%2BXo9U2BezxbO0llCI8C8yPE8QY0IByTTO2yetbc9cQxJ%2F7NVFtXGOrHrWhjyyVOQ2z0YiJddxEm0Rb4%2FqISGFmQi%2BDc2mt837B1T6JKMNi8Ja3FEBTcP%2BjG%2FqcG1lwHIidql9lp0%2BSGLKthi5PhqI9rofGhGpudGK%2F87gON08pracWFX8nRXlntYTUNLjtPtw6Uzmn%2FU6UVe8GNl0kSDcmGJ5t78XrlQxRkacZBuUzmTobfsqaQtYfbcOg5AUyK40gDkCCMbzfrSaZCU0dWotKEd53s%3D&X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20240726T230010Z&X-Amz-SignedHeaders=host&X-Amz-Expires=300&X-Amz-Credential=ASIAWEATRDCJWEYS7CQ2%2F20240726%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Signature=c251d285b0cb0077ce36e59619f80bb7882b8ed1067043e363072cddc506f097" alt="Term Alarms" height="150" width=325"/>
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

### 1. Rodar o Projeto

Para rodar o backend do projeto, execute o comando abaixo a partir do diretório raiz do projeto:

```bash
go run cmd/server/main.go
```
