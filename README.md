# Prova Back-End

> Maior número romano em uma palavra

### Ajustes e melhorias

O projeto ainda está em desenvolvimento e as próximas atualizações serão voltadas nas seguintes tarefas:

- [x] API GraphQL em GO
- [x] Mutation Search
- [x] Lógica da Mutation
- [x] Conteinerizar Aplicação 



## ⚙️ Tecnologias

* [GO 1.18](https://go.dev/)
* [Docker](https://www.docker.com/)
* [Visual Studio Code](https://code.visualstudio.com/)
* [GraphQL](https://graphql.org/)

## 💻 Pré-requisitos

Antes de começar, verifique se você atendeu aos seguintes requisitos:
* Você instalou a versão mais recente de `Docker`


## 🚀 Instalando

Para instalar o projeto, siga estas etapas:

Dentro do projeto, execute o comando para criar uma nova imagem a partir da Dockerfile:
```
docker build --tag=prova-backend:1.0 .
```


## ☕ Usando

Para rodar o projeto, siga estas etapas:

```
docker run -p 8080:8080 prova-backend:1.0

```
Você pode acessar a rota http://localhost:8080/graphql e usar o playground do GraphQL.
Um exemplo de Query para testar:

```graphql
mutation {
 search(text: "AXXBLX") {
  number
  value
 }
}
```
