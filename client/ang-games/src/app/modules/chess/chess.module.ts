import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ChessComponent } from '../../components/chess/chess.component';
import { Routes  } from '@angular/router';
import {MatSelectModule} from '@angular/material/select';
import {MatCheckboxModule} from '@angular/material/checkbox';


const routes: Routes = [
  {
    path: 'chess',
    component: ChessComponent
  }
];

@NgModule({
  declarations: [ChessComponent],
  imports: [
    CommonModule,
    MatSelectModule,
    MatCheckboxModule
  ]
})
export class ChessModule { }
