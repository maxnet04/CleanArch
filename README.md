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

![alt text](/docs/rest-1.png)

- ListOrders

![alt text](/docs/rest-2.png)


### Executando a aplicação com GraphQL

Abra o browser e va para: http://localhost:8080/

- CreateOrder

`mutation createOrder {
  createOrder(input: {Id: "tr6544y", Price:185.77, Tax:15.33}) {
    Id
    Price
    Tax
    FinalPrice
  }
}`

![alt text](/docs/graphi-1.png)

- ListOrders

`query listOrders {
  listOrders{
    Id
    Price
    Tax
    FinalPrice
  }
}`

![alt text](/docs/graphi-2.png)

### Executando a aplicação com gRPC

Instale o evans cli para testar grpc:
```bash
go install github.com/ktr0731/evans@latest
```

Execute apontando para o arquivo .proto
```bash
evans --proto internal/infra/grpc/protofiles/order.proto
```

![alt text](/docs/grpc2.png)


Servicos disponiveis para serem chamados:

![alt text](/docs/grpc1.png)


### visualzando as mensagens com HabbitMQ

Acesse o RabbitMQ em: http://localhost:15672/

![alt text](/docs/rabbit.png)