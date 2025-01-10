package router

import (
	"net/http"

	timer "github.com/PauloHPMKT/cronometro/controllers"
)

func Router() {
	http.HandleFunc("/timer", timer.HandleTimer)
}
