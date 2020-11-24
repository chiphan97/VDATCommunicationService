import { Component, OnInit } from '@angular/core';
import {FormControl, FormGroup, Validators} from '@angular/forms';

@Component({
  selector: 'app-article-editor',
  templateUrl: './article-editor.component.html',
  styleUrls: ['./article-editor.component.sass']
})
export class ArticleEditorComponent implements OnInit {

  public formArticle: FormGroup;

  constructor() {
    this.formArticle = this.createFormGroup();
  }

  ngOnInit(): void {
  }

  private createFormGroup(): FormGroup {
    return new FormGroup({
      title: new FormControl('', [Validators.required])
    });
  }
}
