import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';
import {ChatPageComponent} from './chat-page.component';
import {NzGridModule, NzResultModule} from 'ng-zorro-antd';
import {NzResizableModule} from 'ng-zorro-antd/resizable';
import {ChatPageRouting} from './chat-page.routing';
import {ChatModule} from '../../component/chat/chat.module';


@NgModule({
  declarations: [ChatPageComponent],
  imports: [
    CommonModule,
    NzGridModule,
    NzResizableModule,
    NzResultModule,
    ChatModule,
    ChatPageRouting
  ]
})
export class ChatPageModule {
}
