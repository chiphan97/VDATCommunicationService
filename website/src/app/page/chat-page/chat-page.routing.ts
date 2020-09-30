import {NgModule} from '@angular/core';
import {Routes, RouterModule} from '@angular/router';
import {ChatPageComponent} from './chat-page.component';

const routes: Routes = [
  {
    path: '',
    component: ChatPageComponent
  },
  {
    path: ':groupId',
    component: ChatPageComponent
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class ChatPageRouting {
}
