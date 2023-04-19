import { Component,OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { UserService } from '../user.service';
import { User } from '../user';
import { Router } from '@angular/router';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatButtonModule } from '@angular/material/button';
import { Deck } from '../gatordex';

@Component({
  selector: 'app-view-dex',
  templateUrl: './view-dex.component.html',
  styleUrls: ['./view-dex.component.css']
})
export class ViewDexComponent {
    
}
