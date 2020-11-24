import {AfterViewInit, Component, OnInit} from '@angular/core';
import EditorJS from '@editorjs/editorjs';
import * as Header from '@editorjs/header';
import * as SimpleImage from '@editorjs/simple-image';
import * as Checklist from '@editorjs/checklist';
import List from '@editorjs/list';
import Underline from '@editorjs/underline';
import * as Embed from '@editorjs/embed';
import * as Quote from '@editorjs/quote';
import * as Delimiter from '@editorjs/delimiter';
import * as Table from '@editorjs/table';
import * as Marker from '@editorjs/marker';
import * as Warning from '@editorjs/warning';
import * as Paragraph from '@editorjs/paragraph';

@Component({
  selector: 'app-editor',
  templateUrl: './editor.component.html',
  styleUrls: ['./editor.component.sass']
})
export class EditorComponent implements OnInit, AfterViewInit {

  private editorInstance: EditorJS;

  constructor() {
  }

  ngOnInit(): void {
  }

  ngAfterViewInit() {
    this.editorInstance = new EditorJS({
      holder: 'editor',
      autofocus: true,
      minHeight: 150,
      tools: {
        header: {
          class: Header,
          shortcut: 'CTRL+SHIFT+H',
          config: {
            levels: [1, 2, 3, 4, 5, 6],
            defaultLevel: 3
          }
        },
        image: SimpleImage,
        checklist: {
          class: Checklist,
          inlineToolbar: true,
        },
        list: {
          class: List,
          inlineToolbar: true,
        },
        embed: {
          class: Embed,
          config: {
            services: {
              youtube: true,
              coub: true
            }
          }
        },
        quote: {
          class: Quote,
          inlineToolbar: true,
          shortcut: 'CTRL+SHIFT+O',
          config: {
            quotePlaceholder: 'Enter a quote',
            captionPlaceholder: 'Quote\'s author',
          },
        },
        delimiter: Delimiter,
        underline: Underline,
        table: {
          class: Table,
          inlineToolbar: true,
          config: {
            rows: 2,
            cols: 3,
          },
        },
        Marker: {
          class: Marker,
          shortcut: 'CTRL+SHIFT+M',
        },
        warning: {
          class: Warning,
          inlineToolbar: true,
          shortcut: 'CMD+SHIFT+W',
          config: {
            titlePlaceholder: 'Title',
            messagePlaceholder: 'Message',
          },
        },
        paragraph: {
          class: Paragraph,
          inlineToolbar: true,
        },
      }
    });
  }

}
