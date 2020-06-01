import { Component, OnInit, Input } from '@angular/core';
import { ChessResult, Combination } from '../../service/chess-result-model'



@Component({
  selector: 'app-chess-paginator',
  templateUrl: './chess-paginator.component.html',
  styleUrls: ['./chess-paginator.component.scss']
})
export class ChessPaginatorComponent implements OnInit {

  @Input()
  chessResult: ChessResult;
  @Input()
  dimension: number;


  dataSource: Combination[];

  constructor() { }

  ngOnInit(): void {
    this.dataSource = this.chessResult.combinations.slice(0, 4);
  }

  fetch(page: number): void {
    this.dataSource = this.chessResult.combinations.slice(page * 4, page * 4 + 4);
  }
}
