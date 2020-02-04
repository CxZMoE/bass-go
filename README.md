# bass-go
Un4seen Development Bass 音频库的Go绑定 

# 依赖
Linux `libbass`  
Windows `bass.dll`  
前往[官网](http://www.un4seen.com/bass.html)下载

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
