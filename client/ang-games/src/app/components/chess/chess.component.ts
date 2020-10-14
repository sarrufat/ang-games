import { Component, OnInit, Input } from '@angular/core';
import { FormControl } from '@angular/forms';
import { PieceInput, ChessService } from '../../service/chess.service';
import { TaskId, ChessResult } from '../../service/chess-result-model';
import { timer, Subscription } from 'rxjs';



@Component({
  selector: 'app-chess',
  templateUrl: './chess.component.html',
  styleUrls: ['./chess.component.scss'],
})

export class ChessComponent implements OnInit {

  booards = new FormControl();
  boardList: string[] = ['4x4', '5x5', '6x6', '7x7', '8x8'];
  selected = '4x4';
  taskId: TaskId;
  formDisabled = false;
  chessResult: ChessResult;
  dimension = 4;

  pieces: PieceInput[] = [
    { label: '\u2654', letter: 'K', npieces: 0 },
    { label: '\u2655', letter: 'Q', npieces: 0 },
    { label: '\u2656', letter: 'R', npieces: 0 },
    { label: '\u2657', letter: 'B', npieces: 0 },
    { label: '\u2658', letter: 'N', npieces: 0 }
  ];
  input = '1';
  tmout: Subscription;

  constructor(private service: ChessService) { }

  ngOnInit(): void {
  }

  onSolve(): void {
    console.log('onSolve:' + this.selected);
    this.dimension = +this.selected.substring(0, 1);
    let sum = 0;
    for (const p of this.pieces) {
      sum += p.npieces;
    }
    this.chessResult = undefined;
    if (sum >= this.dimension / 2) {
      this.service.solve(this.selected, this.pieces).subscribe(result => {
        this.taskId = result;
        this.formDisabled = true;
        this.tmout = timer(1000, 1000).subscribe(observer => {
          this.service.checkCompletion(this.taskId).subscribe(chessResult => {
            console.log('chessResult.done ' + chessResult.done);
            if (chessResult.done === true) {
              console.log(chessResult.combinations)
              this.tmout.unsubscribe();
              this.formDisabled = false;
              this.chessResult = chessResult;
              this.chessResult.combLenght = this.chessResult.combinations.length;
            //  this.chessResult.numCombinations = this.chessResult.combinations.length;
              // tslint:disable-next-line:max-line-length
              this.chessResult.msg = 'Found ' + this.chessResult.numCombinations + ' in ' + this.chessResult.ms + 'ms. with ' + this.chessResult.iterations + ' path iterations';
            }
          });
        });
      });
    }
  }

}
