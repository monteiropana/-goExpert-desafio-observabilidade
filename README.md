# -goExpert-desafio-cleanArch
- Para subir a infra do banco de dados e RabbitMQ: docker compose up --build
- Para executar a aplicacao, entrar na pasta cmd: go build
- O arquivo executavel cmd sera criado, no linux/mac usar o comando: ./cmd e no windows: cmd.exe
- Para testar o endpoints HTTP, so executar o send Request em create_order.http e list_order.http
- GraphQL
Para testar a consulta ListOrders no GraphQL, execute a seguinte query, na porta http://localhost:8080/:

query {
  listOrders {
    id
    Price
    FinalPrice
    Tax
  }
}

gRPC
Para testar o serviço ListOrders com gRPC, use o seguinte comando no terminal:
grpcurl -plaintext -proto <caminho_para_protofile> <endereço_servidor> pb.OrderService/ListOrders
HTTP
