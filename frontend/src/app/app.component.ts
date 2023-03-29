import { Component } from '@angular/core';
import { MatToolbarModule } from '@angular/material/toolbar';
import { Router } from '@angular/router';
import { MatButtonModule } from '@angular/material/button';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'Gator-Dex';

  constructor(private router: Router,private toolbar: MatToolbarModule) { }

  clickHome() {
    this.router.navigate(['/home']);
  }

  clickLogin(){
    this.router.navigate(['/login']);
  }

  clickNewUser(){
    this.router.navigate(['/new-user'])
  }

  clickUserList(){
    this.router.navigate(['/user-list']);
  }
}