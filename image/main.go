package main

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/shubhanshu7/Gophercises/image/primitive"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	html := ` 
	<html>
    <body>
        <form action="/upload" method="post" enctype="multipart/form-data">
            <div class="file-uploader__message-area">
                <p>Select a file to upload</p>
            </div>
            <div class="file-chooser">
                <input class="file-chooser__input" type="file" name="image" id="image">
            </div>
                <br><input class="file-uploader__submit-button" type="submit" value="Upload">
        </form>
    </body>
</html>`
	fmt.Fprint(w, html)

}

//uploadHandler to upload image that is persisted throughout one request
func uploadHandler(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("image")
	if err == nil {
		defer file.Close()
		ext := filepath.Ext(header.Filename)[1:]
		saveFile, err := tempfile("", ext)
		if err == nil {
			io.Copy(saveFile, file)
			http.Redirect(w, r, "modify/"+filepath.Base(saveFile.Name()), http.StatusFound)
		}
	}
	errorResponse(w, err)
}

//modifyHandler is used for the show all the transformed image based on input
func modifyHandler(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("./imgs/" + filepath.Base(r.URL.Path))
	if err == nil {
		ext := filepath.Ext(file.Name())[1:]
		mode := r.FormValue("mode")
		if mode == "" {
			generateAllMode(w, r, file, ext)
			return
		}
		number := r.FormValue("number")
		if number == "" {
			gegenerateSingleMode(w, r, file, ext, mode)
			return
		}
		http.Redirect(w, r, "/imgs/"+filepath.Base(file.Name()), http.StatusFound)
	}
	errorResponse(w, err)
}

//gegenerateSingleMode will create images using one mode but different number of shap shapes
func gegenerateSingleMode(w http.ResponseWriter, r *http.Request, file io.ReadSeeker, ext, mode string) {
	var a, b, c string
	var err error
	a, err = genrateImage(file, ext, mode, "50")
	if err == nil {
		file.Seek(0, 0)
		b, err = genrateImage(file, ext, mode, "100")
		if err == nil {
			file.Seek(0, 0)
			c, err = genrateImage(file, ext, mode, "150")
			if err == nil {
				html := `<html><body>
						{{range .}}
							<a href="/modify/{{.Name}}?mode={{.Mode}}&number={{.Number}}">
							<img style="width: 20%;" src="/{{.Name}}">
							</a>
						{{end}}
						</body></html>`
				tpl := template.Must(template.New("").Parse(html))
				type Images struct {
					Name   string
					Mode   int
					Number int
				}
				images := []Images{
					{a, 2, 50}, {b, 2, 100}, {c, 2, 150},
				}

				tpl.Execute(w, images)
			}
		}
	}

}

//generateAllMode will create images with different modes
func generateAllMode(w http.ResponseWriter, r *http.Request, file io.ReadSeeker, ext string) {
	var a, b, c, d string
	var err error
	a, err = genrateImage(file, ext, "2", "30")
	if err == nil {
		file.Seek(0, 0)
		b, err = genrateImage(file, ext, "3", "30")
		if err == nil {
			file.Seek(0, 0)
			c, err = genrateImage(file, ext, "4", "30")
			if err == nil {
				file.Seek(0, 0)
				d, err = genrateImage(file, ext, "5", "30")
				if err == nil {
					fmt.Print("Hello")
					file.Seek(0, 0)
					html := `<html><body>
					{{range .}}
						<a href="/modify/{{.Name}}?mode={{.Mode}}">
						<img style="width: 20%;" src="/{{.Name}}">
						</a>
					{{end}}
					</body></html>`
					tpl := template.Must(template.New("").Parse(html))
					type Images struct {
						Name string
						Mode int
					}
					images := []Images{
						{a, 2}, {b, 3}, {c, 4}, {d, 5},
					}

					tpl.Execute(w, images)
				}
			}
		}
	}

	errorResponse(w, err)
}

//genrateImage will generate new image using Tranform from primitive package
func genrateImage(file io.Reader, ext, mode, number string) (string, error) {
	var out io.Reader
	var err error
	var fileName string
	out, err = primitive.Transform(file, ext, mode, number)
	if err == nil {
		var outFile *os.File
		outFile, err = tempfile("", ext)
		if err == nil {
			io.Copy(outFile, out)
			fileName = outFile.Name()
		}
	}

	return fileName, nil
}

func errorResponse(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

//tempFile will create temporary file in same directory
func tempfile(prefix, ext string) (*os.File, error) {
	var in, out *os.File
	var err error
	in, err = ioutil.TempFile("./imgs/", prefix)
	if err == nil {
		defer os.Remove(in.Name())
		out, err = os.Create(fmt.Sprintf("%s.%s", in.Name(), ext))
	}
	return out, err
}

// getHandlers will return the router mux with handlers
func getHandlers() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/upload", uploadHandler)
	fileServer := http.FileServer(http.Dir("./imgs/"))
	mux.Handle("/imgs/", http.StripPrefix("/imgs", fileServer))
	mux.HandleFunc("/modify/", modifyHandler)
	return mux
}

var listenAndServerFunc = http.ListenAndServe

func main() {
	go func() {
		t := time.NewTicker(5 * time.Minute)
		for {
			<-t.C
		}
	}()
	log.Fatal(listenAndServerFunc(":3020", getHandlers()))
}
