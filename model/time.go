package model

import (
    "fmt"
    "log"
    "regexp"
    "strconv"
    "time"
)

const (
    nsToMs = int64(1e6)
)

var (
    parsePattern *regexp.Regexp
)

func init() {
    /* The parsePattern does the three timestamp formats the API supports,
       and walking the subgroups (skipping the first) and excluding
       all zero-length entries will give you, in order, strings
       for: year, mo, day, hour, min, sec, millis.

       Millis can be blank.  */
    parsePattern = regexp.MustCompile(`\A(\d\d\d\d)(?:(?:(\d\d)(\d\d)(\d\d)(\d\d)(\d\d)(?:\.(\d{1,3}))?)|(?:-(\d\d)-(\d\d)(?:(?:T(\d\d)\:(\d\d):(\d\d)(?:\.(\d{1,3}))?Z)|(?: (\d\d)\:(\d\d)\:(\d\d)(?:\.(\d{1,3}))?))))\z`)
}


// Timestamps are represented as milliseconds since the epoch
type Timestamp int64

func (t Timestamp) Time() time.Time {
    return time.Unix(int64(t) / 1000, (int64(t) % 1000) * 1e6)
}


func ToTimestamp(t time.Time) Timestamp {
    return Timestamp(t.UnixNano() / nsToMs)
}

func ParseTimestamp(stamp string) Timestamp {
    matches := parsePattern.FindStringSubmatch(stamp)[1:]
    fields := make([]int, 7)
    log.Print("Got matches: ", matches)
    var i, j int
    for i, j = 0, 0; i < len(matches) && j < len(fields); i++ {
        if matches[i] != "" {
            fields[j], _ = strconv.Atoi(matches[i])
            j++
        }
    }
    if j < 6 {
        panic(fmt.Sprintf("Invalid timestamp: %s", stamp))
    }
    ts := time.Date(fields[0], time.Month(fields[1]), fields[2], fields[3], fields[4], fields[5], fields[6] * int(nsToMs), UTC)
    return ToTimestamp(ts)
}
