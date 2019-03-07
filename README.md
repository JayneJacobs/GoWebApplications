### Templates in Go

# Tol install local assets:

1. install NodeJs(https://nodejs.org)
2. install `grunt-cli` globally via `npm install -g grunt-cli`
3. run `npm install` from root directory of the project
4. run `grunt` from the root of the project to compile CSS

## to run the application
1. install `http0server` via `npm install -g http-server`
2.[ run `http-server` in applications root. ]* navigate to `http://localhost:8080/html/home.html`
''
## HandleFunc
type HandlerFunc func(ResponseWriter, * Request)
#### has method
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request)

### notFoundHandler

func NOtFoundHandler Handler - returns 404

### RedirectHandler (similar to middle ware)
func RedirectHJandler(url string, code int) Handler

### StripPrefix
func StripPrefix(prefix string, h handler) Handler

### FileServer
func FileServer(root FileSystem) Handler
type FileSystem interface {
    Open(name string) (File, error)
}

type Dir string
func (d Dir) Open(name string) (File, error)


# Templates

### Example
Dear {{.FirstName}}

We would like to personally invite you to {{.Event}}!

Sincerely
{{.AuthorName}}




type Data struct {
    FirstName string
    Event string
    AuthorName string
}