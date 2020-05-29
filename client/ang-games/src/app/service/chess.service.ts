import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable, observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class ChessService {

  constructor(private http: HttpClient) { }

  public solve(dim: string, pieces: PieceInput[]) {
    const body = { dim, pieces };
    const future = this.http.post('/v1/games/chess', body);
    future.subscribe(param => {
      console.log(param);
    });
  }
}


export class PieceInput {
  label: string;
  letter: string;
  npieces: number;
}
