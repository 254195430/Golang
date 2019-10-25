package main

type Time struct {
	// sec 给出自公元 1 年1 月1 日00:00:00
	// 开始的秒数
	sec int64

	// nsec 指定了一秒内的纳秒偏移，
	// 这个值是非零值，
	// 必须在[0, 999999999]范围内
	nsec int32

	// loc 指定了一个Location，
	// 用于决定该时间对应的当地的分、小时、
	// 天和年的值
	// 只有Time 的零值，其loc 的值是nil
	// 这种情况下，认为处于UTC 时区
	loc *Location
}

func Now() Time {
	sec, nsec := now()
	return Time{sec + unixToInternal, nsec, Local}
}
