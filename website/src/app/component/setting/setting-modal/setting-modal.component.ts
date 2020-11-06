import { Component, OnInit } from '@angular/core';
import {SoundSetting} from '../../../model/sound-setting.model';
import AudioSettingData from '../../../data/audio-setting.data';
import {Setting} from '../../../model/setting.model';
import {NzModalRef, NzModalService} from 'ng-zorro-antd';

@Component({
  selector: 'app-setting-modal',
  templateUrl: './setting-modal.component.html',
  styleUrls: ['./setting-modal.component.sass']
})
export class SettingModalComponent implements OnInit {

  public setting: Setting;
  public audioData: Array<SoundSetting> = AudioSettingData.getListSound();
  public audioPlay;
  public loading: boolean;

  constructor(private modalRef: NzModalRef) {
    this.setting = new Setting();
  }

  ngOnInit(): void {
  }

  public async playSound() {
    if (!!this.setting.sound) {
      this.audioPlay = new Audio(`../../assets/audios/${this.setting.sound.fileName}`);
      this.audioPlay.load();
      this.audioPlay.play()
        .then(() => this.audioPlay = null);
    }
  }

  public async onChangeAudio() {
    await this.playSound();
  }

  public onDestroyModal(): void {
    this.modalRef.destroy();
  }
}
