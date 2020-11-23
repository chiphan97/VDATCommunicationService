import { Component, OnInit, Input } from '@angular/core';

@Component({
  selector: 'app-reply-thread-header',
  templateUrl: './reply-thread-header.component.html',
  styleUrls: ['./reply-thread-header.component.sass']
})
export class ReplyThreadHeaderComponent implements OnInit {
  @Input() closeReplyThread: () => void;
  constructor() { }

  ngOnInit(): void {
  }
}
