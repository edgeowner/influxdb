// Generated by tmpl
// https://github.com/benbjohnson/tmpl
//
// DO NOT EDIT!
// Source: array_cursor_test.gen.go.tmpl

package reads

import (
	"context"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/influxdata/flux/execute"
	"github.com/influxdata/flux/values"
	"github.com/influxdata/influxdb/v2/storage/reads/datatypes"
	"github.com/influxdata/influxdb/v2/tsdb/cursors"
)

type MockFloatArrayCursor struct {
	CloseFunc func()
	ErrFunc   func() error
	StatsFunc func() cursors.CursorStats
	NextFunc  func() *cursors.FloatArray
}

func (c *MockFloatArrayCursor) Close()                     { c.CloseFunc() }
func (c *MockFloatArrayCursor) Err() error                 { return c.ErrFunc() }
func (c *MockFloatArrayCursor) Stats() cursors.CursorStats { return c.StatsFunc() }
func (c *MockFloatArrayCursor) Next() *cursors.FloatArray  { return c.NextFunc() }

func TestNewAggregateArrayCursor_Float(t *testing.T) {
	t.Run("Mean", func(t *testing.T) {
		want := &floatWindowMeanArrayCursor{
			FloatArrayCursor: &MockFloatArrayCursor{},
			res:              cursors.NewFloatArrayLen(1),
			tmp:              &cursors.FloatArray{},
		}

		agg := &datatypes.Aggregate{
			Type: datatypes.AggregateTypeMean,
		}

		got, _ := newAggregateArrayCursor(context.Background(), agg, &MockFloatArrayCursor{})

		if diff := cmp.Diff(got, want, cmp.AllowUnexported(floatWindowMeanArrayCursor{})); diff != "" {
			t.Fatalf("did not get expected cursor; -got/+want:\n%v", diff)
		}
	})
}

func TestNewWindowAggregateArrayCursorMonths_Float(t *testing.T) {
	t.Run("Mean", func(t *testing.T) {
		window := execute.Window{
			Every:  values.MakeDuration(int64(time.Hour), 0, false),
			Period: values.MakeDuration(int64(time.Hour), 0, false),
		}

		want := &floatWindowMeanArrayCursor{
			FloatArrayCursor: &MockFloatArrayCursor{},
			res:              cursors.NewFloatArrayLen(MaxPointsPerBlock),
			tmp:              &cursors.FloatArray{},
			window:           window,
		}

		agg := &datatypes.Aggregate{
			Type: datatypes.AggregateTypeMean,
		}

		got, _ := newWindowAggregateArrayCursor(context.Background(), agg, window, &MockFloatArrayCursor{})

		if diff := cmp.Diff(got, want, cmp.AllowUnexported(floatWindowMeanArrayCursor{})); diff != "" {
			t.Fatalf("did not get expected cursor; -got/+want:\n%v", diff)
		}
	})
}

func TestNewWindowAggregateArrayCursor_Float(t *testing.T) {
	t.Run("Mean", func(t *testing.T) {
		window := execute.Window{
			Every:  values.MakeDuration(0, 1, false),
			Period: values.MakeDuration(0, 1, false),
		}

		want := &floatWindowMeanArrayCursor{
			FloatArrayCursor: &MockFloatArrayCursor{},
			res:              cursors.NewFloatArrayLen(MaxPointsPerBlock),
			tmp:              &cursors.FloatArray{},
			window:           window,
		}

		agg := &datatypes.Aggregate{
			Type: datatypes.AggregateTypeMean,
		}

		got, _ := newWindowAggregateArrayCursor(context.Background(), agg, window, &MockFloatArrayCursor{})

		if diff := cmp.Diff(got, want, cmp.AllowUnexported(floatWindowMeanArrayCursor{})); diff != "" {
			t.Fatalf("did not get expected cursor; -got/+want:\n%v", diff)
		}
	})

}

type MockIntegerArrayCursor struct {
	CloseFunc func()
	ErrFunc   func() error
	StatsFunc func() cursors.CursorStats
	NextFunc  func() *cursors.IntegerArray
}

func (c *MockIntegerArrayCursor) Close()                      { c.CloseFunc() }
func (c *MockIntegerArrayCursor) Err() error                  { return c.ErrFunc() }
func (c *MockIntegerArrayCursor) Stats() cursors.CursorStats  { return c.StatsFunc() }
func (c *MockIntegerArrayCursor) Next() *cursors.IntegerArray { return c.NextFunc() }

func TestNewAggregateArrayCursor_Integer(t *testing.T) {
	t.Run("Mean", func(t *testing.T) {
		want := &integerWindowMeanArrayCursor{
			IntegerArrayCursor: &MockIntegerArrayCursor{},
			res:                cursors.NewFloatArrayLen(1),
			tmp:                &cursors.IntegerArray{},
		}

		agg := &datatypes.Aggregate{
			Type: datatypes.AggregateTypeMean,
		}

		got, _ := newAggregateArrayCursor(context.Background(), agg, &MockIntegerArrayCursor{})

		if diff := cmp.Diff(got, want, cmp.AllowUnexported(integerWindowMeanArrayCursor{})); diff != "" {
			t.Fatalf("did not get expected cursor; -got/+want:\n%v", diff)
		}
	})
}

