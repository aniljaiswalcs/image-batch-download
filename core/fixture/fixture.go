package fixture

import (
	"bufio"
	"context"
	"os"
	"strings"
)

type Fixture struct {
	Path      string
	BatchSize int
}

func (f *Fixture) LoadExecute(_ context.Context, batchExecutor func(urls []string) error) error {
	file, err := os.Open(f.Path)
	if err != nil {
		return err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	urls := make([]string, 0, f.BatchSize)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		urlTest := scanner.Text()
		url := strings.Split(urlTest, " ")
		if urlTest == "" {
			continue
		}

		urls = append(urls, url...)
		if len(urls) == 0 {
			return nil
		}
	}
	return batchExecutor(urls)
}
