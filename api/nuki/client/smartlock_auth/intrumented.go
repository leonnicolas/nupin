// This code is auto generated. DO NOT EDIT.

package smartlock_auth

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

func (c *InstrumentedClientService) SmartlockAuthResourceDeleteDelete(_c0 *SmartlockAuthResourceDeleteDeleteParams, _c1 runtime.ClientAuthInfoWriter, _c2 ...ClientOption) (*SmartlockAuthResourceDeleteDeleteNoContent, error) {
	t := time.Now()
	defer func() {
		c.hv.WithLabelValues("SmartlockAuthResourceDeleteDelete").Observe(time.Since(t).Seconds())
	}()
	_a0, err := c.ClientService.SmartlockAuthResourceDeleteDelete(_c0, _c1, _c2...)
	if err != nil {
		c.cv.WithLabelValues("SmartlockAuthResourceDeleteDelete", "error").Inc()
		return _a0, err
	}
	c.cv.WithLabelValues("SmartlockAuthResourceDeleteDelete", "success").Inc()

	return _a0, err
}

func (c *InstrumentedClientService) SmartlockAuthResourceGetGet(_c0 *SmartlockAuthResourceGetGetParams, _c1 runtime.ClientAuthInfoWriter, _c2 ...ClientOption) (*SmartlockAuthResourceGetGetOK, error) {
	t := time.Now()
	defer func() {
		c.hv.WithLabelValues("SmartlockAuthResourceGetGet").Observe(time.Since(t).Seconds())
	}()
	_a0, err := c.ClientService.SmartlockAuthResourceGetGet(_c0, _c1, _c2...)
	if err != nil {
		c.cv.WithLabelValues("SmartlockAuthResourceGetGet", "error").Inc()
		return _a0, err
	}
	c.cv.WithLabelValues("SmartlockAuthResourceGetGet", "success").Inc()

	return _a0, err
}

func (c *InstrumentedClientService) SmartlockAuthResourcePostPost(_c0 *SmartlockAuthResourcePostPostParams, _c1 runtime.ClientAuthInfoWriter, _c2 ...ClientOption) (*SmartlockAuthResourcePostPostNoContent, error) {
	t := time.Now()
	defer func() {
		c.hv.WithLabelValues("SmartlockAuthResourcePostPost").Observe(time.Since(t).Seconds())
	}()
	_a0, err := c.ClientService.SmartlockAuthResourcePostPost(_c0, _c1, _c2...)
	if err != nil {
		c.cv.WithLabelValues("SmartlockAuthResourcePostPost", "error").Inc()
		return _a0, err
	}
	c.cv.WithLabelValues("SmartlockAuthResourcePostPost", "success").Inc()

	return _a0, err
}

func (c *InstrumentedClientService) SmartlockAuthsResourceGetGet(_c0 *SmartlockAuthsResourceGetGetParams, _c1 runtime.ClientAuthInfoWriter, _c2 ...ClientOption) (*SmartlockAuthsResourceGetGetOK, error) {
	t := time.Now()
	defer func() {
		c.hv.WithLabelValues("SmartlockAuthsResourceGetGet").Observe(time.Since(t).Seconds())
	}()
	_a0, err := c.ClientService.SmartlockAuthsResourceGetGet(_c0, _c1, _c2...)
	if err != nil {
		c.cv.WithLabelValues("SmartlockAuthsResourceGetGet", "error").Inc()
		return _a0, err
	}
	c.cv.WithLabelValues("SmartlockAuthsResourceGetGet", "success").Inc()

	return _a0, err
}

