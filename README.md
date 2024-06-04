# Pós Go Expert

## Desafios

### Multithreads

Neste desafio você terá que usar o que aprendemos com Multithreading e APIs para buscar o resultado mais rápido entre duas APIs distintas.

As duas requisições serão feitas simultaneamente para as seguintes APIs:

`https://brasilapi.com.br/api/cep/v1/01153000 + cep`

`http://viacep.com.br/ws/" + cep + "/json/`

Os requisitos para este desafio são:

- Acatar a API que entregar a resposta mais rápida e descartar a resposta mais lenta.

- O resultado da request deverá ser exibido no command line com os dados do endereço, bem como qual API a enviou.

- Limitar o tempo de resposta em 1 segundo. Caso contrário, o erro de timeout deve ser exibido.

### Como testar?

- Via terminal, acesse a pasta do projeto;
- Execute o seguinte comando:
  - Substitua o valor `78550005` por outro CEP de sua preferencia.
```bash
go run main.go --cep 78550005
```
#### Comportamento esperado

- Se a API mais rápida for da BrasilAPI, a resposta deve se parecer com o exemplo abaixo:

```bash
Received content from.: Brasil API [https://brasilapi.com.br/]

Content.: {"cep":"78550005","state":"MT","City":"Sinop","neighborhood":"Setor Residencial Sul","street":"Rua das Tílias","service":"open-cep","location":{"type":"Point","coordinates":{"longitude":"-55.5006515","latitude":"-11.865263"}}}
```
- Caso seja a API da ViaCEP, a resposta deve se parecer com o exemplo abaixo:

```bash
Received content from.: Via CEP [https://viacep.com.br/]

Content.: {"cep":"78550-005","logradouro":"Rua das Tílias","complemento":"","bairro":"Setor Residencial Sul","localidade":"Sinop","uf":"MT","ibge":"5107909","ddd":"66","gia":"","siafi":"8985"}
```