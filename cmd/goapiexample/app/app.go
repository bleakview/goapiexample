package app

import (
	"log"
	"net/http"
	"os"

	"github.com/bleakview/goapiexample/cmd/goapiexample/handler"
	"github.com/bleakview/goapiexample/cmd/goapiexample/initdb"
	"github.com/bleakview/goapiexample/cmd/goapiexample/model"
	"github.com/glebarez/sqlite"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	_ "github.com/bleakview/goapiexample/cmd/goapiexample/docs"
)

// App has router and db instances
type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

// Initilize app
func (app *App) Initialize() {
	var db *gorm.DB
	var err error

	//if no mysql connection string found then switch to sqlite
	mySqlDsn, isMySqlDsnSet := os.LookupEnv("DB_CONNECTION_URL")
	if isMySqlDsnSet {
		db, err = gorm.Open(mysql.Open(mySqlDsn), &gorm.Config{})
	} else {
		db, err = gorm.Open(sqlite.Open("books.db"), &gorm.Config{})
	}

	if err != nil {
		log.Fatal("Could not connect database")
	}

	//create default tables if required
	app.DB = model.DBMigrate(db)
	initdb.InitDB(db)
	app.Router = mux.NewRouter()
	// add routing for swagger docs
	app.Router.PathPrefix("/docs/").Handler(httpSwagger.WrapHandler)
	app.setRouters()
}

// Set all required routers
func (app *App) setRouters() {
	// Routing for handling the projects
	app.Get("/books", app.GetAllBooks)
	app.Get("/books/{id}", app.GetBookById)
	app.Post("/books/", app.CreateBook)
	app.Put("/books/{id}", app.UpdateBook)
}

func (app *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	app.Router.HandleFunc(path, f).Methods("GET")
}

func (app *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	app.Router.HandleFunc(path, f).Methods("POST")
}

func (app *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	app.Router.HandleFunc(path, f).Methods("PUT")
}

// @Summary Gets all books
// @Description Get all recorded books
// @ID GetAllBooks
// @Accept  json
// @Produce  json
// @Success 200 {array} model.Book "Book array"
// @Router /books [get]
func (app *App) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	handler.GetAllBooks(app.DB, w, r)
}

// GetStringByInt example
// @Summary Gets a book
// @Description Gets a single book by id
// @ID GetBookById
// @Accept  json
// @Produce  json
// @Param   id      path   string     true  "Book id"
// @Success 200 {object} model.Book	"found book"
// @Router /books/{id} [get]
func (app *App) GetBookById(w http.ResponseWriter, r *http.Request) {
	handler.GetBook(app.DB, w, r)
}

// @Summary Create Book
// @Description Create single book with given parameters
// @ID CreateBook
// @Accept  json
// @Produce  json
// @Param   book      body model.BookRequest true  "Book data"
// @Success 200 {object} model.Book string	""
// @Router /books/ [post]
func (app *App) CreateBook(w http.ResponseWriter, r *http.Request) {
	handler.CreateBook(app.DB, w, r)
}

// @Summary Update book
// @Description Update book with give id
// @ID UpdateBook
// @Accept  json
// @Produce  json
// @Param   id      path   string     true  "Book id"
// @Param   book      body model.BookRequest true  "Book data"
// @Success 200 {string} string	""
// @Router /books/{id} [put]
func (app *App) UpdateBook(w http.ResponseWriter, r *http.Request) {
	handler.UpdateBook(app.DB, w, r)
}

func (app *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, app.Router))
}
