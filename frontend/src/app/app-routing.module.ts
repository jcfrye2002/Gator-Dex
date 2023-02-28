import { UserDetailsComponent } 
   from './user-details/user-details.component';
import { CreateUserComponent } 
   from './create-user/create-user.component';
import { NgModule } from '@angular/core';
import { Routes, RouterModule } 
   from '@angular/router';
import {UserListComponent } 
  from './user-list/user-list.component';
import { UpdateUserComponent } 
   from './update-user/update-user.component';
import { LoginComponent } from './login/login.component';

const routes: Routes = [
{ path: 'login',component: LoginComponent},
  { path: 'users', component: UserListComponent },
  { path: 'add', component: CreateUserComponent },
  { path: 'update/:id', component: UpdateUserComponent },
  { path: 'details/:id', component: UserDetailsComponent }, 
  { path: '', redirectTo: '/login', pathMatch: 'full' },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule],
  declarations: [LoginComponent]
})
export class AppRoutingModule { }