# Arquitetura limpa GO

Registrar pedidos, calcular o preço total e retornar uma lista de pedidos. Foi prejetado utiliaznao Clean Arquitecture implementando modelos de comunicação REST , GRPC e GraphQl alem de criar os eventos de  mensagens de criação de pedidos no RabbitMQ 

Abaixo os endpoints disponivels

Endpoint REST (POST / order e GET /orders)
Service CreateOrder e ListOrders com 
Query CreateOrder e ListOrders GraphQL


### Prerequisites

- [Docker]([https://www.docker.com/) - Necessario para subir Mysql e rabbitMQ

### Para rodar

PAra roadr a aplicação e subir todos os serviços basta executar:

`docker compose up -d`

### Executando a aplicação com REST 

Use o plugin do vscode api rest para fazer chamada utilizando os arquivos disponiveis na pasta API na raiz do projeto

Endpoints disponíveis:

POST create_order http://localhost:8000/order

GET list_orders http://localhost:8000/orders

- CreateOrder

![alt text](https://github.com/maxnet04/CleanArch/blob/main/images/rest-1.PNG)

- ListOrders

![alt text](https://github.com/maxnet04/CleanArch/blob/main/images/rest-2.PNG)


### Executando a aplicação com GraphQL

Abra o browser e va para: http://localhost:8080/

- CreateOrder

```graphql
mutation CreateOrder {
  createOrder(input:{ id: "bb", Price: 10.20, Tax: 10}) {
    id
    Price
    Tax
    FinalPrice
  }
}
```

![alt text](https://github.com/maxnet04/CleanArch/blob/main/images/graphi-1.PNG)

- ListOrders

```graphql
query ListOrder {
  listOrders {
    id
    Price
    Tax
    FinalPrice
  }
}
```

![alt text](https://github.com/maxnet04/CleanArch/blob/main/images/graphi-2.PNG)

### Executando a aplicação com gRPC

Instale o evans cli para testar grpc:
```bash
go install github.com/ktr0731/evans@latest
```

Execute apontando para o arquivo .proto
```bash
evans --proto internal/infra/grpc/protofiles/order.proto
```

![alt text](https://github.com/maxnet04/CleanArch/blob/main/images/grpc2.PNG)


Servicos disponiveis para serem chamados:

![alt text](https://github.com/maxnet04/CleanArch/blob/main/images/grpc.PNG)


### visualzando as mensagens com HabbitMQ

Acesse o RabbitMQ em: http://localhost:15672/

![alt text](https://github.com/maxnet04/CleanArch/blob/main/images/rabbit.PNG)