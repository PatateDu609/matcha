package public

import (
    "os"
    "fmt"
    "io/ioutil"
    "net/http"
	"github.com/PatateDu609/matcha/utils/log"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
    log.Logger.Infof("File Upload Endpoint Hit")

    // maximum upload of 10 MB files.
    r.ParseMultipartForm(10 << 20)
    file, handler, err := r.FormFile("myFile")
    num := r.FormValue("number")
    user := r.FormValue("user")
    if err != nil {
        log.Logger.Infof("Error Retrieving the File")
        fmt.Println(err)
        return
    }
    defer file.Close()
    log.Logger.Infof("Uploaded File: %+v\n", handler.Filename)
    log.Logger.Infof("File Size: %+v\n", handler.Size)
    log.Logger.Infof("MIME Header: %+v\n", handler.Header)

    if err := os.Mkdir("upload/" + user, os.ModePerm); err != nil {
        fmt.Println(err)
    }

    tempFile, err := os.CreateTemp("upload/" + user, "image-*.png")
    if err != nil {
        fmt.Println(err)
    }
    defer tempFile.Close()

    fileBytes, err := ioutil.ReadAll(file)
    if err != nil {
        fmt.Println(err)
    }
    tempFile.Write(fileBytes)
    fmt.Fprintf(w, "Successfully Uploaded File\n")

    thepath:= "/upload/" + user + tempFile.Name()
    setImage(thepath, num, user, r)
}
