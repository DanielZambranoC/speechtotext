package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/gabriel-vasile/mimetype"
	"github.com/joho/godotenv"
	"github.com/watson-developer-cloud/go-sdk/speechtotextv1"
)

var archivo = flag.String("file", "", "Ruta del archivo de audio")
var modelo = flag.String("model", "es-ES_BroadbandModel", "Modelo del idioma, defecto: es-ES_BroadbandModel")

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	flag.Parse()

	authenticator := &core.IamAuthenticator{
		ApiKey: os.Getenv("API_KEY"),
	}

	options := &speechtotextv1.SpeechToTextV1Options{
		Authenticator: authenticator,
	}

	speechToText, speechToTextErr := speechtotextv1.NewSpeechToTextV1(options)

	if speechToTextErr != nil {
		panic(speechToTextErr)
	}

	speechToText.SetServiceURL(os.Getenv("SERVICE_URL"))

	model := *modelo

	files := [1]string{*archivo}
	for _, file := range files {

		mime, err := mimetype.DetectFile(file)
		if err != nil {
			panic(err)
		}

		var audioFile io.ReadCloser
		var audioFileErr error
		audioFile, audioFileErr = os.Open(file)
		if audioFileErr != nil {
			panic(audioFileErr)
		}
		result, _, responseErr := speechToText.Recognize(
			&speechtotextv1.RecognizeOptions{
				Audio:                     audioFile,
				ContentType:               core.StringPtr(mime.String()),
				Timestamps:                core.BoolPtr(true),
				WordAlternativesThreshold: core.Float32Ptr(0.9),
				Model:                     &model,
			},
		)
		if responseErr != nil {
			panic(responseErr)
		}

		x := result.Results
		for _, a := range x {
			for _, t := range a.Alternatives {
				fmt.Println(*t.Transcript)
			}
		}

	}
}
