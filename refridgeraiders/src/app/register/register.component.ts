import { Component, OnInit } from '@angular/core';
import { DataService } from '../data.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css']
})
export class RegisterComponent implements OnInit {
  registerUserData = {
    firstN: "",
    lastN: "",
    user: "",
    password: ""
  }
  constructor(private dataService: DataService, private router: Router) { }
  ngOnInit() {
  }
  registerUser() {
    this.dataService.registerUser(this.registerUserData)
      .subscribe(
        res => console.log(res),
        err => console.log(err)
      )
      this.router.navigate(['/login']);
  }
}