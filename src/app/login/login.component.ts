import { Component, OnInit } from '@angular/core';
import { FormGroup, FormControl,Validators } from '@angular/forms';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})
export class LoginComponent implements OnInit{
  loginForm!: FormGroup;
hide: any;
  constructor(){ }

  ngOnInit(): void {
    this.loginForm = new FormGroup(
     {
      email : new FormControl('',[Validators.required,Validators.email]), //could easily solve problems with text but idk
      password : new FormControl('',[Validators.required,Validators.minLength(6)])
     }
    );
  }

  onLogin(){

  }
}
