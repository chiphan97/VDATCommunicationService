import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';
import {CommunityComponent} from './community.component';
import {CommunityRouting} from './community.routing';


@NgModule({
  declarations: [
    CommunityComponent
  ],
  imports: [
    CommonModule,
    CommunityRouting
  ]
})
export class CommunityModule {
}
