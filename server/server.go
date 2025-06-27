package server

import (
	gogurt_proto "github.com/fozzbuzz/gogurt-rpc/proto"
)

type gogurtVendor struct{}

func CreateGogurtVendor() *gogurtVendor {
	s := &gogurtVendor{}

	return s
}

func (v *gogurtVendor) GetGogurt(req *gogurt_proto.GetGogurtRequest) (*gogurt_proto.Gogurt, error) {
	gogurt := gogurt_proto.Gogurt{
		Size:      120,
		Remaining: 120,
		Flavor:    "Cherry",
		Open:      false,
	}

	return &gogurt, nil
}
