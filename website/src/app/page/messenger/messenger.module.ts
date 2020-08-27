import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';
import {MessengerComponent} from './messenger.component';
import {NzGridModule} from 'ng-zorro-antd';
import {NzResizableModule} from 'ng-zorro-antd/resizable';
import {ComponentModule} from '../../component/component.module';


@NgModule({
  declarations: [MessengerComponent],
  imports: [
    CommonModule,
    NzGridModule,
    NzResizableModule,
    ComponentModule
  ]
})
export class MessengerModule {
}
