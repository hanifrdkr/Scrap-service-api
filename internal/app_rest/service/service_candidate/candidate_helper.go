package service_candidate

import (
	"fmt"
	"github.com/jlaffaye/ftp"
	"helicopter-hr/config"
	"io"
	"math/rand"
	"mime/multipart"
	"strings"
	"time"
)

const (
	ProvDIYogyakarta = "DI YOGYAKARTA"
	ProvSumatraUtara = "SUMATERA UTARA"
	ProvSumatraSelatan = "SUMATERA SELATAN"
)

func UploadFile(cfg *config.ConfigApp, fileName string, fileHeader *multipart.FileHeader) error {
	conn, err := ftp.Dial(cfg.FTP.HostName)
	if err != nil {
		return err
	}

	err = conn.Login(cfg.FTP.Username, cfg.FTP.Password)
	if err != nil {
		return err
	}

	// Open the file from multipart.FileHeader
	file, err := fileHeader.Open()
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	// Create a reader from the file
	reader, writer := io.Pipe()

	// Copy the file to the pipe writer in a goroutine
	go func() {
		defer writer.Close()
		_, err := io.Copy(writer, file)
		if err != nil {
			fmt.Errorf("failed to copy file: %v\n", err)
		}
	}()

	destinationFile := cfg.FTP.Directory + "/" + fileName
	// Store the file on the FTP server
	err = conn.Stor(destinationFile, reader)
	if err != nil {
		return fmt.Errorf("failed to upload file to FTP server: %v", err)
	}

	if err := conn.Quit(); err != nil {
		return fmt.Errorf("failed to close conn FTP server: %v", err)
	}

	return nil
}

func GenerateRandomString() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const keyLength = 10

	rand.Seed(time.Now().UnixNano())
	shortKey := make([]byte, keyLength)
	for i := range shortKey {
		shortKey[i] = charset[rand.Intn(len(charset))]
	}
	return string(shortKey)
}

func ConvertProvince(input string) string {
	input = strings.TrimLeft(strings.ToUpper(input), " ")

	if strings.ToUpper(input) == "DI YOGYAKARTA" || strings.ToUpper(input) == "DAERAH ISTIMEWA YOGYAKARTA" {
		input = ProvDIYogyakarta
	} else if strings.ToUpper(input) == "SUMATRA UTARA" {
		input = ProvSumatraUtara
	} else if strings.ToUpper(input) == "SUMATRA SELATAN" {
		input = ProvSumatraSelatan
	}

	return input
}

func ConvertRegency(regency string) string {
	if !strings.Contains(strings.ToUpper(regency), "KAB.") && !strings.Contains(strings.ToUpper(regency), "KABUPATEN") &&
		!strings.Contains(strings.ToUpper(regency), "KOTA") {
		return strings.ToUpper("KOTA " + regency)
	}

	return strings.ToUpper(strings.Replace(strings.ToUpper(regency), "KAB.", "KABUPATEN", -1))
}
