# Unico - Feira livre
## Pré-requisitos
1 - Docker          20.10.17 ou superior
2 - Docker Compose  1.25.0 ou superior

#### Para Rodar a aplicação
Vamos executar em dois passos para evitar que aplicação suba antes do banco estar completamente disponível

1 - Crie a network
```shell
docker network create unico_network
```

2 - Para subir o banco de dados e o pgAdmin
```shell
docker-compose up -d db pgAdmin
```

3 - Para subir a aplicação
```shell
docker-compose up app
```

#### Para visualização dos logs
```shell
tail -f log.txt
```
#### Execução dos tests
```shell
go test -coverprofile cover.out  -v ./...
```

#### Tests via postman
Na raiz do repository há uma collection do postman a ser importada e por ser utilizada para testes. Nome arquivo: Fairs.postman_collection.json

Observação: Crie e utilize o Environment do postman para configurar o host a ser utilizado. Veja o exemplo:
![alt text](https://github.com/Deivisson/free_fairs/blob/master/postman_environment_default.png?raw=true)
![alt text](https://github.com/Deivisson/free_fairs/blob/master/postman_environment_debug.png?raw=true)

### TODO
1 - Documentar os EndPoints no Swagger para uma melhor compreensão das mesmas.
2 - Criar shell Script para iniciar a aplicação e não precisar criar a network manualmente
3 - Melhorar a cobertura de testes e explorar outros tipos de testes como de integração

