// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: iap/v1/iap_service.proto

package iappb

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

	commonpb "github.com/code-payments/flipchat-protobuf-api/generated/go/common/v1"
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

	_ = commonpb.Platform(0)
)

// Validate checks the field values on OnPurchaseCompletedRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *OnPurchaseCompletedRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OnPurchaseCompletedRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// OnPurchaseCompletedRequestMultiError, or nil if none found.
func (m *OnPurchaseCompletedRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *OnPurchaseCompletedRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if _, ok := _OnPurchaseCompletedRequest_Platform_InLookup[m.GetPlatform()]; !ok {
		err := OnPurchaseCompletedRequestValidationError{
			field:  "Platform",
			reason: "value must be in list [APPLE GOOGLE]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetReceipt() == nil {
		err := OnPurchaseCompletedRequestValidationError{
			field:  "Receipt",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetReceipt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, OnPurchaseCompletedRequestValidationError{
					field:  "Receipt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, OnPurchaseCompletedRequestValidationError{
					field:  "Receipt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetReceipt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OnPurchaseCompletedRequestValidationError{
				field:  "Receipt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetAuth() == nil {
		err := OnPurchaseCompletedRequestValidationError{
			field:  "Auth",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetAuth()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, OnPurchaseCompletedRequestValidationError{
					field:  "Auth",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, OnPurchaseCompletedRequestValidationError{
					field:  "Auth",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetAuth()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OnPurchaseCompletedRequestValidationError{
				field:  "Auth",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return OnPurchaseCompletedRequestMultiError(errors)
	}

	return nil
}

// OnPurchaseCompletedRequestMultiError is an error wrapping multiple
// validation errors returned by OnPurchaseCompletedRequest.ValidateAll() if
// the designated constraints aren't met.
type OnPurchaseCompletedRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OnPurchaseCompletedRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OnPurchaseCompletedRequestMultiError) AllErrors() []error { return m }

// OnPurchaseCompletedRequestValidationError is the validation error returned
// by OnPurchaseCompletedRequest.Validate if the designated constraints aren't met.
type OnPurchaseCompletedRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OnPurchaseCompletedRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OnPurchaseCompletedRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OnPurchaseCompletedRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OnPurchaseCompletedRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OnPurchaseCompletedRequestValidationError) ErrorName() string {
	return "OnPurchaseCompletedRequestValidationError"
}

// Error satisfies the builtin error interface
func (e OnPurchaseCompletedRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOnPurchaseCompletedRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OnPurchaseCompletedRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OnPurchaseCompletedRequestValidationError{}

var _OnPurchaseCompletedRequest_Platform_InLookup = map[commonpb.Platform]struct{}{
	1: {},
	2: {},
}

// Validate checks the field values on OnPurchaseCompletedResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *OnPurchaseCompletedResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OnPurchaseCompletedResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// OnPurchaseCompletedResponseMultiError, or nil if none found.
func (m *OnPurchaseCompletedResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *OnPurchaseCompletedResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Result

	if len(errors) > 0 {
		return OnPurchaseCompletedResponseMultiError(errors)
	}

	return nil
}

// OnPurchaseCompletedResponseMultiError is an error wrapping multiple
// validation errors returned by OnPurchaseCompletedResponse.ValidateAll() if
// the designated constraints aren't met.
type OnPurchaseCompletedResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OnPurchaseCompletedResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OnPurchaseCompletedResponseMultiError) AllErrors() []error { return m }

// OnPurchaseCompletedResponseValidationError is the validation error returned
// by OnPurchaseCompletedResponse.Validate if the designated constraints
// aren't met.
type OnPurchaseCompletedResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OnPurchaseCompletedResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OnPurchaseCompletedResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OnPurchaseCompletedResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OnPurchaseCompletedResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OnPurchaseCompletedResponseValidationError) ErrorName() string {
	return "OnPurchaseCompletedResponseValidationError"
}

// Error satisfies the builtin error interface
func (e OnPurchaseCompletedResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOnPurchaseCompletedResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OnPurchaseCompletedResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OnPurchaseCompletedResponseValidationError{}

// Validate checks the field values on Receipt with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Receipt) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Receipt with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in ReceiptMultiError, or nil if none found.
func (m *Receipt) ValidateAll() error {
	return m.validate(true)
}

func (m *Receipt) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetValue()); l < 1 || l > 4096 {
		err := ReceiptValidationError{
			field:  "Value",
			reason: "value length must be between 1 and 4096 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return ReceiptMultiError(errors)
	}

	return nil
}

// ReceiptMultiError is an error wrapping multiple validation errors returned
// by Receipt.ValidateAll() if the designated constraints aren't met.
type ReceiptMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ReceiptMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ReceiptMultiError) AllErrors() []error { return m }

// ReceiptValidationError is the validation error returned by Receipt.Validate
// if the designated constraints aren't met.
type ReceiptValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ReceiptValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ReceiptValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ReceiptValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ReceiptValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ReceiptValidationError) ErrorName() string { return "ReceiptValidationError" }

// Error satisfies the builtin error interface
func (e ReceiptValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sReceipt.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ReceiptValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ReceiptValidationError{}
