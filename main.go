package main

import ("fmt"
        "net/http")

func home_page(page http.ResponseWriter, r *http.Request){
  fmt.Fprintf(page, "Я хочу отдохнуть")
}

func other_page(page http.ResponseWriter, r *http.Request){
  fmt.Fprintf(page, "Рандомная инфа")
}

func main() {
  http.HandleFunc("/", home_page)//создание главной страницы
  http.HandleFunc("/sleep/", other_page) //создание дополнительной страницы
  http.ListenAndServe(":6485", nil)//Порт, и дополнительные настройки сервера (в данном случае nil)
}
