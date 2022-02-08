//go:build !fast
// +build !fast

package test

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	ydb "github.com/ydb-platform/ydb-go-sdk/v3"
	"github.com/ydb-platform/ydb-go-sdk/v3/balancers"
	"github.com/ydb-platform/ydb-go-sdk/v3/config"
	"github.com/ydb-platform/ydb-go-sdk/v3/retry"
	"github.com/ydb-platform/ydb-go-sdk/v3/scripting"
	"github.com/ydb-platform/ydb-go-sdk/v3/table"
	"github.com/ydb-platform/ydb-go-sdk/v3/trace"
)

func TestScripting(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	db, err := ydb.New(
		ctx,
		ydb.WithConnectionString(os.Getenv("YDB_CONNECTION_STRING")),
		ydb.WithAnonymousCredentials(),
		ydb.With(
			config.WithRequestTimeout(time.Second*2),
			config.WithStreamTimeout(time.Second*2),
			config.WithOperationTimeout(time.Second*2),
			config.WithOperationCancelAfter(time.Second*2),
		),
		ydb.WithBalancer(balancers.SingleConn()),
		ydb.WithConnectionTTL(time.Millisecond*10000),
		ydb.WithLogger(
			trace.DriverConnEvents,
			ydb.WithNamespace("ydb"),
			ydb.WithOutWriter(os.Stdout),
			ydb.WithErrWriter(os.Stderr),
			ydb.WithMinLevel(ydb.TRACE),
		),
		ydb.WithUserAgent("scripting"),
	)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		// cleanup connection
		if e := db.Close(ctx); e != nil {
			t.Fatalf("db close failed: %+v", e)
		}
	}()
	// Execute
	if err = retry.Retry(ctx, true, func(ctx context.Context) (err error) {
		res, err := db.Scripting().Execute(
			ctx,
			"SELECT 1+1",
			table.NewQueryParameters(),
		)
		if err != nil {
			return err
		}
		defer func() {
			_ = res.Close()
		}()
		if !res.NextResultSet(ctx) {
			return retry.RetryableError(
				fmt.Errorf("no result sets"),
				retry.WithBackoff(retry.BackoffTypeNoBackoff),
			)
		}
		if !res.NextRow() {
			return retry.RetryableError(
				fmt.Errorf("no rows"),
				retry.WithBackoff(retry.BackoffTypeSlowBackoff),
			)
		}
		var sum int32
		if err = res.Scan(&sum); err != nil {
			return fmt.Errorf("scan failed: %w", err)
		}
		if sum != 2 {
			return fmt.Errorf("unexpected sum: %v", sum)
		}
		return res.Err()
	}); err != nil {
		t.Fatalf("Execute failed: %v", err)
	}
	// StreamExecute
	if err = retry.Retry(ctx, true, func(ctx context.Context) (err error) {
		res, err := db.Scripting().StreamExecute(
			ctx,
			"SELECT 1+1",
			table.NewQueryParameters(),
		)
		if err != nil {
			return err
		}
		defer func() {
			_ = res.Close()
		}()
		if !res.NextResultSet(ctx) {
			return retry.RetryableError(
				fmt.Errorf("no result sets"),
				retry.WithBackoff(retry.BackoffTypeNoBackoff),
				retry.WithDeleteSession(),
			)
		}
		if !res.NextRow() {
			return retry.RetryableError(
				fmt.Errorf("no rows"),
				retry.WithBackoff(retry.BackoffTypeFastBackoff),
			)
		}
		var sum int32
		if err = res.Scan(&sum); err != nil {
			return fmt.Errorf("scan failed: %w", err)
		}
		if sum != 2 {
			return fmt.Errorf("unexpected sum: %v", sum)
		}
		return res.Err()
	}); err != nil {
		t.Fatalf("StreamExecute failed: %v", err)
	}
	// ExplainPlan
	if err = retry.Retry(ctx, true, func(ctx context.Context) (err error) {
		res, err := db.Scripting().Explain(
			ctx,
			"SELECT 1+1",
			scripting.ExplainModePlan,
		)
		if err != nil {
			return err
		}
		if res.Plan == "" {
			return fmt.Errorf("empty plan")
		}
		return nil
	}); err != nil {
		t.Fatalf("Explain failed: %v", err)
	}
	// ExplainValidate
	if err = retry.Retry(ctx, true, func(ctx context.Context) (err error) {
		res, err := db.Scripting().Explain(
			ctx,
			"SELECT 1+1",
			scripting.ExplainModeValidate,
		)
		if err != nil {
			return err
		}
		if len(res.ParameterTypes) > 0 {
			return fmt.Errorf("unexpected parameter types")
		}
		return nil
	}); err != nil {
		t.Fatalf("Explain failed: %v", err)
	}
}