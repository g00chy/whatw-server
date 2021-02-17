package database

import (
	"database/sql/driver"
	"errors"
	"time"
)

// XXX: アプリでナノ秒は必要ないので秒まででフォーマット
//{
//    "createdAt": "2015-08-09T15:03:05.135166971Z",
//    "objectID": "55c112a40000a1"
//}
// See: http://qiita.com/taizo/items/2c3a338f1aeea86ce9e2
type ResourceTime struct {
	time.Time
}

func (rt ResourceTime) MarshalJSON() ([]byte, error) {
	t := rt.Time
	if y := t.Year(); y < 0 || y >= 10000 {
		return nil, errors.New("ResourceTime.MarshalJSON: year outside of range [0,9999]")
	}
	return []byte(t.Format(`"` + time.RFC3339 + `"`)), nil
}

func (rt ResourceTime) MarshalText() ([]byte, error) {
	t := rt.Time
	if y := t.Year(); y < 0 || y >= 10000 {
		return nil, errors.New("Time.MarshalText: year outside of range [0,9999]")
	}
	return []byte(t.Format(time.RFC3339)), nil
}

// Sql driver interface

func (rt *ResourceTime) Scan(value interface{}) error {
	rt.Time = value.(time.Time)
	return nil
}

func (rt ResourceTime) Value() (driver.Value, error) {
	return rt.Time, nil
}
