# bass-go
A simple go binding of bass audio library


## Simply play a MP3 file from Url
```
package main


import (
	"fmt"
	bass "github.com/bass-go"
)

func main() {
	bass.Bass_Init()
	handle := bass.Bass_StreamCreateURL("http://music.163.com/song/media/outer/url?id=480097178.mp3", 0, 0, nil, nil)
	bass.Bass_ChannelPlay(handle, 0)
	fmt.Scanln()
}
```
