import { Component,OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { UserService } from '../user.service';
import { User } from '../user';
import { Router } from '@angular/router';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatButtonModule } from '@angular/material/button';
import { Deck } from '../gatordex';

@Component({
  selector: 'app-create-dex',
  templateUrl: './create-dex.component.html',
  styleUrls: ['./create-dex.component.css']
})
export class CreateDexComponent {
  constructor(private userService: UserService,
    private router: Router, private fb:FormBuilder){}

  newCard: FormGroup | undefined;
  cardForm: FormGroup<any>;

  submitted = false;
  deck: Deck = new Deck();
  
    ngOnInit(){
      this.cardForm = this.fb.group({
        front: ['', Validators.required],
        back: ['', Validators.required]
      });
    }
    newDeck (): void{
      this.submitted = false;
      this.deck = new Deck();
    }
    save() {
      this.userService
      .createUser(this.deck).subscribe(data => {
        console.log(data)
        this.deck = new Deck();
        this.gotoList();
      }, 
      error => console.log(error));
    }
    onSubmit(){
      this.submitted = true;
      this.save();
      this.router.navigate(['/home']);
    }

    clickAddCard(){
      location.reload();
    }

    gotoList() {
      this.router.navigate(['/decks']);
    }

    clickHome(){
      this.router.navigate(['/home']);
    }
}
