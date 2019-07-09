package errors

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

//Generic Errors
var NotFound = errors.New("NOT_FOUND")
var GenericError = errors.New("GENERIC_ERROR")

//Contact Errors
var ContactNotFound = errors.New("CONTACT_NOT_FOUND")

//Event Errors
var EventNotFound = errors.New("EVENT_NOT_FOUND")
var LateEvent = errors.New("LATE_EVENT")
var NoTickets = errors.New("NO_TICKETS")

//Template Errors
var TemplateNotFound = errors.New("TEMPLATE_NOT_FOUND")

func ParamNotFound(param string) error {
	return errors.New(fmt.Sprintf("PARAM_%s_NOT_FOUND", strings.ToUpper(param)))
}
