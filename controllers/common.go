package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"liveassembly/models"
	"net/http"
	"strconv"
)

func CreateChannel() models.PushStruct {
	url := "http://localhost:8090/control/get?room=movie" + strconv.Itoa(RoomID)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println()
	}
	var channel Channel
	json.Unmarshal(body, &channel)

	return models.PushStruct{
		PushAddr: "rtmp://localhost:1935/live/" + channel.Channel,
		WatchAddr: models.PushWatch{
			Rtmp: "rtmp://localhost:1935/live/movie" + strconv.Itoa(RoomID),
			Flv:  "http://127.0.0.1:7001/live/movie" + strconv.Itoa(RoomID) + ".flv",
			Hls:  "http://127.0.0.1:7002/live/movie" + strconv.Itoa(RoomID) + ".m3u8",
		},
	}
}
