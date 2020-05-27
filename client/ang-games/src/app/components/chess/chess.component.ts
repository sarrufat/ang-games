import { Component, OnInit } from '@angular/core';
import {FormControl} from '@angular/forms';

class PieceInput {
  label: string;
  checked: boolean;
  npieces: number;
}
@Component({
  selector: 'app-chess',
  templateUrl: './chess.component.html',
  styleUrls: ['./chess.component.scss']
})

export class ChessComponent implements OnInit {

  booards = new FormControl();
  boardList: string[] = ['4x4', '5x5', '6x6', '7x7', '8x8'];
  pieces: PieceInput[] = [
    {label: 'Kings', checked: false, npieces: 0},
    {label: 'Queens', checked: false, npieces: 0},
    {label: 'Bishops', checked: false, npieces: 0},
    {label: 'Rooks', checked: false, npieces: 0},
    {label: 'Knights', checked: false, npieces: 0}

  ];
  constructor() { }

  ngOnInit(): void {
  }

}
