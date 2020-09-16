import { Injectable } from '@angular/core';
import {TranslateService} from '@ngx-translate/core';

@Injectable({
  providedIn: 'root'
})
export class LanguageService {
  private currentLanguage: string;

  constructor(private translateService: TranslateService) {
    this.currentLanguage = 'vi';
  }

  /**
   * Lấy coded ngôn ngữ hiện tại
   */
  getCurrentLanguage() {
    return this.currentLanguage;
  }

  /**
   * Dịch chuỗi sang ngôn ngữ hiện tại
   * @param text
   */
  async translation(text: string): Promise<string> {
    return await this.translateService.get(text).toPromise();
  }

  /**
   * Khởi tạo ngôn ngữ mặc đinh
   */
  setDefaultLanguage() {
    this.translateService.setDefaultLang(this.currentLanguage);
  }

  /**
   * Thay đổi ngôn ngữ hiện tại
   * @param language
   */
  switchLanguage(language: string) {
    this.currentLanguage = language;
    this.translateService.use(this.currentLanguage);
  }
}
