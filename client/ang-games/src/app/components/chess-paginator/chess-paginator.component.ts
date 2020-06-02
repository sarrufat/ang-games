import { Component, OnInit, Input, OnChanges, SimpleChanges } from '@angular/core';
import { ChessResult, Combination } from '../../service/chess-result-model'



@Component({
  selector: 'app-chess-paginator',
  templateUrl: './chess-paginator.component.html',
  styleUrls: ['./chess-paginator.component.scss']
})
export class ChessPaginatorComponent implements OnInit, OnChanges {

  @Input()
  chessResult: ChessResult;
  @Input()
  dimension: number;
  pageSize = 4;

  dataSource: Combination[];

  constructor() { }
  ngOnChanges(changes: SimpleChanges): void {
    if (this.chessResult !== undefined) {
      this.pageSize = this.dimension <= 5 ? 4 : 2;
    }
  }


  ngOnInit(): void {
    this.dataSource = this.chessResult.combinations.slice(0, this.pageSize);
  }

  fetch(page: number): void {
    const pg = this.pageSize;
    this.dataSource = this.chessResult.combinations.slice(page * pg, page * pg + pg);
  }
}
