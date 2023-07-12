package consts

import "errors"

const SpanCTX = "span-ctx"

var (
	GetTrackIdErr = errors.New("获取 track id 错误")
)
