// import { Component,OnInit } from '@angular/core';
// import { FormBuilder, FormGroup, Validators } from '@angular/forms';
// import { Observable } from "rxjs";
// import { UserService } from "../user.service";
// import { User } from "../user";
// import { Router, ActivatedRoute } 
//   from '@angular/router';


// @Component({
//   selector: 'app-login',
//   templateUrl: './login.component.html',
//   styleUrls: ['./login.component.css']
// })
// export class LoginComponent implements OnInit{
// login: FormGroup | undefined;
// loginForm: any;
// id: string;
//   user: User;


// //constructor(private fb: FormBuilder){ }

// constructor(private route: ActivatedRoute,
//   private router: Router,
//     private userService: UserService,private fb: FormBuilder) { }

// ngOnInit(){
//   this.id = this.route.snapshot.params['id'];
//   this.login = this.fb.group({
//     email: ['', Validators.required],
//     password: ['', Validators.required]
//   });
// }

// onSubmit(){

// }

// list(){
//   this.router.navigate(['users']);
// }
// }

import { Component,OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { UserService } from '../user.service';
import { User } from '../user';
import { Router } from '@angular/router';
import { MatFormField, MatFormFieldControl } from '@angular/material/form-field';
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
email: string;
password : string;


constructor(private userService: UserService,
  private router: Router, private fb:FormBuilder){ }

ngOnInit(){
  this.loginForm = this.fb.group({
    Email: ['', Validators.required],
    password: ['', Validators.required]
  });
}

onSubmit(){
  this.loginUser();
  this.submitted = true;
  this.save(); // logic here to sign in 
  this.router.navigate(['/home']);
}

save(){
  
}

clickNewUser(){
  this.router.navigate(['/new-user']);
}
gotoList() {
    this.router.navigate(['/users']);
  }
  loginUser() {
    this.userService.updateUser(this.email, this.password)
      .subscribe(data => {
        console.log(data);
        this.gotoList();
      }, error => console.log(error));
  }
}