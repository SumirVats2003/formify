import axios, { AxiosInstance } from "axios";
import { environment } from "../../environments/environment";

export class AxiosFactory {
  static readonly axiosInstance: AxiosInstance = axios.create({
    baseURL: environment.serverUrl
  });

  static getAxiosInstance() {
    return this.axiosInstance;
  }
}