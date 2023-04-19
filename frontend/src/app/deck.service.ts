import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class DeckService {

  private baseUrl = 'http://localhost:9080/users';

  constructor(private http: HttpClient) { }

  createDeck(deck: Object): Observable<Object> {
    return this.http.post(`${this.baseUrl}`, deck);
  }
}
