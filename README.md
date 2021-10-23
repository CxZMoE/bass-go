# bass-go
<h3>这个库被归档，并且不会再更新，最终日期2021/10/23测试可用</h3>
Un4seen Development Bass 音频库的Go绑定  
任何`收费用途`使用本库需要在[UN4Seen BASS](http://www.un4seen.com/)购买证书。  
`未购买证书情况下使用本绑定库造成的任何纠纷本人概不负责`

BASS is free for non-commercial use. 

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
