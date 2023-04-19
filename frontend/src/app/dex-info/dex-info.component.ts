import { Component,OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { MatFormField, MatFormFieldControl } from '@angular/material/form-field';
import { MatButtonModule } from '@angular/material/button';
import { Deck } from '../gatordex';
import { DeckService } from '../deck.service';
import { HttpClientModule } from '@angular/common/http';
import { HttpClientTestingModule, HttpTestingController } from '@angular/common/http/testing';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';


@Component({
  selector: 'app-dex-info',
  templateUrl: './dex-info.component.html',
  styleUrls: ['./dex-info.component.css']
})
export class DexInfoComponent implements OnInit{
  //need name and class code
  createDex: FormGroup | undefined
  createDexForm: any;

  deck: Deck = new Deck();
  submitted = false;
  deckForm : FormGroup<any>;

  

  constructor(private deckService: DeckService, private router: Router, 
    private fb:FormBuilder){}

    

  ngOnInit() {
    this.createDexForm = this.fb.group({
      gatordeckName:['',Validators.required],
      classCode:['',Validators.required],
    })
  }

  newDeck(){
    this.submitted = false;
    this.deck = new Deck();
  }

  save(){
    this.deckService
    .createDeck(this.deck).subscribe(data=>{
      console.log(data)
      this.deck = new Deck();
      this.toCreate();
    },
    error => console.log(error));
  }

  onSubmit() {
    this.submitted = true;
    this.save();
    this.router.navigate(['/create'])
  }

  toCreate(){
    this.router.navigate(['/create']);
  }
}
