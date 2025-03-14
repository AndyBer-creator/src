package main

import (
	"errors"
	"fmt"
	"net/http"
	"rest_api/pkg/models"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Received %s request for home\n", r.Method)
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	s, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	// Используем помощника render() для отображения шаблона.
	app.render(w, r, "home.page.tmpl", &templateData{
		Snippets: s,
	})
}

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Received %s request for showSnippet\n", r.Method)
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w) // Страница не найдена.
		return
	}

	// Вызываем метода Get из модели Snipping для извлечения данных для
	// конкретной записи на основе её ID. Если подходящей записи не найдено,
	// то возвращается ответ 404 Not Found (Страница не найдена).
	s, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}
	// Используем помощника render() для отображения шаблона.
	app.render(w, r, "show.page.tmpl", &templateData{
		Snippet: s,
	})
}

func (app *application) createSnippetForm(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Received %s request for createSnippetForm\n", r.Method)
	app.render(w, r, "create.page.tmpl", nil) //&templateData{
	// 	//Snippets: s,
	// })
	// if r.Method != http.MethodGet {
	// 	w.Header().Set("Allow", http.MethodGet)
	// 	app.clientError(w, http.StatusMethodNotAllowed)
	// 	return
	// }

	// err := app.render(w, r, "create.page.tmpl", nil)
	// if err != nil {
	// 	fmt.Printf("Error rendering create page: %v\n", err)
	// 	app.serverError(w, err)
	// }
}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Received %s request for createSnippet\n", r.Method)
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	// Получаем данные из формы
	title := r.FormValue("title")
	content := r.FormValue("content")
	expires := r.FormValue("expires")

	// Передаем данные в метод SnippetModel.Insert(), получая обратно
	// ID только что созданной записи в базу данных.
	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// Перенаправляем пользователя на соответствующую страницу заметки.
	http.Redirect(w, r, fmt.Sprintf("/snippet?id=%d", id), http.StatusSeeOther)
}
