package main

import (
	"log"
	"net/http"
	"text/template"
	"path/filepath"
	"sync"
	"flag"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/signature"
	"github.com/stretchr/gomniauth/providers/facebook"
	"github.com/stretchr/gomniauth/providers/github"
	"github.com/stretchr/gomniauth/providers/google"
	"os"
	"github.com/stretchr/objx"
)

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

// ServeHTTP は HTTP リクエストを処理します
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	data := map[string]interface{} {
		"Host": r.Host,
	}
	if authCookie, err := r.Cookie("auth"); err == nil {
		data["UserData"] = objx.MustFromBase64(authCookie.Value)
	}
	t.templ.Execute(w, data) // 監訳注1: t.templ.Execute の戻り値はチェックすべきです。
}

func main() {
	var addr = flag.String("addr", ":8080", "アプリケーションのアドレス")
	flag.Parse() // フラグを解釈します
	// Gomniauth のセットアップ
	gomniauth.SetSecurityKey(signature.RandomKey(64))
	gomniauth.WithProviders(
		facebook.New(os.Getenv("FACEBOOK_APP_ID"), os.Getenv("FACEBOOK_APP_SECRET"), "http://localhost:8080/auth/callback/facebook"),
		github.New(os.Getenv("GITHUB_APP_ID"), os.Getenv("GITHUB_APP_SECRET"), "http://localhost:8080/auth/callback/github"),
		google.New(os.Getenv("GOOGLE_APP_ID"), os.Getenv("GOOGLE_APP_SECRET"), "http://localhost:8080/auth/callback/google"),
	)
	r := newRoom(UseAuthAvatar)
	http.Handle("/chat", MustAuth(&templateHandler{filename: "chat.html"}))
	http.Handle("/login", &templateHandler{filename: "login.html"})
	http.HandleFunc("/auth/", loginHandler)
	http.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		http.SetCookie(w, &http.Cookie{
			Name: "auth",
			Value: "",
			Path: "/",
			MaxAge: -1,
		})
		w.Header()["Location"] = []string{"/chat"}
		w.WriteHeader(http.StatusTemporaryRedirect)
	})
	http.Handle("/room", r)

	// チャットルームを開始します
	go r.run()

	// Web サーバを開始します
	log.Println("Webサーバを開始します。ポート: ", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
