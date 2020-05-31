import { Component, OnInit, ViewChild, ElementRef, AfterViewInit, Input } from '@angular/core';

const BLOCK_SIZE = 50;

@Component({
  selector: 'app-chess-canvas',
  templateUrl: './chess.canvas.component.html',
  styleUrls: ['./chess.canvas.component.scss']
})
export class ChessCanvasComponent implements OnInit, AfterViewInit {

@Input('dimension')cdim = 4;
@ViewChild('chessboard') public canvas: ElementRef;
  private cx: CanvasRenderingContext2D;

  constructor() { }

  ngOnInit(): void {
  }

  public ngAfterViewInit() {
    const canvasEl: HTMLCanvasElement = this.canvas.nativeElement;
    this.cx = canvasEl.getContext('2d');
    canvasEl.width = this.cdim * BLOCK_SIZE;
    canvasEl.height = this.cdim * BLOCK_SIZE;
    this.drawBoard(this.cx);
  }

  private drawBoard(cx: CanvasRenderingContext2D ) {
    for (let x = 0; x < this.cdim; x++) {
      for (let y = 0; y < this.cdim; y++) {
        const bcolor = (x + y) % 2 ? 'rgb(230,200,50)' : 'rgb(90,90,50)';
        cx.fillStyle = bcolor;
        cx.fillRect(x * BLOCK_SIZE, y * BLOCK_SIZE, BLOCK_SIZE, BLOCK_SIZE);
        cx.stroke();
      }
    }
    cx.lineWidth = 2;
    cx.strokeRect(0, 0, this.cdim * BLOCK_SIZE, this.cdim * BLOCK_SIZE);
  }
}
