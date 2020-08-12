package lib

import "github.com/julianlee107/go-common/log"

const (
	TagUndefined = "_undef"
)

const (
	_tag         = "tag"
	_traceId     = "trace_id"
	_spanId      = "span_id"
	_childSpanId = "child_span_id"
	_tagPrefix   = "_com_"
)

type Logger struct {
}

var Log *Logger

type Trace struct {
	TraceId     string
	SpanId      string
	Caller      string
	SrcMethod   string
	HintCode    int64
	HintContent string
}

type TraceContext struct {
	Trace
}

func (l *Logger) Close() {
	log.Close()
}
