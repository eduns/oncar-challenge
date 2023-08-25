# OnCar Challenge

## Backend

Este projeto foi desenvolvido inspirado nas arquiteturas **Arquitetura Limpa** e **Ports and Adapters**,
para acelerar o processo de desenvolvimento, testes e manutenção dos componentes do sistemas

![Arch](assets/clean_arch.jpg)

### Tecnologias Utilizadas

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)
![Postgres](https://img.shields.io/badge/postgres-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white)

### Pré-requisitos

Antes de utilizar a aplicação, certifique-se de ter os seguintes recursos instalados:

- Docker
- Docker Compose
- Linguagem Go

### Testes

Para executar os testes da aplicação, entre no diretório **backend** e execute:

```bash
make run-tests
```

### Executando a aplicação

No diretório **backend**, execute:

```bash
make run-app
```

### Endpoints

>Para listagem de veículos

```http
GET /vehicles HTTP/1.1
Host: localhost:8080
Content-Type: application/json
```

---

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
