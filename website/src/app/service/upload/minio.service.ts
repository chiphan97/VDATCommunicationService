import {Injectable} from '@angular/core';
import { Minio, Client } from '../../../../node_modules/minio';


@Injectable({
    providedIn: 'root'
})

export class MinioService {

    private minioClient: Client;

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
