import { Component } from '@angular/core';
import { AuthService } from '../../services/auth.service';
import { FormControl, ReactiveFormsModule } from '@angular/forms';

@Component({
  selector: 'app-login',
  imports: [ReactiveFormsModule],
  templateUrl: './login.component.html',
  styleUrl: './login.component.css'
})
export class LoginComponent {
  email = new FormControl<string>("");
  password = new FormControl<string>("");

  constructor(
    public authService: AuthService
  ) { }

  loginUser() {
    try {
      console.log("this is the place I wanted to be")
      if (!this.email.value || !this.password.value) return;
      const result = this.authService.loginUser(this.email.value, this.password.value);
      console.log(result);
    } catch (error) {
      console.log(error);
    }
  }
}
