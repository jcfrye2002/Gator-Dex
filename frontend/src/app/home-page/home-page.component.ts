import { Component } from '@angular/core';
import { MatFormField } from '@angular/material/form-field';
import { Router } from '@angular/router';

@Component({
  selector: 'app-home-page',
  templateUrl: './home-page.component.html',
  styleUrls: ['./home-page.component.css']
})
export class HomePageComponent {
  constructor(private router: Router){}

  createDex() {
    this.router.navigate(['/create']);
  }

  viewDex() {
    this.router.navigate(['/view-dex']);
  }
}
