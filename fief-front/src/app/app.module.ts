import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppRoutingModule } from './app-routing.module';
import { NgxGraphModule } from '@swimlane/ngx-graph';
import { AppComponent } from './app.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { BoardComponent } from './board/board.component';

@NgModule({
  declarations: [AppComponent, BoardComponent],
  imports: [
    BrowserModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    NgxGraphModule,
  ],
  providers: [],
  bootstrap: [AppComponent],
})
export class AppModule {}
