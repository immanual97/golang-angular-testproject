import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { UserService } from '../../services/user.service';

@Component({
  selector: 'app-add',
  templateUrl: './add.component.html',
  styleUrls: ['./add.component.css']
})
export class AddComponent implements OnInit {

  constructor(private userService:UserService,private router:Router) { }
 
  name=""
  email=""
  income=0
  ipaddress=""
  date=""

  id=Math.floor(Math.random()*100)
  ngOnInit(): void {
  }

  post(){
    var values={
      id:this.id,
      name:this.name,
      date:this.date,
      email:this.email,
      income:this.income,
      ipaddress:this.ipaddress
    }
    alert(values)
    this.userService.PostUser(values).subscribe(
      response=>{
        alert("Success")
        this.router.navigate(['/get'])
      }
    )
  }


}
