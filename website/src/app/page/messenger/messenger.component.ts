import {Component, OnInit} from '@angular/core';
import {NzResizeEvent} from 'ng-zorro-antd/resizable';

@Component({
  selector: 'app-messenger',
  templateUrl: './messenger.component.html',
  styleUrls: ['./messenger.component.sass']
})
export class MessengerComponent implements OnInit {

  public col = 5;
  id = -1;

  constructor() {
  }

  ngOnInit(): void {
  }

  onResize({col}: NzResizeEvent): void {
    cancelAnimationFrame(this.id);
    this.id = requestAnimationFrame(() => {
      this.col = col;
    });
  }
}
