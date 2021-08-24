module github.com/northwesternmutual/grammes

go 1.13

require (
	github.com/google/uuid v1.1.0
	github.com/gopherjs/gopherjs v0.0.0-20190309154008-847fc94819f9 // indirect
	github.com/gorilla/websocket v1.4.2
	github.com/pkg/errors v0.9.1 // indirect
	github.com/smartystreets/assertions v0.0.0-20190215210624-980c5ac6f3ac // indirect
	github.com/smartystreets/goconvey v0.0.0-20190306220146-200a235640ff
	github.com/stretchr/testify v1.6.1 // indirect
	go.uber.org/atomic v1.3.2 // indirect
	go.uber.org/multierr v1.1.0 // indirect
	go.uber.org/zap v1.9.1
	golang.org/x/sync v0.0.0-20201207232520-09787c993a3a
)

replace github.com/gorilla/websocket v1.4.2 => github.com/wiz-sec/websocket v1.4.3-0.20210824141715-858e166c8e75
