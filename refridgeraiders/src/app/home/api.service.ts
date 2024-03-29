import { Injectable, EventEmitter } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { tap } from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class ApiService {

  private apiUrl = 'https://api.edamam.com';
  private app_id = '9af6c883';
  private app_key = 'c7cf201d3b30404c49c74054a66b9345';

  constructor(private http: HttpClient) { }

  public dataEmitter = new EventEmitter<any>();

  getData(input: string): Observable<any> {
    const url = `https://api.edamam.com/search?q=${input}&app_id=${this.app_id}&app_key=${this.app_key}`;
    return this.http.get<any>(url).pipe(
      tap(data => this.dataEmitter.emit(data))
    );
  }

}
