# Convertir Audio a Texto usando IBM Cloud 

## Sobre el Servicio Speech To Text de IBM®

El servicio IBM® Speech to Text transcribe el audio al texto para habilitar las capacidades de transcripción de voz para las aplicaciones. Este tutorial basado en rizos puede ayudarle a comenzar rápidamente con el servicio. Los ejemplos muestran cómo llamar al método de reconocimiento POST /v1 del servicio para solicitar una transcripción.

## Primeros pasos
1. https://cloud.ibm.com/docs/speech-to-text?topic=speech-to-text-gettingStarted
2. https://cloud.ibm.com/apidocs/speech-to-text?code=go#introduction

## Uso

Agrega en la carpeta raiz un archivo `.env`  con las siguientes variables:

```shell
API_KEY=TU_API_KEY
SERVICE_URL=TU_URL_SERVICIO
```

Para ejecutar el programa se debe pasar el agrumento `file` en la ejecución con la ruta del archivo y su extensión
```shell
go run .\main.go -file="audioflac.flac"
```

### Opcional
Para pasar el modelo del lenguaje como argumento sli se debe utilizar model (por defecto está: *es-ES_BroadbandModel*)

```shell
go run .\main.go -file="audioflac.flac" -model="ar-AR_BroadbandModel" 
```

## Dependencias Principales
### Watson SDKs (Go)
#### https://cloud.ibm.com/docs/speech-to-text?topic=watson-using-sdks
