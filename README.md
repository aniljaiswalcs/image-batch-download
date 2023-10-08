
## Description
1. The ImageDownloaderApp initiates and triggers the ImageDownloaderService.
2. The ImageDownloaderService retrieves batched image URLs from the FixtureLoaderExecutor.
3. The ImageDownloaderService then distributes these batched image URLs among its worker pool.
4. Each worker within the ImageDownloaderService processes its assigned batch of image URLs.
5. Utilizing the ImageDownloaderClient, each ImageDownloaderWorker efficiently downloads the specified image URLs.
6. The downloaded images are stored on the local disk by the ImageDownloaderClient.
7. The ImageDownloaderApp generates a report on STDOUT, detailing downloaded images, unavailable images, skipped images, and more.

# Key Strengths of This Solution
Several underlying implementations set this solution apart:

1. The ImageDownloaderService employs a worker pool comprising 10 workers, with each worker capable of executing 25 concurrent download operations. This parallel approach accelerates image downloading.
2. Through connection pooling, the ImageDownloaderClient optimizes HTTP requests by avoiding the overhead of establishing new connections for each call. This results in significantly reduced latency.
3. The ImageDownloaderClient incorporates an HTTP retry mechanism, allowing failed calls to be retried up to three times, enhancing the solution's robustness.
4. A strategic exponential backoff strategy is applied to the retry mechanism in the ImageDownloaderClient, contributing to improved reliability in the face of connectivity challenges.
5. Utilizing HTTP timeouts, the ImageDownloaderClient prevents application hang-ups due to unexpectedly prolonged tasks, enhancing overall responsiveness.
6. Image IDs are generated using ULID, ensuring that images with the same name do not overwrite each other, thus maintaining data integrity.
7. Instead of relying solely on image extensions in URLs, the ImageDownloaderClient identifies image types based on the content type header. This versatile approach ensures accurate identification irrespective of URL structures.

## Running Tests

To run the tests, execute the following command in your terminal:
```
make test
```

## Run the app
### Prerequisite

Golang 1.20

### Building the app
Before running the app. Please build it first. Execute this command in the terminal:
```bash
make run
```
### Running with Custom Input Fixture

Config.yml make customization:

file with URL path defined with imageurlpath
batchsize also configurale.
Downloaded image same path must be defined at storagerootpath

```bash
imageurlpath: "external-files/images.txt"
batchsize: 25 
storagerootpath: "/go-workspace/src//getsafe/image-batch-download"
```
