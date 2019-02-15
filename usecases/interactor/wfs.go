package interactor

import (
	"time"

	"github.com/mao-wfs/maoctrl/domain"
	"github.com/mao-wfs/maoctrl/usecases/ports"
)

// WFSInteractor is the interactor for MAO controller.
type WFSInteractor struct {
	Handler domain.WFSHandler
}

// NewWFSInteractor creates the new interactor for the WFS usecase.
func NewWFSInteractor(handler domain.WFSHandler) *WFSInteractor {
	return &WFSInteractor{
		Handler: handler,
	}
}

// Initialize initialize MAO wavefront sensor.
func (i *WFSInteractor) Initialize(
	integTime ports.IntegrationTime,
	winFunc ports.WindowFunction,
	swOrder ports.SwitchOrder,
) error {
	wfsConf, err := i.ConfigPort.FormatConfig(integTime, winFunc, swOrder)
	if err != nil {
		return err
	}
	return i.Handler.Initialize(wfsConf)
}

// Finalize finalize MAO wavefront sensor.
func (i *WFSInteractor) Finalize() error {
	return i.Handler.Finalize()
}

// Start starts MAO wavefront sensor.
func (i *WFSInteractor) Start(beginTime time.Time) error {
	return i.Handler.Start(beginTime)
}

// Halt stops MAO wavefront sensor.
func (i *WFSInteractor) Halt() error {
	return i.Handler.Halt()
}
