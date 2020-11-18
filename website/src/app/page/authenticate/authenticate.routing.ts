import {NgModule} from '@angular/core';
import {Routes, RouterModule} from '@angular/router';
import {AuthenticateComponent} from './authenticate/authenticate.component';
import {IntegratedComponent} from './integrated/integrated.component';

const routes: Routes = [
  {
    path: '',
    component: AuthenticateComponent
  },
  {
    path: 'integrated',
    component: IntegratedComponent
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class AuthenticateRouting {
}
