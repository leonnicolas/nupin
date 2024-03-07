// This code is auto generated. DO NOT EDIT.

package account_user

import (
	"time"

	"github.com/go-openapi/runtime"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type InstrumentedClientService struct {
	ClientService
	cv *prometheus.CounterVec
	hv *prometheus.HistogramVec
}

func NewInstrumentedClientService(impl ClientService, r prometheus.Registerer) *InstrumentedClientService {
	hv := promauto.With(r).NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "nuki_http_client_requests_duration_seconds",
			Help:    "the total number of http requests to the Nuki API",
			Buckets: prometheus.DefBuckets,
		}, []string{"method"})

	cv := promauto.With(r).NewCounterVec(
		prometheus.CounterOpts{
			Name: "nuki_http_client_requests",
			Help: "the total number of http requests to the Nuki API",
		}, []string{"method", "result"})

	return &InstrumentedClientService{
		ClientService: impl,
		cv:            cv,
		hv:            hv,
	}
}

func (c *InstrumentedClientService) AccountUserResourceDeleteDelete(_c0 *AccountUserResourceDeleteDeleteParams, _c1 runtime.ClientAuthInfoWriter, _c2 ...ClientOption) (*AccountUserResourceDeleteDeleteNoContent, error) {
	t := time.Now()
	defer func() {
		c.hv.WithLabelValues("AccountUserResourceDeleteDelete").Observe(time.Since(t).Seconds())
	}()
	_a0, err := c.ClientService.AccountUserResourceDeleteDelete(_c0, _c1, _c2...)
	if err != nil {
		c.cv.WithLabelValues("AccountUserResourceDeleteDelete", "error").Inc()
		return _a0, err
	}
	c.cv.WithLabelValues("AccountUserResourceDeleteDelete", "success").Inc()

	return _a0, err
}

func (c *InstrumentedClientService) AccountUserResourceGetGet(_c0 *AccountUserResourceGetGetParams, _c1 runtime.ClientAuthInfoWriter, _c2 ...ClientOption) (*AccountUserResourceGetGetOK, error) {
	t := time.Now()
	defer func() {
		c.hv.WithLabelValues("AccountUserResourceGetGet").Observe(time.Since(t).Seconds())
	}()
	_a0, err := c.ClientService.AccountUserResourceGetGet(_c0, _c1, _c2...)
	if err != nil {
		c.cv.WithLabelValues("AccountUserResourceGetGet", "error").Inc()
		return _a0, err
	}
	c.cv.WithLabelValues("AccountUserResourceGetGet", "success").Inc()

	return _a0, err
}

func (c *InstrumentedClientService) AccountUserResourcePostPost(_c0 *AccountUserResourcePostPostParams, _c1 runtime.ClientAuthInfoWriter, _c2 ...ClientOption) (*AccountUserResourcePostPostOK, *AccountUserResourcePostPostNoContent, error) {
	t := time.Now()
	defer func() {
		c.hv.WithLabelValues("AccountUserResourcePostPost").Observe(time.Since(t).Seconds())
	}()
	_a0, _a1, err := c.ClientService.AccountUserResourcePostPost(_c0, _c1, _c2...)
	if err != nil {
		c.cv.WithLabelValues("AccountUserResourcePostPost", "error").Inc()
		return _a0, _a1, err
	}
	c.cv.WithLabelValues("AccountUserResourcePostPost", "success").Inc()

	return _a0, _a1, err
}

func (c *InstrumentedClientService) AccountUsersResourceGetGet(_c0 *AccountUsersResourceGetGetParams, _c1 runtime.ClientAuthInfoWriter, _c2 ...ClientOption) (*AccountUsersResourceGetGetOK, error) {
	t := time.Now()
	defer func() {
		c.hv.WithLabelValues("AccountUsersResourceGetGet").Observe(time.Since(t).Seconds())
	}()
	_a0, err := c.ClientService.AccountUsersResourceGetGet(_c0, _c1, _c2...)
	if err != nil {
		c.cv.WithLabelValues("AccountUsersResourceGetGet", "error").Inc()
		return _a0, err
	}
	c.cv.WithLabelValues("AccountUsersResourceGetGet", "success").Inc()

	return _a0, err
}

func (c *InstrumentedClientService) AccountUsersResourcePutPut(_c0 *AccountUsersResourcePutPutParams, _c1 runtime.ClientAuthInfoWriter, _c2 ...ClientOption) (*AccountUsersResourcePutPutOK, error) {
	t := time.Now()
	defer func() {
		c.hv.WithLabelValues("AccountUsersResourcePutPut").Observe(time.Since(t).Seconds())
	}()
	_a0, err := c.ClientService.AccountUsersResourcePutPut(_c0, _c1, _c2...)
	if err != nil {
		c.cv.WithLabelValues("AccountUsersResourcePutPut", "error").Inc()
		return _a0, err
	}
	c.cv.WithLabelValues("AccountUsersResourcePutPut", "success").Inc()

	return _a0, err
}
