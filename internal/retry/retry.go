package retry

import (
	"context"
	"errors"
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-log/tflog"
)

const (
	initialBackoff  = 1 * time.Second
	maxBackoff      = 30 * time.Second
	retryTimeoutMsg = "retry timeout exceeded while waiting for the operation to complete, the failure may be due to a transient issue or request cancellation"
)

// Can be overridden at the provider config level
var RetryTimeout = 60 * time.Second

type (
	RetryableFunc func(error) bool
	RetryFunc     func(ctx context.Context) (int, error)
)

// SetRetryTimeout sets the global retry timeout duration.
func SetRetryTimeout(timeout int64) {
	if timeout < 0 {
		return
	}
	RetryTimeout = time.Duration(timeout) * time.Second
}

// Do retries fn until:
// - fn succeeds
// - error is non-retryable
// - context is canceled or times out
func Do(parentCtx context.Context, isRetryable RetryableFunc, fn RetryFunc) error {
	ctx, cancel := context.WithTimeout(parentCtx, RetryTimeout)
	defer cancel()

	backoff := initialBackoff
	attempt := 0

	for {
		attempt++
		_, err := fn(ctx)
		if err == nil {
			return nil
		}

		// Stop retrying on context deadline exceeded, cancellation
		if errors.Is(err, context.DeadlineExceeded) || errors.Is(err, context.Canceled) ||
			errors.Is(ctx.Err(), context.DeadlineExceeded) || errors.Is(ctx.Err(), context.Canceled) {
			// Overriding the ctx deadline/cancellation error message for better user understanding
			return errors.New(retryTimeoutMsg)
		}

		// Stop retrying if error is not retryable
		if isRetryable == nil || !isRetryable(err) {
			return err
		}

		tflog.Debug(ctx, fmt.Sprintf(
			"Transient error detected, retrying request (attempt=%d, backoff=%s, err=%v)",
			attempt, backoff, err,
		))

		// Wait before retrying with exponential backoff
		select {
		case <-ctx.Done():
			return errors.New(retryTimeoutMsg)
		case <-time.After(backoff):
		}

		// Increase backoff for next iteration, capped at maxBackoff
		backoff *= 2
		if backoff > maxBackoff {
			backoff = maxBackoff
		}
	}
}

// TransientErrors determines if an error is retryable based on transient conditions.
// TODO: Currently returns false, treating all errors as non-retryable.
// This can be extended in the future to include predicates for specific transient errors
// such as network errors, temporary service unavailability (5xx errors), etc.
func TransientErrors(err error) bool {
	if err == nil {
		return false
	}
	// Everything is non-retryable for now
	return false
}

// IsNetworkError checks if the error is a network-related error.
func IsNetworkError(err error) bool {
	if err == nil {
		return false
	}

	// Check for net.Error interface
	var netErr net.Error
	if errors.As(err, &netErr) {
		return true
	}

	// TODO: Extend these error patterns as per future requirements.
	// Check for common network error strings (case-insensitive)
	errStr := strings.ToLower(err.Error())
	networkPatterns := []string{
		"connection refused",
		"connection reset",
		"broken pipe",
		"eof",
		"connection closed",
		"dial tcp",
	}

	for _, pattern := range networkPatterns {
		if strings.Contains(errStr, pattern) {
			return true
		}
	}

	return false
}

// IsAlreadyExistsErr checks if the error indicates that the resource already exists.
func IsAlreadyExistsErr(err error) bool {
	if err == nil {
		return false
	}
	errStr := strings.ToLower(err.Error())
	return strings.Contains(errStr, "already exists") || strings.Contains(errStr, "conflict")
}
