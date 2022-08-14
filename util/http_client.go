package util

import (
	"context"
	"io"
	"net/http"
)

//HttpCustomClient a custom http client,
func HttpCustomClient(ctx context.Context, request *http.Request) (response []byte, statusCode int, err error) {
	httpClient := http.Client{}
	httpResp, err := httpClient.Do(request.WithContext(ctx))
	if err != nil {
		return nil, httpResp.StatusCode, err
	}
	defer func(body io.ReadCloser) {
		_ = body.Close()
	}(httpResp.Body)
	response, err = io.ReadAll(httpResp.Body)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return response, httpResp.StatusCode, nil
}
