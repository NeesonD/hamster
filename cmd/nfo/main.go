package main

import (
	"flag"
	"fmt"
	"hamster/internal/aliyunpan"
	"hamster/internal/nfo"
	"strings"
	"time"
)

var (
	aliFileDepth    = flag.Int("share_file_depth", 3, "")
	aliRefreshToken = flag.String("refresh_token", "", "")
	filterFile      = flag.String("filter_file", "", "")
)

func main() {
	now := time.Now()
	flag.Parse()
	aliyunpan.InitALiYunClient(*aliRefreshToken)

	split := strings.Split(*filterFile, ",")
	filterFileMap := map[string]struct{}{}
	for _, s := range split {
		filterFileMap[s] = struct{}{}
	}

	aliyunpan.SyncAliData(*aliFileDepth, filterFileMap)
	fmt.Printf("耗时：%f \n", time.Since(now).Seconds())
	nfo.AppendLink("", "")
}
