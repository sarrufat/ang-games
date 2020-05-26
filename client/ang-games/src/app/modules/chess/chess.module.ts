import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ChessComponent } from '../../components/chess/chess.component';
import { Routes, RouterModule } from '@angular/router';


const routes: Routes = [
  {
    path: 'chess',
    component: ChessComponent
  }
];

@NgModule({
  declarations: [ChessComponent],
  imports: [
    CommonModule
  ]
})
export class ChessModule { }
