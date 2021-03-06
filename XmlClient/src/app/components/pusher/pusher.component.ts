import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import Pusher from 'pusher-js'
import { INotification } from 'src/app/interfaces/notification';

@Component({
  selector: 'app-pusher',
  templateUrl: './pusher.component.html',
  styleUrls: ['./pusher.component.css']
})
export class PusherComponent  implements OnInit {
  username = 'username';
  message : INotification = {} as INotification
  notifications : INotification[] = [];

  constructor(private http: HttpClient) {
    this.notifications.forEach(n => {
      n = {} as INotification
    })
  }

  ngOnInit(): void {
    Pusher.logToConsole = true;

    const pusher = new Pusher('dd3ce2a9c4a58e3577a4', {
      cluster: 'eu'
    });

    const channel = pusher.subscribe('notifications');
    channel.bind('notification', (data: never) => {
      //this.notifications.push(data);
      console.log(data)
    });
  }

  JEDIGOVNA(event : any){
    this.message = event.target.value
  }

  // submit(): void {
  //   this.http.post('http://localhost:9090/notifications/create', {
  //     "content" : this.message.content,
  //     "from" : this.username,
  //     "to" : localStorage.getItem('username'),
  //     "redirectPath" : "/notifications",
  //     "notificationType" : "POST"
  //   }).subscribe(() => console.log(this.notifications));
  // }
}