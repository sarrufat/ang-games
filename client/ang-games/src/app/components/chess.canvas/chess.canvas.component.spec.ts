import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ChessCanvasComponent } from './chess.canvas.component';

describe('Chess.CanvasComponent', () => {
  let component: ChessCanvasComponent;
  let fixture: ComponentFixture<ChessCanvasComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ChessCanvasComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ChessCanvasComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
