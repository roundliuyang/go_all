package adapter

import (
	"fmt"
	"testing"
)

// 旧的文档存储系统
type OldStorage struct{}

func (o *OldStorage) SaveToFile(data string) {
	fmt.Println("Saving data to file:", data)
}

// 新的云端 API
type CloudAPI struct{}

func (c *CloudAPI) UploadToCloud(data string) {
	fmt.Println("Uploading data to cloud:", data)
}

// 适配器
type CloudAdapter struct {
	cloud *CloudAPI
}

func (ca *CloudAdapter) SaveToFile(data string) {
	ca.cloud.UploadToCloud(data)
}

func TestDemo1(t *testing.T) {
	oldStorage := &OldStorage{}
	oldStorage.SaveToFile("Hello from old system!")

	adapter := &CloudAdapter{cloud: &CloudAPI{}}
	adapter.SaveToFile("Hello from new system via adapter!")
}

// 有些音樂播放器只能播放MP3，有些只能播放WAV，如果我們想要一個通用的播放器怎麼辦？
// 一般的 MP3 播放器
type MP3Player struct{}

func (m *MP3Player) PlayMP3(data string) {
	fmt.Println("Playing MP3:", data)
}

// 专门的 WAV 播放器
type WAVPlayer struct{}

func (w *WAVPlayer) PlayWAV(data string) {
	fmt.Println("Playing WAV:", data)
}

// 适配器
type UniversalPlayer struct {
	mp3 *MP3Player
	wav *WAVPlayer
}

func (u *UniversalPlayer) Play(data, fileType string) {
	if fileType == "mp3" {
		u.mp3.PlayMP3(data)
	} else if fileType == "wav" {
		u.wav.PlayWAV(data)
	}
}

func TestDemo2(t *testing.T) {
	player := &UniversalPlayer{mp3: &MP3Player{}, wav: &WAVPlayer{}}
	player.Play("Rock Song", "mp3")
	player.Play("Classical Tune", "wav")
}
