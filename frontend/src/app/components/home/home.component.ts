import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ICustomer } from '../../models/customer';
import { UserService } from '../../services/user.service';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit {

  allcustomers!:ICustomer[];

  constructor(private _userService: UserService,private router:Router) { }

  ngOnInit(): void {
    this.Get_AllCustomer();
  }

  Get_AllCustomer(){
    this._userService.GetAllUsers().subscribe(
      response=>{
        this.allcustomers=response;
      }
    )
  }

  deleteClick(id:any){
    if(confirm('Are you sure?')){
      this._userService.DeleteUser(id).subscribe(
        response=>{
          alert("Deleted")
          this.Get_AllCustomer();
        }
      )
    }
  }

  p:number = 1;
}
