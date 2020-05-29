import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ChessComponent } from '../../components/chess/chess.component';
import { Routes } from '@angular/router';
import { MatSelectModule } from '@angular/material/select';
import { MatCheckboxModule } from '@angular/material/checkbox';
import { MatInputModule } from '@angular/material/input';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatCardModule } from '@angular/material/card';
import { MatButtonModule } from '@angular/material/button';
import { FormsModule} from '@angular/forms';
import { ChessService} from '../../service/chess.service';
import { HttpClientModule } from '@angular/common/http';



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
    MatCheckboxModule,
    MatInputModule,
    MatFormFieldModule,
    MatCardModule,
    MatButtonModule,
    FormsModule,
    HttpClientModule
  ],
  providers: [ChessService]
})
export class ChessModule { }
