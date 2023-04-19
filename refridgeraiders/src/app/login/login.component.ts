import { Component, OnInit } from '@angular/core';
import { DataService } from '../data.service';
import { Router } from '@angular/router';
import { FormBuilder, FormGroup, Validators, FormControlState } from '@angular/forms';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {
  registerForm:any = FormGroup;
  submitted = false;
  invalid = false;
  loginUserData = {
    user: "",
    password: "",
    firstN: "",
    lastN: ""
  }
  constructor(
    private dataService: DataService,
    private router: Router,
    private formBuilder: FormBuilder) 
    { }
  get f() { return this.registerForm.controls; }
  onSubmit() {
    this.submitted = true;
    if (this.registerForm.invalid) {
      return;
    }
    alert('SUCCESS!! :-)\n\n' + JSON.stringify(this.registerForm.value))
    if (this.registerForm.valid) {
      this.loginUser();
    }
    if (this.submitted) {
      this.redirect();
    }
    if(this.dataService.faultyLogin == true) {
      alert("Invalid username or password");
    }
  }
  ngOnInit() {
    this.registerForm = this.formBuilder.group({
      email: ['', [Validators.required, Validators.email]],
      password: ['', [Validators.required, Validators.minLength(6)]]
    });
  }

  loginUser() {
    this.dataService.loginUser(this.loginUserData)
      .subscribe(
        res => console.log(res),
        err => console.log(err)
      )
      if(this.dataService.faultyLogin == true) {
        this.invalid = true;
      }
  }
  
  redirect() {
    this.router.navigate(['/register']);
  }
}
