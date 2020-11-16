import {Injectable} from '@angular/core';
import * as Minio from '../../../../node_modules/minio/dist/main/minio-browser';


@Injectable({
    providedIn: 'root'
})

export class MinioService {

    private readonly vdatBucketName = 'vdat.s3.bucket';
    private minioClient: Minio.Client;

    constructor(){
        console.log('Minio service about to be created');
        this.minioClient = new Minio.Client({
            endPoint: 'localhost',
            port: 9000,
            useSSL: false,
            accessKey: 'minio',
            secretKey: 'minio123'
        });
        console.log('new minio client created');
        this.minioClient.bucketExists(this.vdatBucketName, function( err, exists){
            if (err) {
                console.log('error in bucket')
                return console.error(err)
              }
              if (exists) {
                return console.log('Bucket exists.')
              } else {
                this.minioClient.makeBucket(this.vdatBucketName, 'us-east-1', function(err) {
                    if (err) return console.log(err);
                    console.log('Bucket created successfully in "us-east-1".')
                })
              }
        })
    }

    public uploadFile(file: string) {
        var metaData = {
            'Content-Type': 'application/octet-stream',
            'X-Amz-Meta-Testing': 1234,
            'example': 5678
        }
        this.minioClient.fPutObject(this.vdatBucketName, 'new-File.png', file, 'application/octet-stream', function(err, etag) {
            if (err) {console.log(err)
            }
            console.log('File uploaded successfully .')
          });
    }
}
