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

## 播放本地文件
``` golang
package main


import (
	"fmt"
	bass "github.com/bass-go"
)

func main() {
	err := bass.Init()
	if err == bass.BASS_ERROR_INIT {
		panic("Failed to init bass")
	}
	handle := bass.StreamCreateFile(0, "./KING.mp3", 0, 0)
	bass.ChannelPlay(handle, 0)
	fmt.Scanln()
}
```
