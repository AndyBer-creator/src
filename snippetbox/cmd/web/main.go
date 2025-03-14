package main

import (
	"database/sql"
	"flag"
	"html/template" // Новый импорт
	"log"
	"net/http"
	"os"
	"snippetbox/pkg/models/pgsql"

	_ "github.com/lib/pq"
)

// Добавляем поле templateCache в структуру зависимостей. Это позволит
// получить доступ к кэшу во всех обработчиках.
type application struct {
	errorLog      *log.Logger
	infoLog       *log.Logger
	snippets      *pgsql.SnippetModel
	templateCache map[string]*template.Template
}

func main() {
	addr := flag.String("addr", ":4000", "Сетевой адрес веб-сервера")
	dsn := flag.String("postgres", "postgres://web:pass@/postgres?sslmode=disable", "Название SQL источника данных")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	// Инициализируем новый кэш шаблона...
	templateCache, err := newTemplateCache("./ui/html/")
	if err != nil {
		errorLog.Fatal(err)
	}

	// И добавляем его в зависимостях нашего
	// веб-приложения.
	app := &application{
		errorLog:      errorLog,
		infoLog:       infoLog,
		snippets:      &pgsql.SnippetModel{DB: db},
		templateCache: templateCache,
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Запуск сервера на http://127.0.0.1%s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
