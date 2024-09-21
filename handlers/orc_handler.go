package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"ocr-pdf-api/models"
	"os"
	"os/exec"
	"path/filepath"
)

// OCRHandler handles the OCR request
func OCRHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Converting file.")

	langStringUrl := r.URL.Query().Get("lang")
	lang, err := models.ParseLangString(langStringUrl)

	if err != nil {
		lang = models.EN
	}

	err = r.ParseMultipartForm(10 << 20) // Limit upload size to 10MB
	if err != nil {
		http.Error(w, "Unable to parse multipart form object", http.StatusBadRequest)
		return
	}

	file, handler, err := r.FormFile("pdf")
	if err != nil {
		http.Error(w, "Error loading the provided pdf file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	tempDir := os.TempDir()
	uploadedFilePath := filepath.Join(tempDir, handler.Filename)
	tempFile, err := os.Create(uploadedFilePath)
	if err != nil {
		http.Error(w, "Could not create temporary file", http.StatusInternalServerError)
		return
	}
	defer tempFile.Close()

	_, err = io.Copy(tempFile, file)
	if err != nil {
		http.Error(w, "Error writing the file", http.StatusInternalServerError)
		return
	}

	ocrFilepath := filepath.Join(tempDir, "ocr_"+handler.Filename)

	cmd := exec.Command("ocrmypdf", "--force-ocr", "-l", lang.String(), uploadedFilePath, ocrFilepath)

	output, err := cmd.CombinedOutput() // Capture both stdout and stderr
	if err != nil {
		log.Printf("Error while running the OCR command: %v", err)
		log.Printf("OCR command output: %s", string(output))
		http.Error(w, "Error processing the PDF with OCR", http.StatusInternalServerError)
		return
	}

	ocrFile, err := os.Open(ocrFilepath)
	if err != nil {
		http.Error(w, "Could not open OCR-processed PDF", http.StatusInternalServerError)
		return
	}
	defer ocrFile.Close()

	// Check if the file is empty
	fileInfo, err := ocrFile.Stat()
	if err != nil || fileInfo.Size() == 0 {
		http.Error(w, "Generated file is empty", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"ocr_%s\"", handler.Filename))

	_, err = io.Copy(w, ocrFile)
	if err != nil {
		http.Error(w, "Error sending the file", http.StatusInternalServerError)
		return
	}

	log.Println("Conversion completed.")
}
