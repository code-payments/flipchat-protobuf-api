// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: promoted/v1/promoted_service.proto

package promotedpb

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

// Validate checks the field values on GetPromotedChatsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetPromotedChatsRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetPromotedChatsRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetPromotedChatsRequestMultiError, or nil if none found.
func (m *GetPromotedChatsRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetPromotedChatsRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetTopic()); l < 1 || l > 100 {
		err := GetPromotedChatsRequestValidationError{
			field:  "Topic",
			reason: "value length must be between 1 and 100 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return GetPromotedChatsRequestMultiError(errors)
	}

	return nil
}

// GetPromotedChatsRequestMultiError is an error wrapping multiple validation
// errors returned by GetPromotedChatsRequest.ValidateAll() if the designated
// constraints aren't met.
type GetPromotedChatsRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetPromotedChatsRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetPromotedChatsRequestMultiError) AllErrors() []error { return m }

// GetPromotedChatsRequestValidationError is the validation error returned by
// GetPromotedChatsRequest.Validate if the designated constraints aren't met.
type GetPromotedChatsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetPromotedChatsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetPromotedChatsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetPromotedChatsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetPromotedChatsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetPromotedChatsRequestValidationError) ErrorName() string {
	return "GetPromotedChatsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetPromotedChatsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetPromotedChatsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetPromotedChatsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetPromotedChatsRequestValidationError{}

// Validate checks the field values on GetPromotedChatsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetPromotedChatsResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetPromotedChatsResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetPromotedChatsResponseMultiError, or nil if none found.
func (m *GetPromotedChatsResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetPromotedChatsResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Result

	for idx, item := range m.GetChats() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, GetPromotedChatsResponseValidationError{
						field:  fmt.Sprintf("Chats[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetPromotedChatsResponseValidationError{
						field:  fmt.Sprintf("Chats[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetPromotedChatsResponseValidationError{
					field:  fmt.Sprintf("Chats[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return GetPromotedChatsResponseMultiError(errors)
	}

	return nil
}

// GetPromotedChatsResponseMultiError is an error wrapping multiple validation
// errors returned by GetPromotedChatsResponse.ValidateAll() if the designated
// constraints aren't met.
type GetPromotedChatsResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetPromotedChatsResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetPromotedChatsResponseMultiError) AllErrors() []error { return m }

// GetPromotedChatsResponseValidationError is the validation error returned by
// GetPromotedChatsResponse.Validate if the designated constraints aren't met.
type GetPromotedChatsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetPromotedChatsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetPromotedChatsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetPromotedChatsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetPromotedChatsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetPromotedChatsResponseValidationError) ErrorName() string {
	return "GetPromotedChatsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetPromotedChatsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetPromotedChatsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetPromotedChatsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetPromotedChatsResponseValidationError{}
