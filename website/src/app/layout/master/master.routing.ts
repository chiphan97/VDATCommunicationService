import {NgModule} from '@angular/core';
import {Routes, RouterModule} from '@angular/router';
import {MasterComponent} from './master.component';

const routes: Routes = [
  {
    path: '',
    component: MasterComponent,
    children: [
      {
        path: '',
        pathMatch: 'full',
        redirectTo: 'messages'
      },
      {
        path: 'messages',
        loadChildren: () => import('./../../page/messenger/messenger.module').then(m => m.MessengerModule)
      },
      {
        path: 'article',
        loadChildren: () => import('./../../page/article/article.module').then(m => m.ArticleModule)
      },
      {
        path: 'article/editor',
        loadChildren: () => import('./../../page/article-editor/article-editor.module').then(m => m.ArticleEditorModule)
      }
    ]
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class MasterRouting {
}
