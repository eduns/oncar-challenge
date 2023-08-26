# OnCar Challenge

## BackEnd

Este projeto foi desenvolvido inspirado nas arquiteturas **Arquitetura Limpa** e **Ports and Adapters**,
para acelerar o processo de desenvolvimento, testes e manutenção dos componentes do sistemas

![Arch](assets/clean_arch.jpg)

### Tecnologias Utilizadas

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)

Linguagem rica em funcionalidades e recursos embutidos na biblioteca padrão para tarefas muito comuns, como servidores web e criptografia.

![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)

Permite que o desenvolvimento rápido, empacotamento e deploy de aplicações, sem a limitação de ambientes e dependência de instalar os serviços localmente.

![Postgres](https://img.shields.io/badge/postgres-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white)

Banco Relacional robusto, com grande extensibilidade de operadores, estruturas e tipos de dados, modularidade e excelente suporte à escalabilidade.

### Pré-requisitos

Antes de utilizar a aplicação, certifique-se de ter os seguintes recursos instalados:

- Docker
- Docker Compose
- Linguagem Go

### Endpoints

>Para listagem de veículos

```http
GET /vehicles HTTP/1.1
Host: localhost:8080
Content-Type: application/json
```

>Para cadastrar um lead

```http
POST /leads HTTP/1.1
Host: localhost:8080
Content-Type: application/json
```

### Parâmetros no corpo da requisição

```json
{
  "name":"Some name",
  "email":"mail@mail.com",
  "phone":"+5512881437331",
  "chosenVehicleId":"1da55e15-f253-488d-8590-e0c8a955898b"
}
```

- **name**: Nome do lead
- **email**: E-mail do lead
- **phone**: Telefone do lead
- **chosenVehicleId**: Id de um veículo da lista de veículos

---

## FrontEnd

### Tecnologias Utilizadas

![React](https://img.shields.io/badge/react-%2320232a.svg?style=for-the-badge&logo=react&logoColor=%2361DAFB)

Possui grande flexibilidade para criação e reuso de componentes, eficiência nas atualizações de páginas e é amigável à SEO.

![Vite](https://img.shields.io/badge/vite-%23646CFF.svg?style=for-the-badge&logo=vite&logoColor=white)

Acelera muito o desenvolvimento de páginas web, criando uma estrutura de trabalho ordenada e performática.
Depois, execute

---

## Testes

Para executar os testes da aplicação, no diretório raiz, execute:

```bash
make run-tests
```

### Executando os serviços do projeto

Antes de tudo, preencha os valores das váriáveis nos arquivos **.env.example** e **.env.database.example**

```bash
make run-app
```
