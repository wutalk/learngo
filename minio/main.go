package main

/*
before run this code, first start minio server like this:
docker run -p 9000:9000 --env MINIO_ACCESS_KEY="minio" --env MINIO_SECRET_KEY="miniostorage" metadata-inventory_neo-minio server /tmp/data

more env can be found here:
https://hub.docker.com/layers/minio/minio/latest/images/sha256-d3fcf7345bba4bea545c010707fcaa92f4ae1475a3baffde77fa1d85a8135de4?context=explore
*/

func main() {
	/*
			ctx := context.Background()
			endpoint := "127.0.0.1:9000"
			accessKeyID := "minio"
			secretAccessKey := "miniostorage"
			useSSL := false

			// Initialize minio client object.
			minioClient, err := minio.New(endpoint, &minio.Options{
				Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
				Secure: useSSL,
			})
			if err != nil {
				log.Fatalln(err)
			}

			bucketName := "fp.content"
			// createBucketAndUploadFile(ctx, minioClient)
			// objectCh := minioClient.ListObjects(ctx, "mymusic", minio.ListObjectsOptions{Recursive: true, Prefix: "gopls"})
			objectCh := minioClient.ListObjects(ctx, bucketName, minio.ListObjectsOptions{Recursive: true, Prefix: "LTE19A"})
			count := 0
			for object := range objectCh {
				if object.Err != nil {
					fmt.Println(object.Err)
					return
				}
				count++
				fmt.Printf("objects #%d: %s\n", count, object.Key)
				// prettyPrint(object)
			}
			fmt.Println("finish list objects")

			objectName := "LTE19A/index_20190508T132515.json"
			// objectName := "LTE19A/20190508T132515/files/BOM/FastPass/LTE19A/glbfs/20190508.132515/LTE19A.json"
			err = minioClient.FGetObject(ctx, bucketName, objectName, "/home/owen/testdata/LTE19A_20190508T132515/lte.json", minio.GetObjectOptions{})
			if err != nil {
				fmt.Errorf("fail to find object %#v\n", err)
			}
			fmt.Println("downloaded index file")
			// prettyPrint(obj)

			minioClient.SelectObjectContent(ctx, bucketName, objectName, minio.SelectObjectOptions{})
		}

		func prettyPrint(object interface{}) {
			// MarshalIndent
			objectJSON, err := json.MarshalIndent(object, "", "  ")
			if err != nil {
				log.Fatalf(err.Error())
			}
			fmt.Printf("MarshalIndent funnction output %s\n", string(objectJSON))
		}

		func createBucketAndUploadFile(ctx context.Context, minioClient *minio.Client) {
			// Make a new bucket called mymusic.
			bucketName := "mymusic"
			location := "us-east-1"

			err := minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
			if err != nil {
				// Check to see if we already own this bucket (which happens if you run this twice)
				exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
				if errBucketExists == nil && exists {
					log.Printf("We already own %s\n", bucketName)
				} else {
					log.Fatalln(err)
				}
			} else {
				log.Printf("Successfully created %s\n", bucketName)
			}

			// Upload the zip file
			objectName := "gopls.508022-1GiB-nonames.zip"
			filePath := "/tmp/gopls.508022-1GiB-nonames.zip"
			contentType := "application/zip"

			// Upload the zip file with FPutObject
			n, err := minioClient.FPutObject(ctx, bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
			if err != nil {
				log.Fatalln(err)
			}

			log.Printf("Successfully uploaded %s of size %d\n", objectName, n.Size)
	*/
}
