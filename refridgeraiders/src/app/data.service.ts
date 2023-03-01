import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable, throwError } from 'rxjs';
import { catchError, retry } from 'rxjs/operators';
import { User } from './user';

@Injectable({
  providedIn: 'root'
})

export class DataService {
  private loginUrl = 'http://localhost:4200/User';

  constructor(private http: HttpClient) { }
  loginUser(user: User) {
    return this.http.get<any>(this.loginUrl)
  }
}
