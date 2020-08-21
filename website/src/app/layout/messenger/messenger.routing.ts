import {NgModule} from '@angular/core';
import {Routes, RouterModule} from '@angular/router';
import {ChatComponent} from '../../page/chat/chat.component';
import {MessengerComponent} from './messenger.component';


const routes: Routes = [
  {
    path: '',
    component: MessengerComponent,
    children: [
      {
        path: '',
        component: ChatComponent
      }
    ]
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class MessengerRouting {
}
