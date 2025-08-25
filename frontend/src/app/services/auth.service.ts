import { Injectable } from '@angular/core';
import { AxiosInstance } from 'axios';
import { AxiosFactory } from '../factory/axios.factory';

@Injectable({
  providedIn: 'root'
})
export class AuthService {
  axios: AxiosInstance;

  constructor() {
    this.axios = AxiosFactory.getAxiosInstance();
  }

  loginUser(email: string, password: string) {
    try {
      const result = this.axios.post("/auth/login",
        {
          email, password
        }
      )
      console.log(result)
    } catch (error) {
      console.log(error);
    }
  }
}
