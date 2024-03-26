package address

import (
	"net/http"

	"github.com/claudineyveloso/bookernet.git/internal/types"
	"github.com/claudineyveloso/bookernet.git/internal/utils"
)

type Handler struct {
	addressStore types.AddressStore
}

func NewHandler(addressStore types.AddressStore) *Handler {
	return &Handler{addressStore: addressStore}
}

func (h *Handler) handleCreateAddress(w http.ResponseWriter, r *http.Request) {
	var address types.CreateAddressPayload
	if err := utils.ParseJSON(r, &address); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	err := h.addressStore.CreateAddress(address)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusCreated, address)
}
