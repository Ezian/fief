import { Component, OnInit } from '@angular/core';
import { GameInfo, User } from '@app/_models';
import { GamesService } from '@app/_services';
import { first } from 'rxjs/operators';

type GameInfoView = Partial<GameInfo> & {full: boolean}

@Component({
  selector: 'app-games',
  templateUrl: './games.component.html',
  styleUrls: ['./games.component.less']
})
export class GamesComponent implements OnInit {

  loading = false;
  games: GameInfoView[];

  constructor(private gamesService: GamesService) { }

  ngOnInit() {
    this.updateGames()
  }

  join(game: GameInfo){
    console.log("test")
    this.gamesService.join(game).pipe(first())
    .subscribe({
        next: () => {
          this.updateGames()
        },
        error: error => {
          // TODO handle error through a alert : Check maybe https://jasonwatmore.com/post/2019/07/05/angular-8-alert-toaster-notifications
            this.loading = false;
        }
    });
  }

  updateGames(){
    this.loading = true;
    this.gamesService.getAll().pipe(first()).subscribe(g => {
        this.loading = false;
        this.games = g.map(info => {let gi = info as GameInfoView; gi.full = info.players.joined === info.players.required; return gi});
    });
  }
}
