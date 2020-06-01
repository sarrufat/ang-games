import { Component, OnInit, ViewChild, ElementRef, Input, TemplateRef, AfterViewChecked } from '@angular/core';
import { Combination } from '../../service/chess-result-model';

const BLOCK_SIZE = 50;

@Component({
  selector: 'app-chess-canvas',
  templateUrl: './chess.canvas.component.html',
  styleUrls: ['./chess.canvas.component.scss']
})
export class ChessCanvasComponent implements OnInit, AfterViewChecked {

  @ViewChild('piecesImg') piecesImage: ElementRef;

  // tslint:disable-next-line:no-input-rename
  @Input('dimension') cdim = 4;
  @Input() combination: Combination;
  @ViewChild('chessboard') public canvas: ElementRef;
  private cx: CanvasRenderingContext2D;

  constructor() { }

  ngOnInit(): void {
  }

  public ngAfterViewChecked() {
    const canvasEl: HTMLCanvasElement = this.canvas.nativeElement;
    const piecesImages: HTMLImageElement = this.piecesImage.nativeElement;
    this.cx = canvasEl.getContext('2d');
    canvasEl.width = this.cdim * BLOCK_SIZE;
    canvasEl.height = this.cdim * BLOCK_SIZE;
    this.drawBoard(this.cx);
    this.drawPieces(piecesImages);
  }

  private drawPieces(piecesImages: HTMLImageElement) {
    const pOrder = ['P', 'R', 'N', 'B', 'Q', 'K'];
    for (const pos of this.combination.positions) {
      const delta = pOrder.indexOf(pos.piece) * 50;
      const canvasOffsetX = pos.x * BLOCK_SIZE;
      const canvasOffsetY = pos.y * BLOCK_SIZE;
      this.cx.drawImage(piecesImages, delta, 0, BLOCK_SIZE, BLOCK_SIZE, canvasOffsetX, canvasOffsetY, BLOCK_SIZE, BLOCK_SIZE);

    }
  }

  private drawBoard(cx: CanvasRenderingContext2D) {
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
