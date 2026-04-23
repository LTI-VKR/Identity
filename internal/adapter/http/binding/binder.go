package binding

import (
	"encoding/json"
	"fmt"
	httpError "identity/internal/adapter/http/errors"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func DecodeAndValidate[T any](r *http.Request, dst *T) error {
	decode := json.NewDecoder(r.Body)
	decode.DisallowUnknownFields()

	if err := decode.Decode(dst); err != nil {
		return fmt.Errorf(`%w:`+err.Error(), httpError.ErrInvalidJSON)
	}

	return validate.Struct(dst)
}
