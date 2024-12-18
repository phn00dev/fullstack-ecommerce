package images

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"math/rand"
	"os"
	"strings"
)

func RandStringRunes(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func GetLastXChars(s string, length int) (ext string) {
	ext = string(s[len(s)-length:])
	return
}

func CheckExtension(ext string) bool {
	extensions := []string{"png", "webm", "jpg", "jpeg"}
	for _, v := range extensions {
		if v == ext {
			return true
		}
	}
	return false
}

func UploadFile(ctx *fiber.Ctx, inputName, publicPath, ImagePath string) (*string, error) {
	file, err := ctx.FormFile(inputName)
	if err != nil {
		return nil, err
	}

	fileExt := GetLastXChars(file.Filename, 3)
	if fileExt == "peg" {
		fileExt = "j" + fileExt
	}

	// file Ext validate -->(png , jpg, webm,peg)
	isSupported := CheckExtension(fileExt)
	if !isSupported {
		return nil, errors.New("file formady nädogry")
	}
	fullPath := fmt.Sprintf("%s/%s", publicPath, ImagePath)
	err = os.MkdirAll(fullPath, 0755)
	if err != nil {
		return nil, err
	}

	uniqueId := uuid.New()
	filename := fmt.Sprintf("%s.%s", strings.Replace(uniqueId.String(), "-", "", -1), fileExt)
	filePath := fmt.Sprintf("%s/%s", fullPath, filename)
	err = ctx.SaveFile(file, filePath)
	if err != nil {
		return nil, err
	}
	dbPath := fmt.Sprintf("public/%s/%s", ImagePath, filename)
	return &dbPath, nil
}
