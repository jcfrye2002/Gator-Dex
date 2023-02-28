// import { BrowserModule } 
//  from '@angular/platform-browser';
// import { NgModule } from '@angular/core';
// import { FormsModule } from '@angular/forms';
// import { AppRoutingModule } 
//   from './app-routing.module';
// import { AppComponent } from './app.component';
// import { CreateUserComponent } 
//   from './create-user/create-user.component';
// import { UserDetailsComponent } 
//   from './user-details/user-details.component';
// import { UserListComponent } 
//   from './user-list/user-list.component';
// import { UpdateUserComponent } 
//   from './update-user/update-user.component';
// import { HttpClientModule } 
//   from '@angular/common/http';
// @NgModule({
//   declarations: [
//     AppComponent,
//     CreateUserComponent,
//     UserDetailsComponent,
//     UserListComponent,
//     UpdateUserComponent
//   ],
//   imports: [
//     BrowserModule,
//     AppRoutingModule,
//     FormsModule,
//     HttpClientModule
//   ],
//   providers: [],
//   bootstrap: [AppComponent]
// })
// export class AppModule { }

import { Component, NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { RouterModule, Routes } from '@angular/router';

//local stuff
import { AppComponent } from './app.component';
import { LoginComponent } from './login/login.component';
import { CreateUserComponent } from './create-user/create-user.component';

import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import{ MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';

const appRoutes: Routes = [
  { path: 'login', component: LoginComponent },
  { path: 'new-account', component: CreateUserComponent },
  { path:'', redirectTo: '/login', pathMatch: 'full'}
];

@NgModule({
  imports: [
    BrowserModule,
    FormsModule,
    RouterModule.forRoot(
      appRoutes,
      { enableTracing: true }
    ),
    BrowserAnimationsModule,
    ReactiveFormsModule,
    MatFormFieldModule,
    MatInputModule
  ],
  exports:[
    FormsModule,
    ReactiveFormsModule,
    MatFormFieldModule,
    MatInputModule
  ],
  declarations: [
    AppComponent,
    LoginComponent,
    CreateUserComponent,
  ],
  bootstrap: [ AppComponent ] //hello
})
export class AppModule { }

export class MaterialModule{ }
