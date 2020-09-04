import {Component, OnChanges, OnInit, SimpleChanges} from '@angular/core';
import {UserOnlineService} from '../../service/user-online.service';
import { ActivatedRoute } from '@angular/router';
import {Group} from '../../model/group.model';

@Component({
  selector: 'app-messenger',
  templateUrl: './messenger.component.html',
  styleUrls: ['./messenger.component.sass']
})
export class MessengerComponent implements OnInit {

  public collapseSidebar = false;
  public groupSelected: Group;
  public changed: boolean;

  constructor() {
  }

  ngOnInit(): void {
  }

  onEventChange(isChange: boolean) {
    this.changed = isChange;
  }
}
