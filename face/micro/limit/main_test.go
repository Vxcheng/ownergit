package main

import (
    "sync/atomic"
    "testing"
    "time"
)

// ensure basic behavior of FixedWindowLimiter
func TestFixedWindowLimiter_Allow(t *testing.T) {
    lim := NewFixedWindowLimiter(2)

    // two requests should pass
    if !lim.Allow() {
        t.Error("first allow should be true")
    }
    if !lim.Allow() {
        t.Error("second allow should be true")
    }

    // third request should be blocked
    if lim.Allow() {
        t.Error("third allow should be false")
    }

    // simulate crossing into next second by moving lastSecond back
    prev := atomic.LoadInt64(&lim.lastSecond)
    atomic.StoreInt64(&lim.lastSecond, prev-1)

    if !lim.Allow() {
        t.Error("after second rollover allow should be true")
    }
}

func TestTokenBucketLimiter_AllowAndRefill(t *testing.T) {
    tb := NewTokenBucketLimiter(1, 2) // rate 1 token/sec, capacity 2
    defer tb.Stop()

    // initially capacity tokens
    if !tb.Allow() {
        t.Error("first token should be available")
    }
    if !tb.Allow() {
        t.Error("second token should be available")
    }
    if tb.Allow() {
        t.Error("third token should NOT be available")
    }

    // force refill by rewinding lastFill and calling fill manually
    tb.lastFill = tb.lastFill.Add(-2 * time.Second)
    tb.fill()

    if !tb.Allow() {
        t.Error("token should be available after refill")
    }
}

