#!/bin/bash

sleep 10

echo "RabbitMQ iniciado com sucesso! Configurando a fila..."


#Declara uma fila
curl -u guest:guest -XPUT http://localhost:15672/api/queues/%2f/orders


#Vincula a fila a exchange
curl -u guest:guest -XPOST http://localhost:15672/api/bindings/%2f/e/amq.direct/q/orders