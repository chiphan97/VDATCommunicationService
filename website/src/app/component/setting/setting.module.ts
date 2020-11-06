import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { SettingModalComponent } from './setting-modal/setting-modal.component';
import {
  NzButtonModule,
  NzDividerModule,
  NzGridModule,
  NzIconModule, NzModalModule,
  NzSelectModule,
  NzSwitchModule,
  NzTypographyModule
} from 'ng-zorro-antd';
import {FormsModule} from '@angular/forms';

@NgModule({
  declarations: [SettingModalComponent],
  imports: [
    CommonModule,
    NzGridModule,
    NzSelectModule,
    FormsModule,
    NzIconModule,
    NzDividerModule,
    NzSwitchModule,
    NzTypographyModule,
    NzButtonModule,
    NzModalModule
  ]
})
export class SettingModule { }
