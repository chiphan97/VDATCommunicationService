import {NgModule} from '@angular/core';
import {Routes, RouterModule} from '@angular/router';
import {MessengerComponent} from './messenger.component';

const routes: Routes = [
  {
    path: '',
    component: MessengerComponent
  },
  {
    path: ':groupId',
    component: MessengerComponent
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class MessengerRouting {
}
