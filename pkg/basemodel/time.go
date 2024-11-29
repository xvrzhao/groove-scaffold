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

// Scan 在 gorm 进行读操作将数据库时间值转换为 Time 结构体字段时用到，
// 如果 Time 没有此方法，则会报错：无法将 time.Time 类型赋值到 Time 类型。
func (t *Time) Scan(v interface{}) error {
	if tt, ok := v.(time.Time); ok {
		*t = Time{tt}
		return nil
	}
	return fmt.Errorf("can not convert %v to Time", v)
}

// Value 在 gorm 进行写操作将 Time 结构体字段转为数据库时间值时用到，
// 如果 Time 没有此方法，则会报错：不支持 Time 类型。
func (t Time) Value() (driver.Value, error) {
	var zTime time.Time
	if t.UnixNano() == zTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}
