import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

import { environment } from '@environments/environment';
import { User } from '@app/_models';

export type RegisterUser = {
  login: string,
  email: string,
  password: string,
}


@Injectable({ providedIn: 'root' })
export class UserService {
    constructor(private http: HttpClient) { }

    getAll() {
        return this.http.get<User[]>(`${environment.apiUrl}/users`);
    }


    register(login: string, email: string, password: string) {
      return this.http.post<any>(`${environment.apiUrl}/auth/signup`, { login, email, password } as RegisterUser);
    }

}
