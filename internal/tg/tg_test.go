package tg

import (
	"fmt"
	"strings"
	"testing"
)

func TestName(t *testing.T) {
	txt := `🔍 关键词：#纸钞屋 #西班牙 #动作 #悬疑 #1080p #Netflix

📢 发布频道：阿里云盘吧
🔗 侵删联系：Delete Contact / DMCA`
	fmt.Println(strings.LastIndex(txt, "发布频道："))
	fmt.Println(strings.LastIndex(txt, "侵删联系："))
}
