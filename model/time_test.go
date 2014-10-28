package model

import (
    "fmt"
    "testing"
    "time"
)

var (
    UTC *time.Location
    AprilFool time.Time
    AprilFoolMillis int64
)

func init() {
    UTC, _ = time.LoadLocation("UTC")
    AprilFool = time.Date(2014, 4, 1, 0, 0, 0, 0, UTC)
    AprilFoolMillis = 1396310400000
}

func TestToTimestampNoNanosecs(t *testing.T) {
    ts := ToTimestamp(AprilFool)
    if int64(ts) != AprilFoolMillis {
        t.Error(fmt.Sprintf("Timestamp (%d) does not match expected value (%d).", int64(ts), AprilFoolMillis))
    }
}

func TestToTimestampWithNanosecs(t * testing.T) {
    millis := int64(897)
    nanos := int64(123456)

    tm := AprilFool.Add(time.Millisecond * time.Duration(millis)).Add(time.Nanosecond * time.Duration(nanos))
    ts := ToTimestamp(tm)
    exp := AprilFoolMillis + millis
    if int64(ts) != exp {
        t.Error(fmt.Sprintf("Timestamp (%d) does not match expected value (%d)", int64(ts), exp))
    }
}

func AssertExpectedTime(t *testing.T, exp int64, got Timestamp) {
    if got != Timestamp(exp) {
        t.Error(fmt.Sprintf("Timestamp (%d) does not match expected value (%d)", int64(got), exp))
    }
}

func TestFromShortIsoStampNoMillis(t *testing.T) {
    ts := ParseTimestamp("2014-04-01T00:00:00Z")
    AssertExpectedTime(t, AprilFoolMillis, ts)
}

func TestFromShortIsoStampMillis(t *testing.T) {
    ts := ParseTimestamp("2014-04-01T00:00:00.567Z")
    AssertExpectedTime(t, AprilFoolMillis + int64(567), ts)
}

func TestFromSimpleStampNoMillis(t *testing.T) {
    ts := ParseTimestamp("2014-04-01 00:00:02")
    AssertExpectedTime(t, AprilFoolMillis + int64(2000), ts)
}

func TestFromSimpleStampMillis(t *testing.T) {
    ts := ParseTimestamp("2014-04-01 00:00:00.931")
    AssertExpectedTime(t, AprilFoolMillis + int64(931), ts)
}

func TestFromNumericStampNoMillis(t *testing.T) {
    ts := ParseTimestamp("20140401000000")
    AssertExpectedTime(t, AprilFoolMillis, ts)
}

func TestFromNumericStampMillis(t *testing.T) {
    ts := ParseTimestamp("20140401000000.173")
    AssertExpectedTime(t, AprilFoolMillis + int64(173), ts)
}


