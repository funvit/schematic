package schemagen

import "github.com/funvit/schematic/pkg/scope_error"

var (
	// ErrModelValidationError represents scope of model validation errors.
	ErrModelValidationError = scope_error.New("model validation error")
	// ErrModelFieldValidationError represents scope of model field validation errors.
	ErrModelFieldValidationError = scope_error.New("model field validation error")
)
