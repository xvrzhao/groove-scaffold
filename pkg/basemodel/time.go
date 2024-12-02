package basemodel

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"time"
)

type Time struct {
	time.Time
}

func (t Time) MarshalJSON() (data []byte, err error) {
	data = []byte(strconv.Itoa(int(t.UnixMilli())))
	return
}

func (t *Time) Scan(v interface{}) error {
	if tt, ok := v.(time.Time); ok {
		*t = Time{tt}
		return nil
	}
	return fmt.Errorf("can not convert %v to Time", v)
}

func (t Time) Value() (driver.Value, error) {
	var zTime time.Time
	if t.UnixNano() == zTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}
