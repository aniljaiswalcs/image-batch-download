package app

import (
	"context"
	"io"
	"net/http"
	"os"
	"time"

	cfg "github.com/aniljaiswalcs/image-batch-download/config"
	"github.com/aniljaiswalcs/image-batch-download/core/fixture"
	"github.com/aniljaiswalcs/image-batch-download/core/imagedownloader"
	"github.com/aniljaiswalcs/image-batch-download/core/util"
	imageDownloaderPkg "github.com/aniljaiswalcs/image-batch-download/pkg/imagedownloader"
	"github.com/aniljaiswalcs/image-batch-download/pkg/logger"
	"github.com/oklog/ulid/v2"
)

const (
	unlimited = 0
)

func StartImageDownloaderApp(ctx context.Context, config *cfg.Config) error {
	logger.Init()

	imageDownloader := &imagedownloader.ImageDownloader{
		FixtureLoader: &fixture.Fixture{

			Path:      config.ImageUrlPath,
			BatchSize: int(config.BatchSize),
		},
		DownloaderClient: &imageDownloaderPkg.Client{
			HTTPClient: &imageDownloaderPkg.HTTPClient{
				BaseClient: &http.Client{
					Transport: &http.Transport{
						MaxIdleConns:        250,
						MaxIdleConnsPerHost: 25,
						MaxConnsPerHost:     unlimited,
						IdleConnTimeout:     unlimited,
					},
					Timeout: time.Duration(60) * time.Second,
				},
				RetryOption: imageDownloaderPkg.RetryOption{
					BaseDelay:   time.Duration(50) * time.Millisecond,
					MaxDelay:    time.Duration(3) * time.Second,
					MaxAttempts: 3,
				},
				AcceptedImageContentTypeExtensions: imageDownloaderPkg.CommonImageContentTypeExtensions,
			},
			CreateFileFn: os.Create,
			CopyFileFn:   io.Copy,
		},
		UlidMakerFn:     ulid.Make,
		Workers:         10,
		StorageRootPath: config.BaseRootPath,
	}

	out, err := imageDownloader.DownloadAllImages(ctx)
	if err != nil {
		return err
	}

	return util.JsonStdout(out)
}
