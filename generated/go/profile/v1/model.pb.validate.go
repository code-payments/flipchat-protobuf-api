// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: profile/v1/model.proto

package profilepb

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

// Validate checks the field values on UserProfile with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *UserProfile) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UserProfile with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in UserProfileMultiError, or
// nil if none found.
func (m *UserProfile) ValidateAll() error {
	return m.validate(true)
}

func (m *UserProfile) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetDisplayName()); l < 0 || l > 64 {
		err := UserProfileValidationError{
			field:  "DisplayName",
			reason: "value length must be between 0 and 64 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(m.GetSocialProfiles()) > 1 {
		err := UserProfileValidationError{
			field:  "SocialProfiles",
			reason: "value must contain no more than 1 item(s)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	for idx, item := range m.GetSocialProfiles() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, UserProfileValidationError{
						field:  fmt.Sprintf("SocialProfiles[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, UserProfileValidationError{
						field:  fmt.Sprintf("SocialProfiles[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return UserProfileValidationError{
					field:  fmt.Sprintf("SocialProfiles[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return UserProfileMultiError(errors)
	}

	return nil
}

// UserProfileMultiError is an error wrapping multiple validation errors
// returned by UserProfile.ValidateAll() if the designated constraints aren't met.
type UserProfileMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserProfileMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserProfileMultiError) AllErrors() []error { return m }

// UserProfileValidationError is the validation error returned by
// UserProfile.Validate if the designated constraints aren't met.
type UserProfileValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserProfileValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserProfileValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserProfileValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserProfileValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserProfileValidationError) ErrorName() string { return "UserProfileValidationError" }

// Error satisfies the builtin error interface
func (e UserProfileValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserProfile.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserProfileValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserProfileValidationError{}

// Validate checks the field values on SocialProfile with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SocialProfile) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SocialProfile with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SocialProfileMultiError, or
// nil if none found.
func (m *SocialProfile) ValidateAll() error {
	return m.validate(true)
}

func (m *SocialProfile) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	oneofTypePresent := false
	switch v := m.Type.(type) {
	case *SocialProfile_X:
		if v == nil {
			err := SocialProfileValidationError{
				field:  "Type",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofTypePresent = true

		if all {
			switch v := interface{}(m.GetX()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, SocialProfileValidationError{
						field:  "X",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, SocialProfileValidationError{
						field:  "X",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetX()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SocialProfileValidationError{
					field:  "X",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		_ = v // ensures v is used
	}
	if !oneofTypePresent {
		err := SocialProfileValidationError{
			field:  "Type",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return SocialProfileMultiError(errors)
	}

	return nil
}

// SocialProfileMultiError is an error wrapping multiple validation errors
// returned by SocialProfile.ValidateAll() if the designated constraints
// aren't met.
type SocialProfileMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SocialProfileMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SocialProfileMultiError) AllErrors() []error { return m }

// SocialProfileValidationError is the validation error returned by
// SocialProfile.Validate if the designated constraints aren't met.
type SocialProfileValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SocialProfileValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SocialProfileValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SocialProfileValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SocialProfileValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SocialProfileValidationError) ErrorName() string { return "SocialProfileValidationError" }

// Error satisfies the builtin error interface
func (e SocialProfileValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSocialProfile.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SocialProfileValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SocialProfileValidationError{}

// Validate checks the field values on XProfile with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *XProfile) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on XProfile with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in XProfileMultiError, or nil
// if none found.
func (m *XProfile) ValidateAll() error {
	return m.validate(true)
}

func (m *XProfile) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetUsername()); l < 1 || l > 15 {
		err := XProfileValidationError{
			field:  "Username",
			reason: "value length must be between 1 and 15 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetName()); l < 1 || l > 256 {
		err := XProfileValidationError{
			field:  "Name",
			reason: "value length must be between 1 and 256 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetDescription()); l < 1 || l > 4096 {
		err := XProfileValidationError{
			field:  "Description",
			reason: "value length must be between 1 and 4096 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetProfilePicUrl()); l < 1 || l > 2048 {
		err := XProfileValidationError{
			field:  "ProfilePicUrl",
			reason: "value length must be between 1 and 2048 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for VerifiedType

	// no validation rules for FollowerCount

	if len(errors) > 0 {
		return XProfileMultiError(errors)
	}

	return nil
}

// XProfileMultiError is an error wrapping multiple validation errors returned
// by XProfile.ValidateAll() if the designated constraints aren't met.
type XProfileMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m XProfileMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m XProfileMultiError) AllErrors() []error { return m }

// XProfileValidationError is the validation error returned by
// XProfile.Validate if the designated constraints aren't met.
type XProfileValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e XProfileValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e XProfileValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e XProfileValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e XProfileValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e XProfileValidationError) ErrorName() string { return "XProfileValidationError" }

// Error satisfies the builtin error interface
func (e XProfileValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sXProfile.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = XProfileValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = XProfileValidationError{}
