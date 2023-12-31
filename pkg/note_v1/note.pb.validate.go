// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: note.proto

package note_v1

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

// Validate checks the field values on NoteInfo with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *NoteInfo) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on NoteInfo with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in NoteInfoMultiError, or nil
// if none found.
func (m *NoteInfo) ValidateAll() error {
	return m.validate(true)
}

func (m *NoteInfo) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetTitle()) > 100 {
		err := NoteInfoValidationError{
			field:  "Title",
			reason: "value length must be at most 100 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(m.GetText()) > 4096 {
		err := NoteInfoValidationError{
			field:  "Text",
			reason: "value length must be at most 4096 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetAuthor()) > 100 {
		err := NoteInfoValidationError{
			field:  "Author",
			reason: "value length must be at most 100 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if err := m._validateEmail(m.GetEmail()); err != nil {
		err = NoteInfoValidationError{
			field:  "Email",
			reason: "value must be a valid email address",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return NoteInfoMultiError(errors)
	}

	return nil
}

func (m *NoteInfo) _validateHostname(host string) error {
	s := strings.ToLower(strings.TrimSuffix(host, "."))

	if len(host) > 253 {
		return errors.New("hostname cannot exceed 253 characters")
	}

	for _, part := range strings.Split(s, ".") {
		if l := len(part); l == 0 || l > 63 {
			return errors.New("hostname part must be non-empty and cannot exceed 63 characters")
		}

		if part[0] == '-' {
			return errors.New("hostname parts cannot begin with hyphens")
		}

		if part[len(part)-1] == '-' {
			return errors.New("hostname parts cannot end with hyphens")
		}

		for _, r := range part {
			if (r < 'a' || r > 'z') && (r < '0' || r > '9') && r != '-' {
				return fmt.Errorf("hostname parts can only contain alphanumeric characters or hyphens, got %q", string(r))
			}
		}
	}

	return nil
}

func (m *NoteInfo) _validateEmail(addr string) error {
	a, err := mail.ParseAddress(addr)
	if err != nil {
		return err
	}
	addr = a.Address

	if len(addr) > 254 {
		return errors.New("email addresses cannot exceed 254 characters")
	}

	parts := strings.SplitN(addr, "@", 2)

	if len(parts[0]) > 64 {
		return errors.New("email address local phrase cannot exceed 64 characters")
	}

	return m._validateHostname(parts[1])
}

// NoteInfoMultiError is an error wrapping multiple validation errors returned
// by NoteInfo.ValidateAll() if the designated constraints aren't met.
type NoteInfoMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m NoteInfoMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m NoteInfoMultiError) AllErrors() []error { return m }

// NoteInfoValidationError is the validation error returned by
// NoteInfo.Validate if the designated constraints aren't met.
type NoteInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NoteInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NoteInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NoteInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NoteInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NoteInfoValidationError) ErrorName() string { return "NoteInfoValidationError" }

// Error satisfies the builtin error interface
func (e NoteInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNoteInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NoteInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NoteInfoValidationError{}

// Validate checks the field values on UpdateNoteInfo with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *UpdateNoteInfo) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateNoteInfo with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in UpdateNoteInfoMultiError,
// or nil if none found.
func (m *UpdateNoteInfo) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateNoteInfo) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetTitle()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdateNoteInfoValidationError{
					field:  "Title",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdateNoteInfoValidationError{
					field:  "Title",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTitle()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateNoteInfoValidationError{
				field:  "Title",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetText()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdateNoteInfoValidationError{
					field:  "Text",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdateNoteInfoValidationError{
					field:  "Text",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetText()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateNoteInfoValidationError{
				field:  "Text",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetAuthor()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdateNoteInfoValidationError{
					field:  "Author",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdateNoteInfoValidationError{
					field:  "Author",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetAuthor()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateNoteInfoValidationError{
				field:  "Author",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetEmail()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdateNoteInfoValidationError{
					field:  "Email",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdateNoteInfoValidationError{
					field:  "Email",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetEmail()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateNoteInfoValidationError{
				field:  "Email",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return UpdateNoteInfoMultiError(errors)
	}

	return nil
}

// UpdateNoteInfoMultiError is an error wrapping multiple validation errors
// returned by UpdateNoteInfo.ValidateAll() if the designated constraints
// aren't met.
type UpdateNoteInfoMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateNoteInfoMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateNoteInfoMultiError) AllErrors() []error { return m }

// UpdateNoteInfoValidationError is the validation error returned by
// UpdateNoteInfo.Validate if the designated constraints aren't met.
type UpdateNoteInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateNoteInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateNoteInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateNoteInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateNoteInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateNoteInfoValidationError) ErrorName() string { return "UpdateNoteInfoValidationError" }

// Error satisfies the builtin error interface
func (e UpdateNoteInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateNoteInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateNoteInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateNoteInfoValidationError{}

// Validate checks the field values on Note with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Note) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Note with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in NoteMultiError, or nil if none found.
func (m *Note) ValidateAll() error {
	return m.validate(true)
}

func (m *Note) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if all {
		switch v := interface{}(m.GetNote()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, NoteValidationError{
					field:  "Note",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, NoteValidationError{
					field:  "Note",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetNote()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return NoteValidationError{
				field:  "Note",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetCreatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, NoteValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, NoteValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return NoteValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetUpdatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, NoteValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, NoteValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return NoteValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return NoteMultiError(errors)
	}

	return nil
}

// NoteMultiError is an error wrapping multiple validation errors returned by
// Note.ValidateAll() if the designated constraints aren't met.
type NoteMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m NoteMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m NoteMultiError) AllErrors() []error { return m }

// NoteValidationError is the validation error returned by Note.Validate if the
// designated constraints aren't met.
type NoteValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NoteValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NoteValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NoteValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NoteValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NoteValidationError) ErrorName() string { return "NoteValidationError" }

// Error satisfies the builtin error interface
func (e NoteValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNote.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NoteValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NoteValidationError{}

// Validate checks the field values on CreateRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CreateRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CreateRequestMultiError, or
// nil if none found.
func (m *CreateRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() <= 0 {
		err := CreateRequestValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetNote()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateRequestValidationError{
					field:  "Note",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateRequestValidationError{
					field:  "Note",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetNote()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateRequestValidationError{
				field:  "Note",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreateRequestMultiError(errors)
	}

	return nil
}

// CreateRequestMultiError is an error wrapping multiple validation errors
// returned by CreateRequest.ValidateAll() if the designated constraints
// aren't met.
type CreateRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateRequestMultiError) AllErrors() []error { return m }

// CreateRequestValidationError is the validation error returned by
// CreateRequest.Validate if the designated constraints aren't met.
type CreateRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateRequestValidationError) ErrorName() string { return "CreateRequestValidationError" }

// Error satisfies the builtin error interface
func (e CreateRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateRequestValidationError{}

// Validate checks the field values on CreateResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CreateResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CreateResponseMultiError,
// or nil if none found.
func (m *CreateResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() <= 0 {
		err := CreateResponseValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return CreateResponseMultiError(errors)
	}

	return nil
}

// CreateResponseMultiError is an error wrapping multiple validation errors
// returned by CreateResponse.ValidateAll() if the designated constraints
// aren't met.
type CreateResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateResponseMultiError) AllErrors() []error { return m }

// CreateResponseValidationError is the validation error returned by
// CreateResponse.Validate if the designated constraints aren't met.
type CreateResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateResponseValidationError) ErrorName() string { return "CreateResponseValidationError" }

// Error satisfies the builtin error interface
func (e CreateResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateResponseValidationError{}

// Validate checks the field values on GetRequest with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetRequestMultiError, or
// nil if none found.
func (m *GetRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() <= 0 {
		err := GetRequestValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return GetRequestMultiError(errors)
	}

	return nil
}

// GetRequestMultiError is an error wrapping multiple validation errors
// returned by GetRequest.ValidateAll() if the designated constraints aren't met.
type GetRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetRequestMultiError) AllErrors() []error { return m }

// GetRequestValidationError is the validation error returned by
// GetRequest.Validate if the designated constraints aren't met.
type GetRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetRequestValidationError) ErrorName() string { return "GetRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetRequestValidationError{}

// Validate checks the field values on GetResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetResponseMultiError, or
// nil if none found.
func (m *GetResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetNote()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetResponseValidationError{
					field:  "Note",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetResponseValidationError{
					field:  "Note",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetNote()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetResponseValidationError{
				field:  "Note",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetResponseMultiError(errors)
	}

	return nil
}

// GetResponseMultiError is an error wrapping multiple validation errors
// returned by GetResponse.ValidateAll() if the designated constraints aren't met.
type GetResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetResponseMultiError) AllErrors() []error { return m }

// GetResponseValidationError is the validation error returned by
// GetResponse.Validate if the designated constraints aren't met.
type GetResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetResponseValidationError) ErrorName() string { return "GetResponseValidationError" }

// Error satisfies the builtin error interface
func (e GetResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetResponseValidationError{}

// Validate checks the field values on GetListResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetListResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetListResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetListResponseMultiError, or nil if none found.
func (m *GetListResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetListResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetNotes() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, GetListResponseValidationError{
						field:  fmt.Sprintf("Notes[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetListResponseValidationError{
						field:  fmt.Sprintf("Notes[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetListResponseValidationError{
					field:  fmt.Sprintf("Notes[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return GetListResponseMultiError(errors)
	}

	return nil
}

// GetListResponseMultiError is an error wrapping multiple validation errors
// returned by GetListResponse.ValidateAll() if the designated constraints
// aren't met.
type GetListResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetListResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetListResponseMultiError) AllErrors() []error { return m }

// GetListResponseValidationError is the validation error returned by
// GetListResponse.Validate if the designated constraints aren't met.
type GetListResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetListResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetListResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetListResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetListResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetListResponseValidationError) ErrorName() string { return "GetListResponseValidationError" }

// Error satisfies the builtin error interface
func (e GetListResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetListResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetListResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetListResponseValidationError{}

// Validate checks the field values on UpdateRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *UpdateRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in UpdateRequestMultiError, or
// nil if none found.
func (m *UpdateRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() <= 0 {
		err := UpdateRequestValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetNote()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdateRequestValidationError{
					field:  "Note",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdateRequestValidationError{
					field:  "Note",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetNote()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateRequestValidationError{
				field:  "Note",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return UpdateRequestMultiError(errors)
	}

	return nil
}

// UpdateRequestMultiError is an error wrapping multiple validation errors
// returned by UpdateRequest.ValidateAll() if the designated constraints
// aren't met.
type UpdateRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateRequestMultiError) AllErrors() []error { return m }

// UpdateRequestValidationError is the validation error returned by
// UpdateRequest.Validate if the designated constraints aren't met.
type UpdateRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateRequestValidationError) ErrorName() string { return "UpdateRequestValidationError" }

// Error satisfies the builtin error interface
func (e UpdateRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateRequestValidationError{}

// Validate checks the field values on DeleteRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *DeleteRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in DeleteRequestMultiError, or
// nil if none found.
func (m *DeleteRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() <= 0 {
		err := DeleteRequestValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return DeleteRequestMultiError(errors)
	}

	return nil
}

// DeleteRequestMultiError is an error wrapping multiple validation errors
// returned by DeleteRequest.ValidateAll() if the designated constraints
// aren't met.
type DeleteRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteRequestMultiError) AllErrors() []error { return m }

// DeleteRequestValidationError is the validation error returned by
// DeleteRequest.Validate if the designated constraints aren't met.
type DeleteRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteRequestValidationError) ErrorName() string { return "DeleteRequestValidationError" }

// Error satisfies the builtin error interface
func (e DeleteRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteRequestValidationError{}
