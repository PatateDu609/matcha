package public

import (
    "fmt"
    "io/ioutil"
    "net/http"
	"github.com/PatateDu609/matcha/utils/log"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
    log.Logger.Infof("File Upload Endpoint Hit")

    // Parse our multipart form, 10 << 20 specifies a maximum
    // upload of 10 MB files.
    r.ParseMultipartForm(10 << 20)
    // FormFile returns the first file for the given key `myFile`
    // it also returns the FileHeader so we can get the Filename,
    // the Header and the size of the file
    file, handler, err := r.FormFile("myFile")
    if err != nil {
        log.Logger.Infof("Error Retrieving the File")
        fmt.Println(err)
        return
    }
    defer file.Close()
    log.Logger.Infof("Uploaded File: %+v\n", handler.Filename)
    log.Logger.Infof("File Size: %+v\n", handler.Size)
    log.Logger.Infof("MIME Header: %+v\n", handler.Header)

    // Create a temporary file within our temp-images directory that follows
    // a particular naming pattern
    tempFile, err := ioutil.TempFile("upload", "upload-*.png")
    if err != nil {
        fmt.Println(err)
    }
    defer tempFile.Close()

    // read all of the contents of our uploaded file into a
    // byte array
    fileBytes, err := ioutil.ReadAll(file)
    if err != nil {
        fmt.Println(err)
    }
    // write this byte array to our temporary file
    tempFile.Write(fileBytes)
    // return that we have successfully uploaded our file!
    fmt.Fprintf(w, "Successfully Uploaded File\n")
}
