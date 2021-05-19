package core

import "time"

const (
	defaultTimeLayout        = "15:04"
	default12HoursTimeLayout = "03:04PM"
)

type Formatter struct {
	use12Hours bool
}

func (f *Formatter) timeLayout() string {
	if f.use12Hours {
		return default12HoursTimeLayout
	}
	return defaultTimeLayout
}

func (f *Formatter) TimeString(input time.Time) string {
	return input.Format(f.timeLayout())
}

const (
	defaultRecordKeyLayout        = "2006-01-02-15-04"
	default12HoursRecordKeyLayout = "03:04PM"
)

func (f *Formatter) recordKeyLayout() string {
	if f.use12Hours {
		return defaultRecordKeyLayout
	}
	return default12HoursRecordKeyLayout
}

func (f *Formatter) ParseRecordKeyString(recordKey string) (time.Time, error) {
	return time.Parse(f.recordKeyLayout(), recordKey)
}
