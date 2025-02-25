# Golang Champions 🏆

Um sistema de gerenciamento de campeonatos desenvolvido em Go, permitindo a criação e administração de competições esportivas.

## Funcionalidades

- Criar novo campeonato
- Adicionar atletas
- Criar jogos entre atletas
- Listar atletas participantes
- Acompanhamento de pontuação e desempenho

## Estrutura do Projeto

O projeto está organizado nos seguintes pacotes:

- `athlete`: Gerenciamento de atletas e suas estatísticas
- `championship`: Lógica do campeonato
- `game`: Controle de jogos e partidas
- `main.go`: Ponto de entrada da aplicação

## Como Executar

1. Certifique-se de ter Go 1.22.1 ou superior instalado
2. Clone o repositório: `git clone https://github.com/chrisviana/golang-champions.git`
3. Entre no diretório do projeto: `cd golang-champions`
4. Execute o programa: `go run main.go`


## Menu de Opções

1. Criar novo campeonato
2. Adicionar atleta
3. Criar jogo
4. Listar atletas
5. Sair

## Estrutura de Dados

### Atleta
- Nome
- Pontos ganhos
- Número de vitórias
- Número de derrotas
- Saldo inicial
- Jogos disputados
- Performance

### Campeonato
- Nome
- Data de início
- Data de término
- Número de atletas
- Código hash único
- Lista de atletas
- Lista de jogos

## Contribuição

Sinta-se à vontade para contribuir com o projeto através de pull requests ou reportando issues.

## Licença

Este projeto está sob a licença MIT.
