package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	Title string
	Body  []byte
}

// TODO: 不要ならあとで消す
// func (p *Page) save() error {
// 	filename := p.Title + ".txt"
// 	// 0600 = 権限(読み書き可能な状態)
// 	return ioutil.WriteFile(filename, p.Body, 0600)
// }

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

// テンプレートでPageの情報をHTMLへ出力
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, _ := template.ParseFiles(tmpl + ".html")
	t.Execute(w, p)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func main() {
	// ハンドラーの設定
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)

	// 受信待ち状態
	//  前処理で定義したパスに合致しない場合は終了する
	log.Fatal(http.ListenAndServe(":8080", nil))
}
