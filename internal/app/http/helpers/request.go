package helpers

import (
	"bytes"
	"encoding/json"
	"errors"
	internal_errors "go_notifier/internal/app/http/errors"
	"go_notifier/internal/common"
	"io"
	"net/http"
)

func CreateAndValidateFromRequest(r *http.Request, request common.Validatable) error {
	var buf bytes.Buffer
	_, err := io.Copy(&buf, r.Body)
	if err != nil {
		return internal_errors.Wrap(err, "unable to transform request to buffers byte", internal_errors.RequestErrCopyToBuffer)
	}

	requestData := buf.Bytes()
	r.Body.Close()

	if err := json.Unmarshal(requestData, &request); err != nil {
		return internal_errors.Wrap(err, "unable unmarshal request data", internal_errors.RequestErrUnmarshal)
	}

	err = request.Validate()
	if err != nil {
		return internal_errors.Wrap(err, "request data is not valid", internal_errors.RequestErrNotValid)
	}

	return nil
}

func HandleRequestError(err error, w http.ResponseWriter) {
	var wrapedErr *internal_errors.WrappedError
	if errors.As(err, &wrapedErr) {
		if wrapedErr.ErrorType == internal_errors.RequestErrNotValid {
			http.Error(w, "request is not valid", http.StatusBadRequest)
			return
		}
	}
	http.Error(w, "internal server error", http.StatusInternalServerError)
	return
}