func TestNewWindowAggregateArrayCursorMonths_Integer(t *testing.T) {
	t.Run("Mean", func(t *testing.T) {
		window := execute.Window{
			Every:  values.MakeDuration(int64(time.Hour), 0, false),
			Period: values.MakeDuration(int64(time.Hour), 0, false),
		}

		want := &integerWindowMeanArrayCursor{
			IntegerArrayCursor: &MockIntegerArrayCursor{},
			res:                cursors.NewFloatArrayLen(MaxPointsPerBlock),
			tmp:                &cursors.IntegerArray{},
			window:             window,
		}

		agg := &datatypes.Aggregate{
			Type: datatypes.AggregateTypeMean,
		}

		got, _ := newWindowAggregateArrayCursor(context.Background(), agg, window, &MockIntegerArrayCursor{})

		if diff := cmp.Diff(got, want, cmp.AllowUnexported(integerWindowMeanArrayCursor{})); diff != "" {
			t.Fatalf("did not get expected cursor; -got/+want:\n%v", diff)
		}
	})
}

func TestNewWindowAggregateArrayCursor_Integer(t *testing.T) {
	t.Run("Mean", func(t *testing.T) {
		window := execute.Window{
			Every:  values.MakeDuration(0, 1, false),
			Period: values.MakeDuration(0, 1, false),
		}

		want := &integerWindowMeanArrayCursor{
			IntegerArrayCursor: &MockIntegerArrayCursor{},
			res:                cursors.NewFloatArrayLen(MaxPointsPerBlock),
			tmp:                &cursors.IntegerArray{},
			window:             window,
		}

		agg := &datatypes.Aggregate{
			Type: datatypes.AggregateTypeMean,
		}

		got, _ := newWindowAggregateArrayCursor(context.Background(), agg, window, &MockIntegerArrayCursor{})

		if diff := cmp.Diff(got, want, cmp.AllowUnexported(integerWindowMeanArrayCursor{})); diff != "" {
			t.Fatalf("did not get expected cursor; -got/+want:\n%v", diff)
		}
	})

}

type MockUnsignedArrayCursor struct {
	CloseFunc func()
	ErrFunc   func() error
	StatsFunc func() cursors.CursorStats
	NextFunc  func() *cursors.UnsignedArray
}

func (c *MockUnsignedArrayCursor) Close()                       { c.CloseFunc() }
func (c *MockUnsignedArrayCursor) Err() error                   { return c.ErrFunc() }
func (c *MockUnsignedArrayCursor) Stats() cursors.CursorStats   { return c.StatsFunc() }
func (c *MockUnsignedArrayCursor) Next() *cursors.UnsignedArray { return c.NextFunc() }

func TestNewAggregateArrayCursor_Unsigned(t *testing.T) {
	t.Run("Mean", func(t *testing.T) {
		want := &unsignedWindowMeanArrayCursor{
			UnsignedArrayCursor: &MockUnsignedArrayCursor{},
			res:                 cursors.NewFloatArrayLen(1),
			tmp:                 &cursors.UnsignedArray{},
		}

		agg := &datatypes.Aggregate{
			Type: datatypes.AggregateTypeMean,
		}

		got, _ := newAggregateArrayCursor(context.Background(), agg, &MockUnsignedArrayCursor{})

		if diff := cmp.Diff(got, want, cmp.AllowUnexported(unsignedWindowMeanArrayCursor{})); diff != "" {
			t.Fatalf("did not get expected cursor; -got/+want:\n%v", diff)
		}
	})
}

func TestNewWindowAggregateArrayCursorMonths_Unsigned(t *testing.T) {
	t.Run("Mean", func(t *testing.T) {
		window := execute.Window{
			Every:  values.MakeDuration(int64(time.Hour), 0, false),
			Period: values.MakeDuration(int64(time.Hour), 0, false),
		}

		want := &unsignedWindowMeanArrayCursor{
			UnsignedArrayCursor: &MockUnsignedArrayCursor{},
			res:                 cursors.NewFloatArrayLen(MaxPointsPerBlock),
			tmp:                 &cursors.UnsignedArray{},
			window:              window,
		}

		agg := &datatypes.Aggregate{
			Type: datatypes.AggregateTypeMean,
		}

		got, _ := newWindowAggregateArrayCursor(context.Background(), agg, window, &MockUnsignedArrayCursor{})

		if diff := cmp.Diff(got, want, cmp.AllowUnexported(unsignedWindowMeanArrayCursor{})); diff != "" {
			t.Fatalf("did not get expected cursor; -got/+want:\n%v", diff)
		}
	})
}

func TestNewWindowAggregateArrayCursor_Unsigned(t *testing.T) {
	t.Run("Mean", func(t *testing.T) {
		window := execute.Window{
			Every:  values.MakeDuration(0, 1, false),
			Period: values.MakeDuration(0, 1, false),
		}

		want := &unsignedWindowMeanArrayCursor{
			UnsignedArrayCursor: &MockUnsignedArrayCursor{},
			res:                 cursors.NewFloatArrayLen(MaxPointsPerBlock),
			tmp:                 &cursors.UnsignedArray{},
			window:              window,
		}

		agg := &datatypes.Aggregate{
			Type: datatypes.AggregateTypeMean,
		}

		got, _ := newWindowAggregateArrayCursor(context.Background(), agg, window, &MockUnsignedArrayCursor{})

		if diff := cmp.Diff(got, want, cmp.AllowUnexported(unsignedWindowMeanArrayCursor{})); diff != "" {
			t.Fatalf("did not get expected cursor; -got/+want:\n%v", diff)
		}
	})

}
