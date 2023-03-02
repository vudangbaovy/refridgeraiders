import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable, throwError } from 'rxjs';
import { catchError, retry } from 'rxjs/operators';
import { User } from './user';

@Injectable({
  providedIn: 'root'
})

export class DataService {
  private registerUrl = 'http://localhost:3000/User/Register'
  private loginUrl = 'http://localhost:3000/User';

  constructor(private http: HttpClient) { }

  registerUser(user: User) {
    return this.http.post<any>(this.registerUrl, user)
  }

  loginUser(user: User) {
    return this.http.post<any>(this.loginUrl, user)
  }
}
