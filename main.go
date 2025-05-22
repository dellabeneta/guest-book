package main

import (
    "embed"
    "database/sql"
    "fmt"
    "html/template"
    "log"
    "net/http"
    "time"

    _ "github.com/mattn/go-sqlite3"
)

//go:embed templates
var embeddedTemplates embed.FS

//go:embed static
var staticFiles embed.FS

type Message struct {
    Content   string
    CreatedAt time.Time
}

func main() {
    // Criar/Carregar banco de dados
    db, err := sql.Open("sqlite3", "guestbook.db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Criar tabela (se não existir)
    _, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS messages (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            content TEXT,
            created_at DATETIME DEFAULT CURRENT_TIMESTAMP
        )
    `)
    if err != nil {
        log.Fatal(err)
    }

    // Servir favicon.ico diretamente
    http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
        data, err := staticFiles.ReadFile("static/favicon.ico")
        if err != nil {
            http.NotFound(w, r)
            return
        }
        w.Header().Set("Content-Type", "image/x-icon")
        w.Write(data)
    })

    // Configurar servidor
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if r.Method == "POST" {
            message := r.FormValue("message")
            if message != "" {
                _, err := db.Exec("INSERT INTO messages (content) VALUES (?)", message)
                if err != nil {
                    http.Error(w, err.Error(), http.StatusInternalServerError)
                    return
                }
            }
            http.Redirect(w, r, "/", http.StatusSeeOther)
            return
        }

        // Buscar mensagens
        rows, err := db.Query("SELECT content, created_at FROM messages ORDER BY created_at DESC")
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        defer rows.Close()

        var messages []Message
        for rows.Next() {
            var m Message
            err := rows.Scan(&m.Content, &m.CreatedAt)
            if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }
            messages = append(messages, m)
        }

        // Carregar template
        tmpl, err := template.ParseFS(embeddedTemplates, "templates/index.html")
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        // Renderizar
        tmpl.Execute(w, messages)
    })

    // Iniciar servidor
    fmt.Println("Servidor rodando em http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}