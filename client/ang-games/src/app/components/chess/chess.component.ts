import { Component, OnInit, Input } from '@angular/core';
import { FormControl } from '@angular/forms';
import { PieceInput, ChessService } from '../../service/chess.service';


@Component({
  selector: 'app-chess',
  templateUrl: './chess.component.html',
  styleUrls: ['./chess.component.scss'],
})

export class ChessComponent implements OnInit {

  booards = new FormControl();
  boardList: string[] = ['4x4', '5x5', '6x6', '7x7', '8x8'];
  selected = '4x4';

  pieces: PieceInput[] = [
    { label: '\u2654', letter: 'K', npieces: 1 },
    { label: '\u2655', letter: 'Q', npieces: 2 },
    { label: '\u2656', letter: 'R', npieces: 3 },
    { label: '\u2657', letter: 'B', npieces: 4 },
    { label: '\u2658', letter: 'N', npieces: 5 }
  ];
  input = '1';

  constructor(private service: ChessService) { }

  ngOnInit(): void {
  }

  onSolve(): void {
    console.log('onSolve:' + this.selected);
    this.service.solve(this.selected, this.pieces);
  }

}
