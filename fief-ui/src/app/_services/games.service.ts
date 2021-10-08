import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { environment } from '@environments/environment';

import { GameInfo } from '@app/_models'
@Injectable({
  providedIn: 'root'
})
export class GamesService {

  constructor(private http: HttpClient) { }

  getAll(){
    return this.http.get<GameInfo[]>(`${environment.apiUrl}/games`)
  }
}
