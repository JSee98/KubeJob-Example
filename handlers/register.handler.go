package handlers

import (
	"net/http"
)
func RegisterHandlers() {
    http.HandleFunc("/kubejob", HandleKubeJobPost)
}