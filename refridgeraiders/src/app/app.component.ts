import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { DataService } from './data.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  isLoggedIn: boolean = false;
  title = 'refridgeraiders';
  constructor(private router: Router, private dataService: DataService) {

  }
  clickButton(path: string) {
    this.router.navigate([path]);
  }
  logout() { 
    this.dataService.logout();
  }
  ngOnInit() {
    this.router.events.subscribe(event => {
      if (event.constructor.name === "NavigationEnd") {
       this.isLoggedIn = this.dataService.isLoggedIn;
      }
    })
  }
}