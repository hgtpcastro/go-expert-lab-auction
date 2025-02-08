# go-expert-lab-auction
Concorrência com Golang - Leilão

A configuração é realizada através de variáveis de ambiente declaradas no arquivo `.env`

## Configuração
Ajuste-o o arquivo `.env` presente na pasta _./cmd/auction_, conforme necessidade. Por padrão, os seguintes valores são utilizados:

```sh
BATCH_INSERT_INTERVAL=20s
MAX_BATCH_SIZE=4
AUCTION_INTERVAL=20s

MONGO_INITDB_ROOT_USERNAME: admin
MONGO_INITDB_ROOT_PASSWORD: admin
MONGODB_URL=mongodb://admin:admin@mongodb:27017/auctions?authSource=admin
MONGODB_DB=auctions
```

### Buildar a imagem docker e inicar a aplicação
```bash
    make start
```

### Parar a aplicação
```bash
    make stop
```

### Remover os containers
```bash
    make clean
```

### Rodar os testes de unidade
```bash
    make test
```

### Exemplos de requisições via arquivo .http (VSCode: REST Client Plugin)

Navegue até a pasta api no diretório raiz do projeto


```sh
auction.http
```

## <a name="license"></a> License

Copyright (c) 2025 [Hugo Castro Costa]

[Hugo Castro Costa]: https://github.com/hgtpcastro
