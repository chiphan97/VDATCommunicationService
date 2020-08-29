import {Component, EventEmitter, Input, OnInit, Output} from '@angular/core';
import {MessengerDrawerComponent} from '../messenger-drawer/messenger-drawer.component';
import {NzDrawerService} from 'ng-zorro-antd';

@Component({
  selector: 'app-messenger-header',
  templateUrl: './messenger-header.component.html',
  styleUrls: ['./messenger-header.component.sass']
})
export class MessengerHeaderComponent implements OnInit {

  @Input() collapseSidebar: boolean;
  @Output() collapseSidebarChange = new EventEmitter<boolean>();

  constructor() { }

  ngOnInit(): void {
  }

  switchCollapseSidebar() {
    this.collapseSidebar = !this.collapseSidebar;
    this.collapseSidebarChange.emit(this.collapseSidebar);
  }
}
