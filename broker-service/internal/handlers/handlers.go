package handlers

import (
	"broker/internal/utils"
	"net/http"
)

func Broker(w http.ResponseWriter, r *http.Request) {
	payload := utils.JsonResponse{
		Error:   false,
		Message: "Hit the broker",
	}

	if err := utils.WriteJSON(w, http.StatusOK, payload); err != nil {
		return
	}
}
