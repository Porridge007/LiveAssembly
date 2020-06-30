package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"liveassembly/models"
	"net/http"
	"os/exec"
	"strconv"
	"strings"
)

var PidMap = make(models.FFmpegPid)

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

	thisRoom := RoomID
	RoomID += 1

	return models.PushStruct{
		PushAddr: "rtmp://localhost:1935/live/" + channel.Channel,
		WatchAddr: models.PushWatch{
			Rtmp: "rtmp://localhost:1935/live/movie" + strconv.Itoa(thisRoom),
			Flv:  "http://127.0.0.1:7001/live/movie" + strconv.Itoa(thisRoom) + ".flv",
			Hls:  "http://127.0.0.1:7002/live/movie" + strconv.Itoa(thisRoom) + ".m3u8",
		},
	}
}

func PullStream(pullAddr, pushAddr, roomID string)  {
	cmd := exec.Command("extra/ffmpeg.exe", "-i", pullAddr, "-vcodec", "copy", "-acodec", "copy", "-f","flv",
		pushAddr)
	fmt.Println(cmd)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	cmd.Start()
	//pidStruct :=  models.FFmpegPid{
	//	Channel: roomID,
	//	Pid:     cmd.Process.Pid,
	//}
	PidMap[roomID] = cmd.Process.Pid
	fmt.Println(PidMap)
	cmd.Wait()

}

func GetChannel(watchAddr string) string{
	split := strings.Split(watchAddr, "/")
	return  split[len(split)-1]
}
