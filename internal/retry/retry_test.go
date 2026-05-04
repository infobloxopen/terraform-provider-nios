package retry

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"
)

// TestDo_Success tests that Do returns nil when the function succeeds
func TestDo_Success(t *testing.T) {
	callCount := 0
	fn := func(ctx context.Context) (int, error) {
		callCount++
		return 200, nil
	}

	err := Do(context.Background(), nil, fn)
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
	if callCount != 1 {
		t.Errorf("Expected 1 call, got: %d", callCount)
	}
}

// TestDo_NonRetryableError tests that Do returns immediately on non-retryable errors
func TestDo_NonRetryableError(t *testing.T) {
	callCount := 0
	expectedErr := errors.New("non-retryable error")

	fn := func(ctx context.Context) (int, error) {
		callCount++
		return 500, expectedErr
	}

	isRetryable := func(err error) bool {
		return false // Not retryable
	}

	err := Do(context.Background(), isRetryable, fn)
	if err == nil {
		t.Error("Expected error, got nil")
	}
	if !errors.Is(err, expectedErr) {
		t.Errorf("Expected error %v, got: %v", expectedErr, err)
	}
	if callCount != 1 {
		t.Errorf("Expected 1 call for non-retryable error, got: %d", callCount)
	}
}

// TestDo_RetryableErrorWithEventualSuccess tests retrying until success
func TestDo_RetryableErrorWithEventualSuccess(t *testing.T) {
	// Reduce timeout for faster test execution
	originalTimeout := RetryTimeout
	SetRetryTimeout(5)
	defer func() { RetryTimeout = originalTimeout }()

	callCount := 0
	fn := func(ctx context.Context) (int, error) {
		callCount++
		if callCount < 3 {
			return 500, errors.New("temporary error")
		}
		return 200, nil
	}

	isRetryable := func(err error) bool {
		return true // Always retryable
	}

	err := Do(context.Background(), isRetryable, fn)
	if err != nil {
		t.Errorf("Expected no error after retries, got: %v", err)
	}
	if callCount != 3 {
		t.Errorf("Expected 3 calls, got: %d", callCount)
	}
}

// TestDo_ContextTimeout tests that Do respects context timeout
func TestDo_ContextTimeout(t *testing.T) {
	// Set a very short timeout
	originalTimeout := RetryTimeout
	SetRetryTimeout(1)
	defer func() { RetryTimeout = originalTimeout }()

	callCount := 0
	fn := func(ctx context.Context) (int, error) {
		callCount++
		// Block until the retry context times out or is cancelled
		<-ctx.Done()
		return 500, ctx.Err()
	}

	isRetryable := func(err error) bool {
		return true
	}

	err := Do(context.Background(), isRetryable, fn)
	if err == nil {
		t.Error("Expected timeout error, got nil")
	}
	if err.Error() != retryTimeoutMsg {
		t.Errorf("Expected timeout message, got: %v", err)
	}
}

// TestDo_ContextCancellation tests that Do respects context cancellation
func TestDo_ContextCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	callCount := 0
	fn := func(ctx context.Context) (int, error) {
		callCount++
		if callCount == 2 {
			cancel()
		}
		return 500, errors.New("retryable error")
	}

	isRetryable := func(err error) bool { return true }

	err := Do(ctx, isRetryable, fn)
	if err == nil {
		t.Errorf("expected cancellation error, got nil")
	}
	if err.Error() != retryTimeoutMsg {
		t.Errorf("expected retry timeout message, got: %v", err)
	}
	if callCount < 2 {
		t.Errorf("expected at least 2 calls, got: %d", callCount)
	}
}

// TestDo_NilRetryableFunc tests behavior when isRetryable is nil
func TestDo_NilRetryableFunc(t *testing.T) {
	callCount := 0
	expectedErr := errors.New("some error")

	fn := func(ctx context.Context) (int, error) {
		callCount++
		return 500, expectedErr
	}

	err := Do(context.Background(), nil, fn)
	if err == nil {
		t.Error("Expected error, got nil")
	}
	if !errors.Is(err, expectedErr) {
		t.Errorf("Expected error %v, got: %v", expectedErr, err)
	}
	if callCount != 1 {
		t.Errorf("Expected 1 call when isRetryable is nil, got: %d", callCount)
	}
}

// TestSetRetryTimeout tests the SetRetryTimeout function
func TestSetRetryTimeout(t *testing.T) {
	originalTimeout := RetryTimeout
	defer func() { RetryTimeout = originalTimeout }()

	SetRetryTimeout(10)
	if RetryTimeout != 10*time.Second {
		t.Fatalf("expected 10s, got %v", RetryTimeout)
	}

	SetRetryTimeout(0)
	if RetryTimeout != 0 {
		t.Fatalf("expected 0, got %v", RetryTimeout)
	}

	RetryTimeout = 5 * time.Second
	SetRetryTimeout(-1)
	if RetryTimeout != 5*time.Second {
		t.Fatalf("negative timeout should not change value")
	}
}

// ExampleDo_basicRetry demonstrates basic retry usage
func ExampleDo_basicRetry() {
	// Define a function that might fail
	attemptCount := 0
	operation := func(ctx context.Context) (int, error) {
		attemptCount++
		if attemptCount < 3 {
			return 500, errors.New("temporary failure")
		}
		return 200, nil
	}

	// Define what errors are retryable
	isRetryable := func(err error) bool {
		return err.Error() == "temporary failure"
	}

	// Execute with retry
	err := Do(context.Background(), isRetryable, operation)
	if err != nil {
		fmt.Println("Operation failed:", err)
	} else {
		fmt.Println("Operation succeeded after", attemptCount, "attempts")
	}
	// Output: Operation succeeded after 3 attempts
}

// ExampleDo_networkErrorRetry demonstrates retrying on network errors
func ExampleDo_networkErrorRetry() {
	attempts := 0

	operation := func(ctx context.Context) (int, error) {
		attempts++
		if attempts < 2 {
			return 0, errors.New("dial tcp: connection refused")
		}
		return 200, nil
	}

	isRetryable := func(err error) bool {
		return IsNetworkError(err)
	}

	err := Do(context.Background(), isRetryable, operation)
	if err != nil {
		fmt.Println("Network operation failed:", err)
	} else {
		fmt.Println("Network operation succeeded after", attempts, "attempts")
	}

	// Output: Network operation succeeded after 2 attempts
}

// TestDoWithTimeout_Success tests that DoWithTimeout works with custom timeout
func TestDoWithTimeout_Success(t *testing.T) {
	callCount := 0
	fn := func(ctx context.Context) (int, error) {
		callCount++
		return 200, nil
	}

	err := DoWithTimeout(context.Background(), 5*time.Second, nil, fn)
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
	if callCount != 1 {
		t.Errorf("Expected 1 call, got: %d", callCount)
	}
}

// TestDoWithTimeout_CustomTimeout tests that custom timeout is respected
func TestDoWithTimeout_CustomTimeout(t *testing.T) {
	callCount := 0
	fn := func(ctx context.Context) (int, error) {
		callCount++
		<-ctx.Done()
		return 500, ctx.Err()
	}

	isRetryable := func(err error) bool {
		return true
	}

	// Use a very short timeout
	err := DoWithTimeout(context.Background(), 500*time.Millisecond, isRetryable, fn)
	if err == nil {
		t.Error("Expected timeout error, got nil")
	}
	if err.Error() != retryTimeoutMsg {
		t.Errorf("Expected timeout message, got: %v", err)
	}
}
