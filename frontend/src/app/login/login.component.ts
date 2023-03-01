import { Component,OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit{
login: FormGroup | undefined;
loginForm: any;
clicked:boolean;


constructor(private fb: FormBuilder){ }

ngOnInit(){
  this.loginForm = this.fb.group({
    email: ['', Validators.required],
    password: ['', Validators.required]
  });
}

onSubmit(){
    this.clicked = true;
}
}
