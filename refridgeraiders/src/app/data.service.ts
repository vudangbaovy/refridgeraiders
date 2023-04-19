import { Injectable } from '@angular/core';
import { HttpClient, HttpResponse } from '@angular/common/http';
import { Observable, throwError } from 'rxjs';
import { catchError, retry, map } from 'rxjs/operators';
import { Router } from '@angular/router';
import { User } from './user';

@Injectable({
  providedIn: 'root'
})

export class DataService {
  private registerUrl = 'http://localhost:3000/user/register'
  private loginUrl = 'http://localhost:3000/user';
  private getUserUrl = 'http://localhost:3000/login';

  isLoggedIn: boolean = false;
  isRegistered: boolean = false;
  faultyLogin: boolean = false;
  public redirectUrl: string = '/';
  constructor(
    private http: HttpClient,
    private router: Router
    ) { }

  registerUser(user: any) {
    return this.http.post<any>(this.registerUrl, user).pipe(map((response: any) => {
      // do whatever with your response
      this.isRegistered = true;
      if (this.redirectUrl) {
        this.router.navigate(['login']);
        this.redirectUrl = '/';
      }
    }));
  }

  // loginUser(user: User) {
  //   return this.http.post<any>(this.loginUrl, user)
  // }
  loginUser(user: any): Observable<any> {
    return this.http.post(this.loginUrl, user).pipe(map((response: any) => {
      // do whatever with your response
      if (response.status === 200) {
        this.isLoggedIn = true;
        if (this.redirectUrl) {
          this.router.navigate([this.redirectUrl]);
          this.redirectUrl = '/';
        }
      }
      else {
        this.isLoggedIn = false;
        this.faultyLogin = true;
      }
    }));
  }

  logout() {
    this.isLoggedIn = false;
    this.router.navigate(['login']);
  }

  getUser() {
    return this.http.post(this.getUserUrl, {token: localStorage.getItem('token')});
  }
}
