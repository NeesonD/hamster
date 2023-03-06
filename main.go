package main

import (
	"hamster/config"
	"sync"
)

func init() {
	config.Load()
}

func main() {
	fileInfos := ReadFile(config.Get().DirPaths)
	var (
		wg     sync.WaitGroup
		worker = make(chan struct{}, config.Get().ScrapingWorker)
	)
	for _, info := range fileInfos {
		info := info
		wg.Add(1)
		worker <- struct{}{}
		go func() {
			defer func() {
				wg.Done()
				<-worker
			}()
			scraping(info)
		}()
	}
	wg.Wait()
}

func scraping(info *FileInfo) {

}

//
func getUnionId(info *FileInfo) string {
	switch config.Get().VideoType {
	case config.Movie:

	}
	return ""
}
