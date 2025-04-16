package datetime

import (
	"time"
)

type Mapper interface {
	FromString(rawDatetime string) (time.Time, error)
	ToString(datetime time.Time) string
}

type MapperImpl struct {
	format string
}

func NewMapper(format string) *MapperImpl {
	if format == "" {
		format = time.RFC3339
	}

	return &MapperImpl{
		format: format,
	}
}

func (m *MapperImpl) FromString(rawDatetime string) (time.Time, error) {
	return time.Parse(m.format, rawDatetime)
}

func (m *MapperImpl) ToString(datetime time.Time) string {
	return datetime.Format(m.format)
}
