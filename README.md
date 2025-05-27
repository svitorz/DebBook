````md
# DevBook 📘

API RESTful desenvolvida em **Go (Golang)** com foco em performance, boas práticas e escalabilidade. Inspirado em redes sociais, o DevBook permite que usuários se cadastrem, publiquem conteúdo, sigam outros perfis e curtam publicações.

## ✨ Funcionalidades

- ✅ Autenticação com **JWT**
- ✅ CRUD completo de usuários
- ✅ Sistema de **publicações**
- ✅ Relacionamentos de **seguidores** (follow/unfollow)
- ✅ Sistema de **likes** (curtidas)
- ✅ Arquitetura limpa com **padrão MVC** e camada de **repositórios**
- ✅ **Middlewares** para logging e autenticação
- ✅ Banco de dados relacional (**MySQL**)
- ✅ Integração completa com **Docker** via `docker-compose`

---

## 🚀 Tecnologias Utilizadas

- [Go (Golang)](https://golang.org)
- [MySQL](https://www.mysql.com/)
- [Docker](https://www.docker.com/)
- [JWT (JSON Web Token)](https://jwt.io/)
- [Mux Router](https://github.com/gorilla/mux)
- [GORM (opcional, caso use)](https://gorm.io/)

---

## 📦 Como executar

### Pré-requisitos

- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/)

### Clone o projeto

```bash
git clone https://github.com/svitorz/DevBook.git
cd DevBook
```
````

### Crie o arquivo `.env`

```env
DB_USER=root
DB_PASS=mydba
DB_NAME=devbook
DB_PORT=3307
DB_HOST=127.0.0.1
API_PORT=5000
```

### Execute com Docker Compose

```bash
docker-compose up --build
```

A API estará disponível em: [http://localhost:5000](http://localhost:5000)

---

## 🧪 Endpoints principais

- `POST /login` → autenticação e geração de token JWT

- `POST /usuarios` → cadastro de usuário

- `GET /usuarios` → listagem de usuários

- `GET /usuarios/{id}` → detalhes de usuário

- `PUT /usuarios/{id}` → edição

- `DELETE /usuarios/{id}` → exclusão

- `POST /publicacoes` → criar publicação

- `GET /publicacoes` → listar publicações

- `POST /publicacoes/{id}/curtir` → curtir

- `POST /publicacoes/{id}/descurtir` → remover curtida

- `POST /usuarios/{id}/seguir` → seguir outro usuário

- `POST /usuarios/{id}/parar-de-seguir` → deixar de seguir

---

## 🔐 Autenticação JWT

Ao fazer login, você recebe um token JWT que deve ser enviado no header:

```http
Authorization: Bearer seu_token_jwt
```

---

## 💡 Contribuindo

Pull requests são bem-vindos! Para mudanças maiores, abra uma issue para discutirmos o que você deseja modificar.

---

## 🧑‍💻 Autor

Desenvolvido por **[Vitor Fábio de Castro Souza](https://www.linkedin.com/in/svitorz)**
📧 [vitor.fabio.castro@gmail.com](mailto:vitor.fabio.castro@gmail.com)
🐙 [github.com/svitorz](https://github.com/svitorz)

---

## 📜 Licença

Este projeto está licenciado sob a licença MIT - veja o arquivo [LICENSE](LICENSE) para mais detalhes.

```

---

```
