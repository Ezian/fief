import { Component, OnInit } from '@angular/core';
import { BOARD } from '../mock-board';
import { Board, City } from '../model';

@Component({
  selector: 'app-board',
  templateUrl: './board.component.html',
  styleUrls: ['./board.component.css'],
})
export class BoardComponent implements OnInit {
  board: Board = BOARD;

  constructor() {}

  ngOnInit(): void {}

  x(city: City): number {
    return 50 + 200 * (city.id % 10);
  }

  y(city: City): number {
    return 115 * Math.round(city.id / 10);
  }
}
