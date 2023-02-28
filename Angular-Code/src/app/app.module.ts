import { Component, NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { RouterModule, Routes } from '@angular/router';

//local stuff
import { AppComponent } from './app.component';
import { LoginComponent } from './login/login.component';
import { NewAccountComponent } from './new-account/new-account.component';

import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import{ MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';

const appRoutes: Routes = [
  { path: 'login', component: LoginComponent },
  { path: 'new-account', component: NewAccountComponent },
  {path:'', redirectTo: '/login', pathMatch: 'full'}
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
    NewAccountComponent,
  ],
  bootstrap: [ AppComponent ] //hello
})
export class AppModule { }

export class MaterialModule{ }