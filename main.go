package main

import (
	"fmt"
	"io"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// 1. Проверяем метод запроса (для GET тела обычно нет)
	if r.Method != http.MethodPost && r.Method != http.MethodPut {
		println("Метод не поддерживается", http.StatusMethodNotAllowed)

	}
	println(r.RequestURI)
	// 2. Читаем тело запроса
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Ошибка чтения тела запроса", http.StatusBadRequest)
		return
	}
	defer r.Body.Close() // Важно закрыть тело!

	// 3. Выводим тело
	println("Получено тело: %s\n", body)

	// 4. Отправляем ответ
	fmt.Fprintf(w, "Тело запроса получено: %s", body)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
