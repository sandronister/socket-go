#!/bin/bash

# Endereço IP e porta do servidor TCP
SERVER_IP="127.0.0.1"
SERVER_PORT="3005"
SERVER_PORT2="3006"
SERVER_PORT3="3007"

# Mensagem a ser enviada
MESSAGE="Sua mensagem aqui ruptela"
MESSAGE2="Sua mensagem aqui ominicon"
MESSAGE2="Sua mensagem aqui sinocastle"

# Número de mensagens a serem enviadas
NUM_MESSAGES=1000

for i in $(seq 1 $NUM_MESSAGES); do
  echo $MESSAGE $i| nc $SERVER_IP $SERVER_PORT
  echo $MESSAGE2 $i| nc $SERVER_IP $SERVER_PORT2
  echo $MESSAGE3 $i| nc $SERVER_IP $SERVER_PORT3
done