import { Component } from '@angular/core';
import { DataService } from '../data.service';
import { Router } from '@angular/router';
import { LoginComponent } from '../login/login.component';

@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.css']
})
export class ProfileComponent {
  user = {
    user: "",
    password: "",
    firstN: "",
    lastN: ""
  };
  constructor(private dataService: DataService, private login: LoginComponent) { }
  ngOnInit() {
    const getUser = {
      user: "",
      password: ""
    }
    console.log(this.login.loginUserData)
    this.dataService.getUser(this.login.loginUserData).subscribe((response: any) => {
        console.log(response);
        this.user = response;
      }, (error) => {
        console.error(error);
      });
  }
}
