// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: blob/v1/model.proto

package blobpb

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

// Validate checks the field values on Blob with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Blob) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Blob with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in BlobMultiError, or nil if none found.
func (m *Blob) ValidateAll() error {
	return m.validate(true)
}

func (m *Blob) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetBlobId()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, BlobValidationError{
					field:  "BlobId",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, BlobValidationError{
					field:  "BlobId",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetBlobId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return BlobValidationError{
				field:  "BlobId",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for BlobType

	if all {
		switch v := interface{}(m.GetOwnerId()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, BlobValidationError{
					field:  "OwnerId",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, BlobValidationError{
					field:  "OwnerId",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetOwnerId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return BlobValidationError{
				field:  "OwnerId",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for S3Url

	if all {
		switch v := interface{}(m.GetCreatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, BlobValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, BlobValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return BlobValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return BlobMultiError(errors)
	}

	return nil
}

// BlobMultiError is an error wrapping multiple validation errors returned by
// Blob.ValidateAll() if the designated constraints aren't met.
type BlobMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m BlobMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m BlobMultiError) AllErrors() []error { return m }

// BlobValidationError is the validation error returned by Blob.Validate if the
// designated constraints aren't met.
type BlobValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BlobValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BlobValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BlobValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BlobValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BlobValidationError) ErrorName() string { return "BlobValidationError" }

// Error satisfies the builtin error interface
func (e BlobValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBlob.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BlobValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BlobValidationError{}
