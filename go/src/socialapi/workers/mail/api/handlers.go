package api

import (
	"net/http"
	"net/url"
	"socialapi/workers/common/response"
	"socialapi/workers/mail/models"
)

func Parse(u *url.URL, h http.Header, req *models.Mail) (int, http.Header, interface{}, error) {
	if err := req.Validate(); err != nil {
		// faily silently, we dont want mail parser service to retry on
		// the failed validation
		return response.NewDefaultOK()
	}

	if err := req.Persist(); err != nil {
		return response.NewBadRequest(err)
	}

	return response.NewDefaultOK()
}
