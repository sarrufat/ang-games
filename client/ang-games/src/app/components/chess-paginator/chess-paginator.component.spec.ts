import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ChessPaginatorComponent } from './chess-paginator.component';

describe('ChessPaginatorComponent', () => {
  let component: ChessPaginatorComponent;
  let fixture: ComponentFixture<ChessPaginatorComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ChessPaginatorComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ChessPaginatorComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
