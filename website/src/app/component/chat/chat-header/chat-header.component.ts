import {Component, EventEmitter, Input, OnChanges, OnInit, Output, SimpleChanges} from '@angular/core';
import {Group} from '../../../model/group.model';
import {GroupType} from '../../../const/group-type.const';

@Component({
  selector: 'app-chat-header',
  templateUrl: './chat-header.component.html',
  styleUrls: ['./chat-header.component.sass']
})
export class ChatHeaderComponent implements OnInit, OnChanges {

  @Input() groupSelected: Group;
  @Input() collapseSidebar: boolean;
  @Output() collapseSidebarChange = new EventEmitter<boolean>();

  constructor() { }

  ngOnInit(): void {
  }

  ngOnChanges(changes: SimpleChanges) {
    if (changes.groupSelected && this.groupSelected) {
      this.collapseSidebarChange.emit(true);
    }
  }

  isGroup = (type) => type === GroupType.MANY;

  switchCollapseSidebar() {
    this.collapseSidebar = !this.collapseSidebar;
    this.collapseSidebarChange.emit(this.collapseSidebar);
  }
}
