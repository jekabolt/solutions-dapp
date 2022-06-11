package app

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/jekabolt/solutions-dapp/art-admin/store/nft"
)

// errors

type StatusResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText string `json:"status"`          // user-level status message
	AppCode    int64  `json:"code,omitempty"`  // application-specific error code
	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
}

func (e *StatusResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func ErrInvalidRequest(err error) render.Renderer {
	return &StatusResponse{
		Err:            err,
		HTTPStatusCode: http.StatusBadRequest,
		StatusText:     "Invalid request.",
		ErrorText:      err.Error(),
	}
}

func ErrInternalServerRequest(err error) render.Renderer {
	return &StatusResponse{
		Err:            err,
		HTTPStatusCode: http.StatusInternalServerError,
		StatusText:     http.StatusText(http.StatusInternalServerError),
		ErrorText:      err.Error(),
	}
}

func ErrRender(err error) render.Renderer {
	return &StatusResponse{
		Err:            err,
		HTTPStatusCode: http.StatusUnprocessableEntity,
		StatusText:     "Error rendering response.",
		ErrorText:      err.Error(),
	}
}

func ErrInternalServerError(err error) render.Renderer {
	return &StatusResponse{
		Err:            err,
		HTTPStatusCode: http.StatusInternalServerError,
		StatusText:     http.StatusText(http.StatusInternalServerError),
		ErrorText:      err.Error(),
	}
}

func ErrUnauthorizedError(err error) render.Renderer {
	return &StatusResponse{
		Err:            err,
		HTTPStatusCode: http.StatusUnauthorized,
		StatusText:     http.StatusText(http.StatusUnauthorized),
		ErrorText:      err.Error(),
	}
}

var ErrNotFound = &StatusResponse{HTTPStatusCode: http.StatusNotFound, StatusText: http.StatusText(http.StatusNotFound)}
var StatusOK = &StatusResponse{HTTPStatusCode: http.StatusOK, StatusText: http.StatusText(http.StatusOK)}

// auth

type AuthResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

func NewAuthResponse(ar *AuthResponse) *AuthResponse {
	return ar
}

func (i *AuthResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// nft mints request

type NFTMintResponse struct {
	NFTMintRequest *nft.NFTMintRequest `json:"mint"`
}

func NewNFTMintResponse(mr nft.NFTMintRequest) *NFTMintResponse {
	return &NFTMintResponse{NFTMintRequest: &mr}
}

func (mr *NFTMintResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return mr.NFTMintRequest.Validate()
}

func NewNFTMintRequestListResponse(nftMintRequests []nft.NFTMintRequest) []render.Renderer {
	list := []render.Renderer{}
	for _, nftMr := range nftMintRequests {
		list = append(list, NewNFTMintResponse(nftMr))
	}
	return list
}
