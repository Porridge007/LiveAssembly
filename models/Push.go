package models

type PushWatch struct {
	Rtmp string
	Flv string
	Hls string
}

type PushStruct struct {
	PushAddr string
	WatchAddr PushWatch
}

type PushList []PushStruct
