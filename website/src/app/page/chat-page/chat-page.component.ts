import { Component, OnInit } from '@angular/core';
import {Group} from '../../model/group.model';
import {NzResizeEvent} from 'ng-zorro-antd/resizable';

@Component({
  selector: 'app-chat-page',
  templateUrl: './chat-page.component.html',
  styleUrls: ['./chat-page.component.sass']
})
export class ChatPageComponent implements OnInit {

  public collapseSidebar = false;
  public groupSelected: Group;
  public changed: boolean;
  public colResize = 5;
  private idResize = -1;

  constructor() {
  }

  ngOnInit(): void {
  }

  onEventChange(isChange: boolean) {
    this.changed = isChange;
  }

  onResize({ col }: NzResizeEvent): void {
    cancelAnimationFrame(this.idResize);
    this.idResize = requestAnimationFrame(() => {
      this.colResize = col;
    });
  }
}
