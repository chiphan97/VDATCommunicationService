import {Component, OnChanges, OnInit, SimpleChanges} from '@angular/core';
import {Group} from '../../model/group.model';
import {NzResizeEvent} from 'ng-zorro-antd/resizable';
import {ActivatedRoute, Router} from '@angular/router';

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

  constructor(private route: ActivatedRoute,
              private router: Router) {
    this.route.queryParams
      .subscribe(params => {
        console.log(params);
      });
  }

  ngOnInit(): void {
  }

  onEventChange(isChange: boolean) {
    this.changed = isChange;
  }

  onGroupChange(group: Group) {
    this.router.navigate(['messages', group.id]);
  }

  onResize({col}: NzResizeEvent): void {
    cancelAnimationFrame(this.idResize);
    this.idResize = requestAnimationFrame(() => {
      this.colResize = col;
    });
  }
}
