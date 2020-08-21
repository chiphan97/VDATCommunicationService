import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-messenger',
  templateUrl: './messenger.component.html',
  styleUrls: ['./messenger.component.sass']
})
export class MessengerComponent implements OnInit {

  public isCollapsed = false;

  constructor() { }

  ngOnInit(): void {
  }

}
