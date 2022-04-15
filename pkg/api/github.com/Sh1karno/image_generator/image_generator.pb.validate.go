// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/proto/image_generator.proto

package image_generator

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on GetPlaceholderRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetPlaceholderRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetPlaceholderRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetPlaceholderRequestMultiError, or nil if none found.
func (m *GetPlaceholderRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetPlaceholderRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetWidth() <= 0 {
		err := GetPlaceholderRequestValidationError{
			field:  "Width",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetHeight() <= 0 {
		err := GetPlaceholderRequestValidationError{
			field:  "Height",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := len(m.GetBackgroundColor()); l < 0 || l > 6 {
		err := GetPlaceholderRequestValidationError{
			field:  "BackgroundColor",
			reason: "value length must be between 0 and 6 bytes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !_GetPlaceholderRequest_BackgroundColor_Pattern.MatchString(m.GetBackgroundColor()) {
		err := GetPlaceholderRequestValidationError{
			field:  "BackgroundColor",
			reason: "value does not match regex pattern \"[0-9A-Za-z]+\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := len(m.GetTextColor()); l < 0 || l > 6 {
		err := GetPlaceholderRequestValidationError{
			field:  "TextColor",
			reason: "value length must be between 0 and 6 bytes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !_GetPlaceholderRequest_TextColor_Pattern.MatchString(m.GetTextColor()) {
		err := GetPlaceholderRequestValidationError{
			field:  "TextColor",
			reason: "value does not match regex pattern \"[0-9A-Za-z]+\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := len(m.GetText()); l < 0 || l > 20 {
		err := GetPlaceholderRequestValidationError{
			field:  "Text",
			reason: "value length must be between 0 and 20 bytes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetFoundSize() <= 0 {
		err := GetPlaceholderRequestValidationError{
			field:  "FoundSize",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return GetPlaceholderRequestMultiError(errors)
	}

	return nil
}

// GetPlaceholderRequestMultiError is an error wrapping multiple validation
// errors returned by GetPlaceholderRequest.ValidateAll() if the designated
// constraints aren't met.
type GetPlaceholderRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetPlaceholderRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetPlaceholderRequestMultiError) AllErrors() []error { return m }

// GetPlaceholderRequestValidationError is the validation error returned by
// GetPlaceholderRequest.Validate if the designated constraints aren't met.
type GetPlaceholderRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetPlaceholderRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetPlaceholderRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetPlaceholderRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetPlaceholderRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetPlaceholderRequestValidationError) ErrorName() string {
	return "GetPlaceholderRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetPlaceholderRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetPlaceholderRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetPlaceholderRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetPlaceholderRequestValidationError{}

var _GetPlaceholderRequest_BackgroundColor_Pattern = regexp.MustCompile("[0-9A-Za-z]+")

var _GetPlaceholderRequest_TextColor_Pattern = regexp.MustCompile("[0-9A-Za-z]+")

// Validate checks the field values on GetPlaceholderResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetPlaceholderResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetPlaceholderResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetPlaceholderResponseMultiError, or nil if none found.
func (m *GetPlaceholderResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetPlaceholderResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Image

	if len(errors) > 0 {
		return GetPlaceholderResponseMultiError(errors)
	}

	return nil
}

// GetPlaceholderResponseMultiError is an error wrapping multiple validation
// errors returned by GetPlaceholderResponse.ValidateAll() if the designated
// constraints aren't met.
type GetPlaceholderResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetPlaceholderResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetPlaceholderResponseMultiError) AllErrors() []error { return m }

// GetPlaceholderResponseValidationError is the validation error returned by
// GetPlaceholderResponse.Validate if the designated constraints aren't met.
type GetPlaceholderResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetPlaceholderResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetPlaceholderResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetPlaceholderResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetPlaceholderResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetPlaceholderResponseValidationError) ErrorName() string {
	return "GetPlaceholderResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetPlaceholderResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetPlaceholderResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetPlaceholderResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetPlaceholderResponseValidationError{}
