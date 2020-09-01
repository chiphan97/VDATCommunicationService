import {Inject, Injectable, PLATFORM_ID} from '@angular/core';
import axios, {AxiosRequestConfig, AxiosResponse} from 'axios';
import {isPlatformBrowser} from '@angular/common';

@Injectable({
  providedIn: 'root'
})
export class ApiService {

  private DEFAULT_TIMEOUT = 5000;
  private TOKEN_KEY = 'TOKEN';
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
      url,
      method: 'GET',
      params
    };
    return axios.request(requestConfig);
  }

  /**
   * POST Method
   * @param url string
   * @param data any
   */
  public post(url: string, data: any): Promise<AxiosResponse<any>> {
    const requestConfig: AxiosRequestConfig = {
      url,
      method: 'POST',
      data
    };
    return axios.request(requestConfig);
  }

  /**
   * PUT Method
   * @param url string
   * @param data any
   */
  public put(url: string, data: any): Promise<AxiosResponse<any>> {
    const requestConfig: AxiosRequestConfig = {
      url,
      method: 'PUT',
      data
    };
    return axios.request(requestConfig);
  }

  /**
   * PATCH Method
   * @param url
   * @param data
   */
  public patch(url: string, data: any): Promise<AxiosResponse<any>> {
    const requestConfig: AxiosRequestConfig = {
      url,
      method: 'PATCH',
      data
    };
    return axios.request(requestConfig);
  }

  /**
   * DELETE Method
   * @param url
   */
  public delete(url: string): Promise<AxiosResponse<any>> {
    const requestConfig: AxiosRequestConfig = {
      url,
      method: 'DELETE'
    };
    return axios.request(requestConfig);
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
