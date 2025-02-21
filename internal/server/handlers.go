package server

import (
	"encoding/json"
	"net/http"

	keygen "github.com/JulianoL13/passgen/internal/random"
)

type PasswordRequest struct {
	Size       int  `json:"size"`
	UseUpper   bool `json:"useUpper"`
	UseNum     bool `json:"useNum"`
	UseSpecial bool `json:"useSpecial"`
}

func generateRandomPasswordHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var request PasswordRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		request = PasswordRequest{
			Size:       8,
			UseUpper:   true,
			UseNum:     true,
			UseSpecial: true,
		}
	}

	if request.Size == 0 {
		request.Size = 8
	}

	if request.Size < 4 {
		http.Error(w, "Error: password too short", http.StatusBadRequest)
		return
	}

	password := keygen.GetRandomPass(request.Size, request.UseUpper, request.UseNum, request.UseSpecial)

	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"password": password}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}
