import { Injectable } from '@angular/core';
import {
  HttpRequest,
  HttpResponse,
  HttpHandler,
  HttpEvent,
  HttpInterceptor,
  HTTP_INTERCEPTORS,
} from '@angular/common/http';
import { Observable, of, throwError } from 'rxjs';
import { delay, mergeMap, materialize, dematerialize } from 'rxjs/operators';
import sign from 'jwt-encode';
import { LoginSuccess } from '@app/_services';

let users = [{ login: 'test', password: 'test', email: 'test@test.com' }];
let games = [
  {
    id: 'fake-uid-duh-joined',
    name: 'My Great Game (joined)',
    players: { joined: 2, required: 6 },
    status: 'Waiting for Players...',
    joined: true,
  },
  {
    id: 'fake-uid-duh-joinable',
    name: 'My Great Game (joinable)',
    players: { joined: 2, required: 6 },
    status: 'Waiting for Players...',
    joined: false,
  },
  {
    id: 'fake-uid-duh-full',
    name: 'My Great Game (Ready)',
    players: { joined: 6, required: 6 },
    status: 'Waiting for orders...',
    joined: true,
  },
];

@Injectable()
export class FakeBackendInterceptor implements HttpInterceptor {
  intercept(
    request: HttpRequest<any>,
    next: HttpHandler
  ): Observable<HttpEvent<any>> {
    const { url, method, headers, body } = request;

    // wrap in delayed observable to simulate server api call
    return of(null)
      .pipe(mergeMap(handleRoute))
      .pipe(materialize()) // call materialize and dematerialize to ensure delay even if an error is thrown (https://github.com/Reactive-Extensions/RxJS/issues/648)
      .pipe(delay(500))
      .pipe(dematerialize());

    function handleRoute() {
      switch (true) {
        case url.endsWith('/auth/signin') && method === 'POST':
          return authenticate();
        case url.endsWith('/join') && method === 'POST':
          return join();
        case url.endsWith('/auth/signup') && method === 'POST':
          return signup();
        case url.endsWith('/users') && method === 'GET':
          return getUsers();
        case url.endsWith('/games') && method === 'GET':
          return getGames();
        default:
          // pass through any requests not handled above
          return next.handle(request);
      }
    }

    // route functions

    function authenticate() {
      const { login, password } = body;
      const user = users.find(
        (x) => x.login === login && x.password === password
      );
      if (!user) return error('login or password is incorrect');
      const token = sign(
        {
          authorized: true,
          login: login,
          exp: new Date(Date.now() + 3600 * 1000 * 24),
        },
        'fake-secrets'
      );
      return ok({ token } as LoginSuccess);
    }

    function join() {
      if (!isLoggedIn()) return unauthorized();
      console.log('Join player');
      games[1].players.joined++;
      games[1].joined = true;
      return ok();
    }

    function signup() {
      const { login, email, password } = body;
      const user = users.find((x) => x.login === login);
      if (!!user) return error('User already exists');
      users.push({ login, email, password });
      return ok({
        message: 'User created',
      });
    }

    function getUsers() {
      if (!isLoggedIn()) return unauthorized();
      return ok(users);
    }

    function getGames() {
      if (!isLoggedIn()) return unauthorized();
      return ok(games);
    }

    // helper functions

    function ok(body?) {
      return of(new HttpResponse({ status: 200, body }));
    }

    function error(message) {
      return throwError({ error: { message } });
    }

    function unauthorized() {
      return throwError({ status: 401, error: { message: 'Unauthorised' } });
    }

    function isLoggedIn() {
      return headers.get('Authorization').startsWith('Bearer');
    }
  }
}

@Injectable()
export class TrueBackEndInterceptor implements HttpInterceptor {
  intercept(
    request: HttpRequest<any>,
    next: HttpHandler
  ): Observable<HttpEvent<any>> {
    return next.handle(request);
  }
}

export let fakeBackendProvider = {
  // use fake backend in place of Http service for backend-less development
  provide: HTTP_INTERCEPTORS,
  useClass: FakeBackendInterceptor,
  multi: true,
};
