import { Component, OnInit } from '@angular/core';
import {Form, FormBuilder, FormGroup, Validators} from '@angular/forms'

@Component({
  selector: 'app-new-account',
  templateUrl: './new-account.component.html',
  styleUrls: ['./new-account.component.css']
})

export class NewAccountComponent implements OnInit{
  account: FormGroup | undefined;
  accountForm: any;
  constructor(private fb: FormBuilder){

  }

  ngOnInit(): void {
    this.accountForm = this.fb.group({
      firstName:['',Validators.required],
      lastName:['',Validators.required],
      email:['',Validators.required],
      password:['',Validators.required],
      reEnterPassword:['',Validators.required]
    })
  }

  onSubmit(){
    //new account logic here
  }
}
