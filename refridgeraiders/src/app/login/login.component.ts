import { Component, OnInit } from '@angular/core';
import { DataService } from '../data.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {
  loginUserData = {
    name: "",
    password: "",
    adminlevel: 0,
    allergies: " "
  }
  constructor(private dataService: DataService) { }
  ngOnInit() {
  }
  loginUser() {
    console.log(this.loginUserData);
    this.dataService.loginUser(this.loginUserData)
      .subscribe(
        res => console.log(res),
        err => console.log(err)
      )
  }
}
