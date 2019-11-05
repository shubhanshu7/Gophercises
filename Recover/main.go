package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"runtime/debug"
	"strconv"
	"strings"

	"github.com/alecthomas/chroma/formatters/html"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
	"github.com/gorilla/mux"
)

func main() {
	mux := mux.NewRouter()
	mux.HandleFunc("/debug/", sourceCodeHandler).Methods("GET")
	mux.HandleFunc("/panic/", panicDemo).Methods("GET")
	mux.HandleFunc("/panic-after/", panicAfterDemo).Methods("GET")
	mux.HandleFunc("/", hello)
	log.Fatal(http.ListenAndServe(":3000", devMw(mux)))
}

//sourceCodeHandler shows the requested page in the browser.
//Chroma takes source code and other structured text and converts it into syntax highlighted HTML.
func sourceCodeHandler(w http.ResponseWriter, r *http.Request) {
	path := r.FormValue("path") //pass the path in the URL
	lineStr := r.FormValue("line")
	line, err := strconv.Atoi(lineStr) //converting the string into int to get line no
	if err != nil {
		line = -1
	}
	file, err := os.Open(path) //opens the path
	b := bytes.NewBuffer(nil)
	_, err = io.Copy(b, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var lines [][2]int
	if line > 0 {
		lines = append(lines, [2]int{line, line})
	}
	//Lexers convert source text into a stream of tokens
	lexer := lexers.Get("go") //Recognizes the chroma package for go
	iterator, err := lexer.Tokenise(nil, b.String())
	if err != nil {
		log.Printf(err.Error()) //check error
	}
	//styles specify how token types are mapped to colours
	style := styles.Get("github") //takes a theme
	// formatters convert tokens and styles into formatted output.
	formatter := html.New(html.TabWidth(2), html.WithLineNumbers(), html.LineNumbersInTable(), html.HighlightLines(lines))
	w.Header().Set("Content-Type", "text/html")               //so that browser understand the HTML
	fmt.Fprint(w, "<style>pre { font-size: 1.2em; }</style>") //To incresase the font size
	formatter.Format(w, style, iterator)
}

// devMw is a HTTP middleware that recovers from any panics in our application
// and renders a stack trace.
func devMw(app http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() { // Defer is commonly used to simplify functions that perform various clean-up actions.
			if err := recover(); err != nil { //recover captures all the error and executes program normally after panic
				log.Println(err)       //recover only works in defer funcion
				stack := debug.Stack() //stack trace of the goroutine
				log.Println(string(stack))
				w.WriteHeader(http.StatusInternalServerError) //Provide response handler with status code
				fmt.Fprintf(w, "<h1/panic-after/>panic: %v</h1><pre>%s</pre>", err, makeLinks(string(stack)))
			}
		}()
		app.ServeHTTP(w, r)
	}
}

func panicDemo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>I am panicked,I will screw your code!</h1>")
	funcThatPanics()
}

func panicAfterDemo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Sucessfully Screwed your code!</h1>")
	funcThatPanics()
}

func funcThatPanics() {
	panic("Oh no!")
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello!")
}

// makeLinks makes links of stack traces and highlight the line having panic
func makeLinks(stack string) string {
	lines := strings.Split(stack, "\n") //\n is the seperator
	for li, line := range lines {
		if len(line) == 0 || line[0] != '\t' {
			continue
		}
		file := ""
		for i, ch := range line {
			if ch == ':' {
				file = line[1:i] //url of file till :, starting from space
				break
			}
		}
		var lineStr strings.Builder
		// start iterating after : to get the line no.
		for i := len(file) + 2; i < len(line); i++ {
			if line[i] < '0' || line[i] > '9' { //if range of char is out of range then brake
				break
			}
			lineStr.WriteByte(line[i])
		}
		v := url.Values{} //To Encode the path of the source file , v.Encode()
		v.Set("path", file)
		v.Set("line", lineStr.String())
		lines[li] = "\t<a href=\"/debug/?" + v.Encode() + "\">" + file + ":" + lineStr.String() + "</a>" + line[len(file)+2+len(lineStr.String()):]
	}
	return strings.Join(lines, "\n")
}
