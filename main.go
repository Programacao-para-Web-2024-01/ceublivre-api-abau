package main

import (
    "bufio"
    "encoding/json"
    "fmt"
    "html/template"
    "net/http"
    "os"
)

var (
    tmplRegistro *template.Template
    tmplLogin    *template.Template
    tmplIndex    *template.Template
)

func init() {
    tmplRegistro = template.Must(template.ParseFiles("registro.html"))
    tmplLogin = template.Must(template.ParseFiles("login.html"))
    tmplIndex = template.Must(template.ParseFiles("index.html"))
}

type Usuario struct {
    ID       string `json:"id"`
    Nome     string `json:"nome"`
    Email    string `json:"email"`
    Password string `json:"password"`
}

func salvarUsuario(usuario Usuario) error {
    file, err := os.OpenFile("usuarios.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        return err
    }
    defer file.Close()

    dadosJSON, err := json.Marshal(usuario)
    if err != nil {
        return err
    }

    _, err = file.Write(dadosJSON)
    if err != nil {
        return err
    }

    file.Write([]byte("\n"))

    return nil
}

func lerUsuarios() ([]Usuario, error) {
    var usuarios []Usuario

    file, err := os.Open("usuarios.json")
    if err != nil {
        return usuarios, err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        var usuario Usuario
        err := json.Unmarshal(scanner.Bytes(), &usuario)
        if err != nil {
            return usuarios, err
        }
        usuarios = append(usuarios, usuario)
    }

    return usuarios, scanner.Err()
}

func RegisterUsuario(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        tmplRegistro.Execute(w, nil)
        return
    }

    usuario := Usuario{
        ID:       r.FormValue("id"),
        Nome:     r.FormValue("nome"),
        Email:    r.FormValue("email"),
        Password: r.FormValue("password"),
    }

    err := salvarUsuario(usuario)
    if err != nil {
        http.Error(w, "Erro ao salvar usuário", http.StatusInternalServerError)
        return
    }

    http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func ListarUsuariosAPI(w http.ResponseWriter, r *http.Request) {
    usuarios, err := lerUsuarios()
    if err != nil {
        http.Error(w, "Erro ao ler usuários", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(usuarios)
}

func Login(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        tmplLogin.Execute(w, nil)
        return
    }

    email := r.FormValue("email")
    password := r.FormValue("password")

    usuarios, err := lerUsuarios()
    if err != nil {
        http.Error(w, "Erro ao ler usuários", http.StatusInternalServerError)
        return
    }

    for _, usuario := range usuarios {
        if usuario.Email == email && usuario.Password == password {
            http.Redirect(w, r, "/index", http.StatusSeeOther)
            return
        }
    }

    tmplLogin.Execute(w, struct{ Error bool }{true})
}

func Index(w http.ResponseWriter, r *http.Request) {
    tmplIndex.Execute(w, nil)
}

func main() {
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
    http.HandleFunc("/", RegisterUsuario)
    http.HandleFunc("/api/usuarios", ListarUsuariosAPI)
    http.HandleFunc("/login", Login)
    http.HandleFunc("/index", Index)

    fmt.Println("Servidor iniciado em http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}

