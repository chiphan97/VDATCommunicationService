import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-list-user-chat',
  templateUrl: './list-user-chat.component.html',
  styleUrls: ['./list-user-chat.component.sass']
})
export class ListUserChatComponent implements OnInit {

  loading = false;
  data = [
    {
      title: 'User 1'
    },
    {
      title: 'User 2'
    },
    {
      title: 'User 3'
    },
    {
      title: 'User 4'
    }
  ];

  constructor() { }

  ngOnInit(): void {
  }

}
