import { Component,OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { UserService } from '../user.service';
import { User } from '../user';
import { Router } from '@angular/router';
import { MatFormField } from '@angular/material/form-field';
import { MatButtonModule } from '@angular/material/button';


@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit{
login: FormGroup | undefined;
loginForm: any;

submitted = false;


constructor(private userService: UserService,
  private router: Router, private fb:FormBuilder){ }

ngOnInit(){
  this.loginForm = this.fb.group({
    Email: ['', Validators.required],
    password: ['', Validators.required]
  });
}

onSubmit(){
  this.submitted = true;
  this.save();
}

save(){
  
}
}