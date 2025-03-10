// Code generated by tmpl; DO NOT EDIT.
// https://github.com/benbjohnson/tmpl
//
// Source: engine.gen.go.tmpl

package tsm1

import (
	"context"

	"github.com/cnosdb/db/query"
)

// buildFloatCursor creates a cursor for a float field.
func (e *Engine) buildFloatCursor(ctx context.Context, measurement, seriesKey, field string, opt query.IteratorOptions) floatCursor {
	key := SeriesFieldKeyBytes(seriesKey, field)
	cacheValues := e.Cache.Values(key)
	keyCursor := e.KeyCursor(ctx, key, opt.SeekTime(), opt.Ascending)
	return newFloatCursor(opt.SeekTime(), opt.Ascending, cacheValues, keyCursor)
}

// buildIntegerCursor creates a cursor for a integer field.
func (e *Engine) buildIntegerCursor(ctx context.Context, measurement, seriesKey, field string, opt query.IteratorOptions) integerCursor {
	key := SeriesFieldKeyBytes(seriesKey, field)
	cacheValues := e.Cache.Values(key)
	keyCursor := e.KeyCursor(ctx, key, opt.SeekTime(), opt.Ascending)
	return newIntegerCursor(opt.SeekTime(), opt.Ascending, cacheValues, keyCursor)
}

// buildUnsignedCursor creates a cursor for a unsigned field.
func (e *Engine) buildUnsignedCursor(ctx context.Context, measurement, seriesKey, field string, opt query.IteratorOptions) unsignedCursor {
	key := SeriesFieldKeyBytes(seriesKey, field)
	cacheValues := e.Cache.Values(key)
	keyCursor := e.KeyCursor(ctx, key, opt.SeekTime(), opt.Ascending)
	return newUnsignedCursor(opt.SeekTime(), opt.Ascending, cacheValues, keyCursor)
}

// buildStringCursor creates a cursor for a string field.
func (e *Engine) buildStringCursor(ctx context.Context, measurement, seriesKey, field string, opt query.IteratorOptions) stringCursor {
	key := SeriesFieldKeyBytes(seriesKey, field)
	cacheValues := e.Cache.Values(key)
	keyCursor := e.KeyCursor(ctx, key, opt.SeekTime(), opt.Ascending)
	return newStringCursor(opt.SeekTime(), opt.Ascending, cacheValues, keyCursor)
}

// buildBooleanCursor creates a cursor for a boolean field.
func (e *Engine) buildBooleanCursor(ctx context.Context, measurement, seriesKey, field string, opt query.IteratorOptions) booleanCursor {
	key := SeriesFieldKeyBytes(seriesKey, field)
	cacheValues := e.Cache.Values(key)
	keyCursor := e.KeyCursor(ctx, key, opt.SeekTime(), opt.Ascending)
	return newBooleanCursor(opt.SeekTime(), opt.Ascending, cacheValues, keyCursor)
}
