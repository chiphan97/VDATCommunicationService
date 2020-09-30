import {NgModule} from '@angular/core';
import {Routes, RouterModule} from '@angular/router';
import {AuthGuard} from './guard/auth.guard';


const routes: Routes = [
  {
    path: '',
    loadChildren: () => import('./layout/master/master.module').then(m => m.MasterModule),
    canActivateChild: [AuthGuard]
  },
  {
    path: 'auth',
    loadChildren: () => import('./layout/auth/auth.module').then(m => m.AuthModule),
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRouting {
}