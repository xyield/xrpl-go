package client

import "github.com/xyield/xrpl-go/model/client/server"

type Server interface {
	Fee(*server.FeeRequest) (*server.FeeResponse, XRPLResponse, error)
	Manifest(*server.ManifestRequest) (*server.ManifestResponse, XRPLResponse, error)
	ServerInfo(*server.ServerInfoRequest) (*server.ServerInfoResponse, XRPLResponse, error)
	ServerState(*server.ServerStateRequest) (*server.ServerStateRequest, XRPLResponse, error)
}

type serverImpl struct {
	client Client
}

func (s *serverImpl) Fee(req *server.FeeRequest) (*server.FeeResponse, XRPLResponse, error) {
	res, err := s.client.SendRequest(req)
	if err != nil {
		return nil, nil, err
	}
	var fr server.FeeResponse
	err = res.GetResult(&fr)
	if err != nil {
		return nil, nil, err
	}
	return &fr, res, nil
}

func (s *serverImpl) Manifest(req *server.ManifestRequest) (*server.ManifestResponse, XRPLResponse, error) {
	res, err := s.client.SendRequest(req)
	if err != nil {
		return nil, nil, err
	}
	var mr server.ManifestResponse
	err = res.GetResult(&mr)
	if err != nil {
		return nil, nil, err
	}
	return &mr, res, nil
}

func (s *serverImpl) ServerInfo(req *server.ServerInfoRequest) (*server.ServerInfoResponse, XRPLResponse, error) {
	res, err := s.client.SendRequest(req)
	if err != nil {
		return nil, nil, err
	}
	var sir server.ServerInfoResponse
	err = res.GetResult(&sir)
	if err != nil {
		return nil, nil, err
	}
	return &sir, res, nil
}

func (s *serverImpl) ServerState(req *server.ServerStateRequest) (*server.ServerStateResponse, XRPLResponse, error) {
	res, err := s.client.SendRequest(req)
	if err != nil {
		return nil, nil, err
	}
	var ssr server.ServerStateResponse
	err = res.GetResult(&ssr)
	if err != nil {
		return nil, nil, err
	}
	return &ssr, res, nil
}
