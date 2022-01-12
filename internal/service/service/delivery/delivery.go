package delivery

import (
	"github.com/just4n4cc/tp-sem2-db/internal/response"
	"github.com/just4n4cc/tp-sem2-db/internal/service/service"
	"github.com/just4n4cc/tp-sem2-db/pkg/logger"
	"net/http"
)

const logMessage = "service:service:delivery:"

type Delivery struct {
	useCase service.UseCase
}

func NewDelivery(useCase service.UseCase) *Delivery {
	return &Delivery{
		useCase: useCase,
	}
}

func (h *Delivery) ServiceClear(w http.ResponseWriter, r *http.Request) {
	message := logMessage + "ServiceClear:"
	logger.Debug(message + "started")

	err := h.useCase.ServiceClear()
	if err != nil {
		response.UnknownError(&w, err, message)
		return
	}

	w.WriteHeader(http.StatusOK)
	response.SetBody(&w, nil, message)
}

func (h *Delivery) ServiceStatus(w http.ResponseWriter, r *http.Request) {
	message := logMessage + "ServiceStatus:"
	logger.Debug(message + "started")

	s, err := h.useCase.ServiceStatus()
	if err != nil {
		response.UnknownError(&w, err, message)
		return
	}

	w.WriteHeader(http.StatusOK)
	response.SetBody(&w, s, message)
}