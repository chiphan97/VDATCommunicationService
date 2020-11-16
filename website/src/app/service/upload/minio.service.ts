import {Injectable} from '@angular/core';
//import { Client } from 'minio';
import * as Minio from '../../../../node_modules/minio/dist/main/minio-browser';


@Injectable({
    providedIn: 'root'
})

export class MinioService {

    private minioClient;

    constructor(){
        this.minioClient = new Minio.Client({
            endPoint: 'play.min.io',
            port: 9000,
            useSSL: true,
            accessKey: 'minio',
            secretKey: 'minio123'
        });
    }
}
