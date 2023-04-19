import { Component } from '@angular/core';
import { MatFormField } from '@angular/material/form-field';
import { Router } from '@angular/router';
import { MatButtonModule } from '@angular/material/button';
import { Deck } from '../gatordex';
import { DeckService } from '../deck.service';

@Component({
  selector: 'app-home-page',
  templateUrl: './home-page.component.html',
  styleUrls: ['./home-page.component.css']
})
export class HomePageComponent {
  constructor(private deckService: DeckService, private router: Router){}

  deck: Deck = new Deck();

  save(){
    this.deckService
    .createDeck(this.deck).subscribe(data =>{
      console.log(data)
      this.deck = new Deck

    },
      error => console.log(error));
  }

  createDex() {
    //this.save();
    this.router.navigate(['/make-dex']);
  }

  viewDex() {
    this.router.navigate(['/view-dex']);
  }
}
