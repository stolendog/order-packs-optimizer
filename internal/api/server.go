package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/stolendog/order-packs-optimizer/internal/app"
	"github.com/stolendog/order-packs-optimizer/internal/domain"
)

// TODO move to separate handlers and make server struct

func Start(app *app.App) error {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		http.ServeFile(w, r, "web/index.html")
	})

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	mux.HandleFunc("/api/packs", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handleGetPacks(app, w, r)
		case http.MethodPost:
			handlePostPacks(app, w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

	})

	mux.HandleFunc("/api/calculate", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		var req CalculatePacksRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Bad Request: "+err.Error(), http.StatusBadRequest)
			return
		}

		result, err := app.CalculatePacks(req.OrderQuantity)
		if err != nil {
			http.Error(w, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
			return
		}

		packsUsed := make([]PackUsage, 0, len(result.PacksUsed))
		for size, quantity := range result.PacksUsed {
			packsUsed = append(packsUsed, PackUsage{
				PackSize: int(size),
				Quantity: quantity,
			})
		}

		resp := CalculatePacksResponse{
			PacksUsed: packsUsed,
		}

		respondJSON(w, http.StatusOK, resp)
	})

	fmt.Println("Starting server on :9999")

	if err := http.ListenAndServe("0.0.0.0:9999", mux); err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}

	return nil
}

func respondJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "failed to encode JSON", http.StatusInternalServerError)
	}
}

func handleGetPacks(app *app.App, w http.ResponseWriter, r *http.Request) {
	packList, err := app.GetAllPacks()
	if err != nil {
		http.Error(w, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	packs := make([]PackInfo, 0, len(packList))
	for _, pack := range packList {
		packs = append(packs, PackInfo{
			Size: int(pack.Size),
		})
	}
	respondJSON(w, http.StatusOK, PackListResponse{Packs: packs})
}

func handlePostPacks(app *app.App, w http.ResponseWriter, r *http.Request) {
	var req PackListRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Bad Request: "+err.Error(), http.StatusBadRequest)
		return
	}

	var packs []domain.Pack
	for _, packSize := range req.Packs {
		pack, err := domain.NewPack(packSize)
		if err != nil {
			http.Error(w, "Bad Request: "+err.Error(), http.StatusBadRequest)
			return
		}
		packs = append(packs, pack)
	}

	if err := app.StorePackList(packs); err != nil {
		http.Error(w, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
