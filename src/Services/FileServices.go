package services

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
)

// TODO
//Guardar archivo con nombre y la version
//Generar sistema de ERRORES para el cliente
//Generar sistema de actualizacion de archivos si la version es diferente
//Generar sistema de almacenamiento db de versiones (Estilo scrapper api)

func GetNewVersionOfServer() {
	_, err := os.Stat("Files")
	if os.IsNotExist(err) {
		log.Print("Files folder does not exist, creating it...")
		_CreateFilesFolder()
	}
	_, err = os.Stat("Files/BeamMP_Server.zip")
	if os.IsExist(err) {
		log.Print("File already exists, deleting it...")
		_DeleteFilesFolder()
		return
	}
	out, err := os.Create("Files/BeamMP_Server.zip")
	if err != nil {
		panic(err)
	}
	defer out.Close()
	archive, err := http.Get("https://beammp.com/server/BeamMP_Server.zip")
	if err != nil {
		panic(err)
	}
	defer archive.Body.Close()
	if archive.StatusCode != http.StatusOK {
		fmt.Println("ERROR: ", archive.Status)
		return
	}
	log.Print("File downloaded successfully!")
	io.Copy(out, archive.Body)
}

func SendBeamMPServerFilesService(w http.ResponseWriter, r *http.Request) {
	_, err := os.Stat("Files/BeamMP_Server.zip")
	if os.IsNotExist(err) {
		log.Print("File does not exist, downloading it...")
		GetNewVersionOfServer()
	}
	http.ServeFile(w, r, "Files/BeamMP_Server.zip")
}

func _CreateFilesFolder() {
	cmd := exec.Command("mkdir", "Files")
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

func _DeleteFilesFolder() {
	cmd := exec.Command("rm", "-rf", "Files")
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}
