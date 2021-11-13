import { Component, OnInit } from '@angular/core';
import { AlertService } from '@app/_alert';
import { GameInfo, User } from '@app/_models';
import { GamesService } from '@app/_services';
import { first } from 'rxjs/operators';

type GameInfoView = Partial<GameInfo> & { full: boolean }

@Component({
  selector: 'app-games',
  templateUrl: './games.component.html',
  styleUrls: ['./games.component.less']
})
export class GamesComponent implements OnInit {

  loading = false;
  games: GameInfoView[];

  constructor(private gamesService: GamesService, private alertService: AlertService) { }

  ngOnInit() {
    this.updateGames()
  }

  join(game: GameInfo) {
    this.gamesService.join(game).pipe(first())
      .subscribe({
        next: () => {
          this.alertService.info('Game joined.', {autoClose: true})
          this.updateGames()
        },
        error: error => {
          this.alertService.error('Unable to join game... Try again', {autoClose: true})
          this.loading = false;
        }
      });
  }

  updateGames() {
    this.loading = true;
    this.gamesService.getAll().pipe(first()).subscribe(g => {
      this.loading = false;
      this.games = g.map(info => { let gi = info as GameInfoView; gi.full = info.players.joined === info.players.required; return gi });
    });
  }
}
