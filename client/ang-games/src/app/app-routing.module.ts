import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { AppComponent } from './app.component';
import {ChessComponent} from './components/chess/chess.component';


const routes: Routes = [
  {
    path: 'chess',
    component: ChessComponent,
    children: [
      {path: '', component: ChessComponent}
    ],
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
