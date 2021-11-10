import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { ICustomer } from '../models/customer';

@Injectable({
  providedIn: 'root'
})
export class UserService {

  constructor(private _http:HttpClient) { }
  
  GetAllUsers():Observable<ICustomer[]>{
    let cust = this._http.get<ICustomer[]>('http://localhost:8080/api/getData/');
    return cust;
  }
  PostUser(cust:ICustomer):Observable<number>{
    let customer=this._http.post<number>('http://localhost:8080/api/postData/',cust);
    return customer;
  }
  DeleteUser(id:number):Observable<number>{
    let customer=this._http.delete<number>('http://localhost:8080/api/deleteRow/'+id);
    return customer;
  }

}
