package types

import "errors"

func (c *ConvReq) Validate() error {
	if c == nil {
		return errors.New("conv req is nil")
	}

	return c.BaseParams.Validate()
}

func (s *BaseConv) Validate() error {
	if s == nil {
		return errors.New("base params is nil")
	}
	// pid
	if s.PID == "" {
		return errors.New("pid is empty")
	}
	// channel
	if s.Channel == "" {
		return errors.New("channel is empty")
	}
	// adid
	if s.AdID == "" {
		return errors.New("adid is empty")
	}
	// brand
	if s.Brand == "" {
		return errors.New("brand is empty")
	}

	return nil
}