func (c *InstrumentedClientService) SmartlockAuthsResourcePutPut(_c0 *SmartlockAuthsResourcePutPutParams, _c1 runtime.ClientAuthInfoWriter, _c2 ...ClientOption) (*SmartlockAuthsResourcePutPutNoContent, error) {
	t := time.Now()
	defer func() {
		c.hv.WithLabelValues("SmartlockAuthsResourcePutPut").Observe(time.Since(t).Seconds())
	}()
	_a0, err := c.ClientService.SmartlockAuthsResourcePutPut(_c0, _c1, _c2...)
	if err != nil {
		c.cv.WithLabelValues("SmartlockAuthsResourcePutPut", "error").Inc()
		return _a0, err
	}
	c.cv.WithLabelValues("SmartlockAuthsResourcePutPut", "success").Inc()

	return _a0, err
}

func (c *InstrumentedClientService) SmartlocksAuthsResourceDeleteDelete(_c0 *SmartlocksAuthsResourceDeleteDeleteParams, _c1 runtime.ClientAuthInfoWriter, _c2 ...ClientOption) (*SmartlocksAuthsResourceDeleteDeleteNoContent, error) {
	t := time.Now()
	defer func() {
		c.hv.WithLabelValues("SmartlocksAuthsResourceDeleteDelete").Observe(time.Since(t).Seconds())
	}()
	_a0, err := c.ClientService.SmartlocksAuthsResourceDeleteDelete(_c0, _c1, _c2...)
	if err != nil {
		c.cv.WithLabelValues("SmartlocksAuthsResourceDeleteDelete", "error").Inc()
		return _a0, err
	}
	c.cv.WithLabelValues("SmartlocksAuthsResourceDeleteDelete", "success").Inc()

	return _a0, err
}

func (c *InstrumentedClientService) SmartlocksAuthsResourceGetGet(_c0 *SmartlocksAuthsResourceGetGetParams, _c1 runtime.ClientAuthInfoWriter, _c2 ...ClientOption) (*SmartlocksAuthsResourceGetGetOK, error) {
	t := time.Now()
	defer func() {
		c.hv.WithLabelValues("SmartlocksAuthsResourceGetGet").Observe(time.Since(t).Seconds())
	}()
	_a0, err := c.ClientService.SmartlocksAuthsResourceGetGet(_c0, _c1, _c2...)
	if err != nil {
		c.cv.WithLabelValues("SmartlocksAuthsResourceGetGet", "error").Inc()
		return _a0, err
	}
	c.cv.WithLabelValues("SmartlocksAuthsResourceGetGet", "success").Inc()

	return _a0, err
}

func (c *InstrumentedClientService) SmartlocksAuthsResourcePostPost(_c0 *SmartlocksAuthsResourcePostPostParams, _c1 runtime.ClientAuthInfoWriter, _c2 ...ClientOption) (*SmartlocksAuthsResourcePostPostNoContent, error) {
	t := time.Now()
	defer func() {
		c.hv.WithLabelValues("SmartlocksAuthsResourcePostPost").Observe(time.Since(t).Seconds())
	}()
	_a0, err := c.ClientService.SmartlocksAuthsResourcePostPost(_c0, _c1, _c2...)
	if err != nil {
		c.cv.WithLabelValues("SmartlocksAuthsResourcePostPost", "error").Inc()
		return _a0, err
	}
	c.cv.WithLabelValues("SmartlocksAuthsResourcePostPost", "success").Inc()

	return _a0, err
}

func (c *InstrumentedClientService) SmartlocksAuthsResourcePutPut(_c0 *SmartlocksAuthsResourcePutPutParams, _c1 runtime.ClientAuthInfoWriter, _c2 ...ClientOption) (*SmartlocksAuthsResourcePutPutNoContent, error) {
	t := time.Now()
	defer func() {
		c.hv.WithLabelValues("SmartlocksAuthsResourcePutPut").Observe(time.Since(t).Seconds())
	}()
	_a0, err := c.ClientService.SmartlocksAuthsResourcePutPut(_c0, _c1, _c2...)
	if err != nil {
		c.cv.WithLabelValues("SmartlocksAuthsResourcePutPut", "error").Inc()
		return _a0, err
	}
	c.cv.WithLabelValues("SmartlocksAuthsResourcePutPut", "success").Inc()

	return _a0, err
}
