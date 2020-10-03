import {Inject, Injectable, PLATFORM_ID} from '@angular/core';
import {isPlatformBrowser} from '@angular/common';
import axios, {AxiosRequestConfig, AxiosResponse} from 'axios';
import {StorageConst} from '../../const/storage.const';

@Injectable({
  providedIn: 'root'
})
export class ApiService {

  private DEFAULT_TIMEOUT = 5000;
  private TOKEN_KEY = StorageConst.KC_ACCESS_TOKEN;
  private readonly isBrowser: boolean;

  constructor(@Inject(PLATFORM_ID) platformId: any) {
    this.isBrowser = isPlatformBrowser(platformId);
    this.setDefaultRequestConfig();
  }

  /**
   * GET Method
   * @param url string
   * @param params any
   */
  public get(url: string, params?: any): Promise<AxiosResponse<any>> {
    const requestConfig: AxiosRequestConfig = {
      params
    };
    return axios.get(url, requestConfig);
  }

  /**
   * POST Method
   * @param url string
   * @param data any
   */
  public post(url: string, data: any): Promise<AxiosResponse<any>> {
    return axios.post(url, data);
  }

  /**
   * PUT Method
   * @param url string
   * @param data any
   */
  public put(url: string, data: any): Promise<AxiosResponse<any>> {
    return axios.put(url, data);
  }

  /**
   * PATCH Method
   * @param url
   * @param data
   */
  public patch(url: string, data: any): Promise<AxiosResponse<any>> {
    return axios.patch(url, data);
  }

  /**
   * DELETE Method
   * @param url
   * @param params
   */
  public delete(url: string, params?: any): Promise<AxiosResponse<any>> {
    const requestConfig: AxiosRequestConfig = {
      params
    };
    return axios.delete(url, requestConfig);
  }

  /**
   * Set default request config axios
   * @private
   */
  private setDefaultRequestConfig(): void {
    axios.defaults.timeout = this.DEFAULT_TIMEOUT;
    axios.defaults.responseType = 'json';
    axios.defaults.timeoutErrorMessage = 'Cannot connect to server';
    axios.defaults.headers.common.Authorization = `Bearer ${this.getToken()}`;
  }

  /**
   * Get Token
   * @private
   */
  private getToken(): string {
    if (this.isBrowser) {
      const token = localStorage.getItem(this.TOKEN_KEY);
      return token && token.length > 0 ? token : '';
    }

    return '';
  }
}
