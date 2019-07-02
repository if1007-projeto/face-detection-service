# Face Detection Processor

Serviço responsável por ler os frames do kafka, detectar todas as faces presentes no frame e publicar novamente no kafka.

## Requisitos

A aplicação foi testada com as ferramentas e versões abaixo.

* go - 1.10.4

### Bibliotecas

* [pigo](https://github.com/esimov/pigo)
* [gg](https://github.com/fogleman/gg)

## Instalação

Após instalar o go, instale as bibliotecas necessárias utilizando o comando abaixo:

```bash
go get -u github.com/esimov/pigo/cmd/pigo
go get -u github.com/fogleman/gg
go get -u github.com/disintegration/imaging
go get -u github.com/lovoo/goka
```

## Como utilizar

Primeiro, suba os containers do kafka utilizando o comando abaixo.

```bash
docker-compose up
```

Depois disso, compile a aplicação e execute.

```bash
go build
./face-detection-processor
```
