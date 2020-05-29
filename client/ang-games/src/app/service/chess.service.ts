import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable, observable } from 'rxjs';
import { ChessResult, Combination, Position, TaskId } from './chess-result-model';

@Injectable({
  providedIn: 'root'
})
export class ChessService {

  constructor(private http: HttpClient) { }

  public solve(dim: string, pieces: PieceInput[]): Observable<TaskId> {
    const body = { dim, pieces };
    console.log(JSON.stringify(body));
    return this.http.post('/v1/games/chess', body) as Observable<TaskId>;
  }

  public checkCompletion(id: TaskId): Observable<ChessResult> {
    const future = this.http.get('/v1/games/chess/' + id.taskId);
    return future as Observable<ChessResult>;
  }
}


export class PieceInput {
  label: string;
  letter: string;
  npieces: number;
}
