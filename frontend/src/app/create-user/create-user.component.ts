import { UserService } from '../user.service';
import { User } from '../user';
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { MatFormFieldModule } from '@angular/material/form-field';


@Component({
  selector: 'app-create-user',
  templateUrl: './create-user.component.html',
  styleUrls: ['./create-user.component.css']
})
export class CreateUserComponent implements OnInit {
  createUser: FormGroup | undefined
  createUserForm: any;

  user: User = new User();
  submitted = false;
accountForm: FormGroup<any>;

  constructor(private userService: UserService,
    private router: Router, private fb:FormBuilder) { }

    
  ngOnInit() {
    this.accountForm = this.fb.group({
      firstName:['',Validators.required],
      lastName:['',Validators.required],
      email:['',Validators.required],
      password:['',Validators.required],
  })
}

  newUser(): void {
    this.submitted = false;
    this.user = new User();
  }

  save() {
    this.userService
    .createUser(this.user).subscribe(data => {
      console.log(data)
      this.user = new User();
      this.gotoList();
    }, 
    error => console.log(error));
  }

  onSubmit() {
    this.submitted = true;
    this.save();    
  }

  gotoList() {
    this.router.navigate(['/users']);
  }
}