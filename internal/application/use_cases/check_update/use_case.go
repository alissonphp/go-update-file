package check_update

import (
	"encoding/json"
	"fmt"
	"github.com/alissonphp/go-update-file/internal/domain"
	"io"
	"net/http"
	"os"
	"time"
)

type checkUpdateService struct{}

func CheckApi() error {
	resp, err := http.Get("https://pokeapi.co/api/v2/pokemon-form/132")
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	var result domain.Result
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return err
	}

	imageResp, err := http.Get(result.Sprites.FrontDefault)
	if err != nil {
		return err
	}

	defer imageResp.Body.Close()

	fileName := fmt.Sprintf("static/pos_images/pos_image_%v.png", time.Now().Unix())

	file, err := os.Create(fileName)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = io.Copy(file, imageResp.Body)
	if err != nil {
		return err
	}

	println("Imagem baixada com sucesso.")

	return nil

}
