# Face Detection Service

Serviço responsável por ler os frames do kafka, detectar todas as faces presentes no frame e publicar novamente no kafka.

## Requisitos

A aplicação foi testada com as ferramentas e versões abaixo.

* go - 1.10.4

### Bibliotecas

* [pigo](github.com/esimov/pigo)
* [gg](github.com/fogleman/gg)

## Instalação

Após instalar o go, instale as bibliotecas necessárias utilizando o comando abaixo:

```bash
go get -u github.com/esimov/pigo/cmd/pigo
go get -u github.com/fogleman/gg
```

## Como utilizar

```bash
go build
./pigo-sample
```