package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type OrderRequest struct {
	OrderID   string  `json:"orderId"`
	ProductID string  `json:"productId"`
	Quantity  int     `json:"quantity"`
	Amount    float64 `json:"amount"`
}

func createOrderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req OrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"orderId":   req.OrderID,
		"productId": req.ProductID,
		"quantity":  req.Quantity,
		"status":    "ORDER_CREATED_AND_QUEUED",
		"timestamp": time.Now().UTC().Format(time.RFC3339),
	})
}

func main() {
	http.HandleFunc("/api/v1/orders", createOrderHandler)
	_ = http.ListenAndServe(":8080", nil)
}
