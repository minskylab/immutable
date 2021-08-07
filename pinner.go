package immutable

import (
	"encoding/json"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
)

type PinataResponse struct {
	IpfsHash  string
	Timestamp time.Time
}

func pinAssetToIPFS(cfg *Config, filepath string) (*PinataResponse, error) {
	// data, err :=os.ReadFile(filepath)
	// Create a Resty Client
	client := resty.New()
	client = client.SetHostURL("https://api.pinata.cloud")

	// "pinataMetadata"
	r, err := client.R().
		EnableTrace().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetHeader("pinata_api_key", cfg.Pinata.APIKey).
		SetHeader("pinata_secret_api_key", cfg.Pinata.APISecret).
		SetFile("file", filepath).
		// SetFormData(map[string]string{
		// 	"pinataMetadata": "{\"name\":\"immutable\"}",
		// }).
		Post("/pinning/pinFileToIPFS")
	if err != nil {
		return nil, errors.WithStack(err)
	}

	pinata := &PinataResponse{}

	if err := json.Unmarshal(r.Body(), pinata); err != nil {
		return nil, errors.WithStack(err)
	}

	return pinata, nil
}

func PinDocumentToIPFS(cfg *Config) (*PinataResponse, error) {
	return pinAssetToIPFS(cfg, finalResultPath(cfg))
}
