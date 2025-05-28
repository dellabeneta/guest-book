## Guest Book

Uma aplicação web simples de livro de visitas feita em Go, utilizando SQLite e templates HTML.

Ambiente de Staging e Produção, em branches separadas.

### Pré-requisitos

- [Go](https://golang.org/dl/) 1.16 ou superior
- [SQLite3](https://www.sqlite.org/download.html) (opcional, pois o Go criará o banco automaticamente)

### Como rodar

1. **Clone o repositório:**

   ```sh
   git clone <URL-do-repositório>
   cd guest-book
   ```

2. **Instale as dependências:**

   ```sh
   go mod tidy
   ```

3. **Execute a aplicação diretamente:**

   ```sh
   go run main.go
   ```

   **Ou, se preferir, compile para um binário:**

   ```sh
   go build -o guestbook main.go
   ./guestbook
   ```

   > **Atenção:** O arquivo `guestbook.db` (banco de dados) será criado no mesmo diretório do binário. Para manter a persistência dos dados, mantenha o arquivo `guestbook.db` junto do executável.

4. **Acesse no navegador:**

   Abra [http://localhost:8080](http://localhost:8080)

### Estrutura do Projeto

- `main.go`: Código principal da aplicação.
- `templates/`: Contém o template HTML (`index.html`).
- `guestbook.db`: Banco de dados SQLite criado automaticamente.
- `migrations/`: Scripts SQL para migração (opcional).

### Funcionalidades

- Envie mensagens pelo formulário na página principal.
- As mensagens são salvas no banco de dados e exibidas em ordem decrescente de data.

### Observações

- O banco de dados é criado automaticamente no arquivo `guestbook.db` na primeira execução.
- Para redefinir as mensagens, basta apagar o arquivo `guestbook.db`.


Feito com Go e amor 💚