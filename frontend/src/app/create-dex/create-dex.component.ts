import { Component,OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { UserService } from '../user.service';
import { User } from '../user';
import { Router } from '@angular/router';
import { MatFormField } from '@angular/material/form-field';
import { MatButtonModule } from '@angular/material/button';

@Component({
  selector: 'app-create-dex',
  templateUrl: './create-dex.component.html',
  styleUrls: ['./create-dex.component.css']
})
export class CreateDexComponent {
  constructor(private userService: UserService,
    private router: Router, private fb:FormBuilder){}

  newCard: FormGroup | undefined;
  cardForm: any;

  submitted = false;

    ngOnInit(){
      this.cardForm = this.fb.group({
        front: ['', Validators.required],
        back: ['', Validators.required]
      });
    }

    onSubmit(){
      this.submitted = true;
      this.router.navigate(['/home']);
    }

    clickAddCard(){
      location.reload();
    }

    clickHome(){
      this.router.navigate(['/home']);
    }
}
