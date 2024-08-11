#!/bin/bash

# Endereço IP e porta do servidor TCP
SERVER_IP="127.0.0.1"
SERVER_PORT="3005"

# Mensagem a ser enviada
MESSAGE="Sua mensagem aqui"

# Número de mensagens a serem enviadas
NUM_MESSAGES=1000

for i in $(seq 1 $NUM_MESSAGES); do
  echo $MESSAGE $i| nc $SERVER_IP $SERVER_PORT
done