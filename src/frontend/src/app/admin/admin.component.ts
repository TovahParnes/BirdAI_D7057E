import { Component } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { environment } from 'src/environments/environment';
import { AnalyzeResponse,DeleteResponse,AdminResponse } from 'src/assets/components/components';

@Component({
  selector: 'app-admin',
  templateUrl: './admin.component.html',
  styleUrls: ['./admin.component.css']
})
export class AdminComponent {

constructor(
  private http: HttpClient,
){}

  createAdmin(_id:string,access:string,userId:string){
    const authKey = localStorage.getItem("auth");
    if(authKey){
      this.sendCreateAdmin(authKey,_id,access,userId).subscribe(
        (response: AnalyzeResponse) => {
        }
      )
    }
  }

  makeMeAdmin(){
    const authKey = localStorage.getItem("auth");
    if(authKey){
      this.sendMakeMeAdmin(authKey).subscribe(
        (response: AdminResponse) => {
        }
      )
    }
  }

  updateAdmin(_id:string,access:string,userId:string){
    const authKey = localStorage.getItem("auth");
    if(authKey){
      this.sendUpdateAdmin(authKey,_id,access,userId).subscribe(
        (response: AdminResponse) => {
        }
      )
    }
  }

  deleteAdmin(userId:string){
    const authKey = localStorage.getItem("auth");
    if(authKey){
      this.sendDeleteAdmin(authKey,userId).subscribe(
        (response: DeleteResponse) => {
        }
      )
    }
  }

  getCurrentAdmin(){
    const authKey = localStorage.getItem("auth");
    if(authKey){
      this.sendGetCurrentAdmin(authKey).subscribe(
        (response: AdminResponse) => {
          console.log(response);
        }
      )
    }
  }

  getAdminByID(adminID:string){
    const authKey = localStorage.getItem("auth");
    if(authKey){
      this.sendGetAdminByID(authKey,adminID).subscribe(
        (response: AdminResponse) => {
          console.log(response);
        }
      )
    }
  }

  listSetOfAdmins(setOfAdmins:string,searchParameterForAdmin:string){
    const authKey = localStorage.getItem("auth");
    if(authKey){
      this.sendListSetOfAdmins(authKey,setOfAdmins,searchParameterForAdmin).subscribe(
        (response: AdminResponse) => {
          console.log(response);
        }
      )
    }
  }

  sendListSetOfAdmins(token:string,setOfAdmin:string,searchParameterForAdmin:string){
    const header = {
      'Authorization': `Bearer ${token}`
    };
    const request = "?set="+setOfAdmin+"&search="+searchParameterForAdmin
    return this.http.get<AdminResponse>(environment.identifyRequestURL+"/admins/list"+request,{ headers: header });
  }

  sendGetAdminByID(token:string,adminID:string){
    const header = {
      'Authorization': `Bearer ${token}`
    };
    return this.http.get<AdminResponse>(environment.identifyRequestURL+"/admins/"+adminID,{ headers: header });
  }

  sendGetCurrentAdmin(token:string){
    const header = {
      'Authorization': `Bearer ${token}`
    };
    return this.http.get<AdminResponse>(environment.identifyRequestURL+"/admins/me",{ headers: header });
  }

  sendUpdateAdmin(token:string,_id: string, access:string, userId:string){
    const header = {
      'Authorization': `Bearer ${token}`
    };
    const body = {
      "_id": _id,
      "access": access,
      "userId": userId,
    }
    return this.http.patch<AdminResponse>(environment.identifyRequestURL+"/admins/"+userId,body,{ headers: header });
  }

  sendMakeMeAdmin(token:string){
    const header = {
      'Authorization': `Bearer ${token}`
    };
    //Not listed in documentation but this call requires a body the content of which has no effect.
    const body = {
      "dummy": "dummy",
    }
    return this.http.post<AdminResponse>(environment.identifyRequestURL+"/admins/me",body,{ headers: header });
  }

  sendCreateAdmin(token: string, _id: string, access:string, userId:string){
    const header = {
      'Authorization': `Bearer ${token}`
    };
    const body = {
      "_id": _id,
      "access": access,
      "userId": userId,
    }
    return this.http.post<AnalyzeResponse>(environment.identifyRequestURL+"/admins/",body,{ headers: header });
  }

  sendDeleteAdmin(token:string,userId:string){
    const header = {
      'Authorization': `Bearer ${token}`
    };
    const body = {
      "userId": userId,
    };
    return this.http.delete<DeleteResponse>(environment.identifyRequestURL+"/admins/"+userId,{ headers: header });
  }

  createInput(request:string){
    const createAdmin_id = document.getElementById("createAdmin_id") as HTMLInputElement;
    const createAdminAccess = document.getElementById("createAdminAccess") as HTMLInputElement;
    const createAdminUserId = document.getElementById("createAdminUserId") as HTMLInputElement;
    const deleteAdminUserId = document.getElementById("deleteAdminUserId") as HTMLInputElement;
    const getAdminId = document.getElementById("getAdminByID") as HTMLInputElement;
    const getSetOfAdmins = document.getElementById("getSetOfAdmins") as HTMLInputElement;
    const getSearchParameterForAdmin = document.getElementById("getSearchParameterForAdmins") as HTMLInputElement;
    if(request == "createAdmin"){
      this.createAdmin(createAdmin_id.value,createAdminAccess.value,createAdminUserId.value)
      this.clearModifyInputs(createAdmin_id,createAdminAccess,createAdminUserId);
    }else if(request == "updateAdmin"){
      this.updateAdmin(createAdmin_id.value,createAdminAccess.value,createAdminUserId.value)
      this.clearModifyInputs(createAdmin_id,createAdminAccess,createAdminUserId);
    }else if(request == "deleteAdmin"){
      this.deleteAdmin(deleteAdminUserId.value)
      deleteAdminUserId.value ='';
    }else if(request == "getAdminByID"){
      this.getAdminByID(getAdminId.value);
      this.clearGetInputs(getAdminId,getSetOfAdmins,getSearchParameterForAdmin);
    }else if(request == "getListSetOfAdmins"){
      this.listSetOfAdmins(getSetOfAdmins.value,getSearchParameterForAdmin.value);
      this.clearGetInputs(getAdminId,getSetOfAdmins,getSearchParameterForAdmin);
    }else if(request == "getCurrentAdmin"){
      this.getCurrentAdmin();
      this.clearGetInputs(getAdminId,getSetOfAdmins,getSearchParameterForAdmin);
    }
  }

  clearModifyInputs(createAdmin_id:HTMLInputElement,createAdminAccess:HTMLInputElement,createAdminUserId:HTMLInputElement){
    createAdmin_id.value = '';
    createAdminAccess.value = '';
    createAdminUserId.value = '';
  }

  clearGetInputs(getAdminId:HTMLInputElement,getSetOfAdmins:HTMLInputElement,getSearchParameterForAdmin:HTMLInputElement){
    getSetOfAdmins.value = '';
    getSearchParameterForAdmin.value = '';
    getAdminId.value = '';
  }

  ngOnInit(): void {
    const userId = localStorage.getItem("userId");
    console.log(userId)
  }
  
}
