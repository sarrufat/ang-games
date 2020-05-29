import { Component, OnInit, Input } from '@angular/core';
import { FormControl } from '@angular/forms';
import { PieceInput, ChessService } from '../../service/chess.service';
import { TaskId } from '../../service/chess-result-model';
import { timer, Observable, observable, Subscription} from 'rxjs';



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
    this.service.solve(this.selected, this.pieces).subscribe(result => {
      this.taskId = result;
      this.formDisabled = true;
      this.tmout = timer(10000).subscribe(observer => {
        this.formDisabled = false;
        this.tmout.unsubscribe();
      });
    });
  }

}
