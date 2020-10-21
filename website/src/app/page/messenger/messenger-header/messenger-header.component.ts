import {Component, EventEmitter, Input, OnInit, Output} from '@angular/core';
import {Group} from '../../../model/group.model';

@Component({
  selector: 'app-messenger-header',
  templateUrl: './messenger-header.component.html',
  styleUrls: ['./messenger-header.component.sass']
})
export class MessengerHeaderComponent implements OnInit {

  @Input() groupSelected: Group;

  @Input() collapseSidebarRight: boolean;
  @Output() collapseSidebarRightChange = new EventEmitter<boolean>();

  constructor() { }

  ngOnInit(): void {
  }

  // region Event
  public onSwitchCollapseSidebar() {
    this.collapseSidebarRightChange.emit(!this.collapseSidebarRight);
  }
  // endregion

}
