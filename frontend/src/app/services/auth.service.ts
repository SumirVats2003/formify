import { Injectable } from '@angular/core';
import { AxiosInstance } from 'axios';
import { AxiosFactory } from '../factory/axios.factory';

@Injectable({
  providedIn: 'root'
})
export class AuthService {
  axios: AxiosInstance = AxiosFactory.getAxiosInstance();

  loginUser(email: string, password: string) {
    try {
      const result = this.axios.get("/auth/login", {
        params: {
          email, password
        }
      })
      return result;
    } catch (error) {
      console.log(error);
      return null;
    }
  }
}
