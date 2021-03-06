import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Router } from '@angular/router';
import { JwtHelperService } from '@auth0/angular-jwt';
import { BehaviorSubject, Observable } from 'rxjs';
import { map } from 'rxjs/operators';
import { LoggedUserDto } from 'src/app/interfaces/logged-user-dto';
import { NewUser } from 'src/app/interfaces/new-user';
import { environment } from 'src/environments/environment';

@Injectable({
  providedIn: 'root'
})
export class UserService {

  private currentUserSubject: BehaviorSubject<LoggedUserDto>;
  public currentUser: Observable<LoggedUserDto>;
  private user! : LoggedUserDto;
  private loginStatus = new BehaviorSubject<boolean>(false);
  private decoded! : any

  private apiServerUrl = environment.apiBaseUrl;
  constructor(private http: HttpClient, private router : Router, private jwtHelper :JwtHelperService) {
    this.currentUserSubject = new BehaviorSubject<LoggedUserDto>(
      JSON.parse(localStorage.getItem('currentUser')!)
    );
    this.currentUser = this.currentUserSubject.asObservable();
   }

   public get currentUserValue(): LoggedUserDto {
    return this.currentUserSubject.value;
}

public getUserValue() : LoggedUserDto {
  console.log("Token" + this.currentUserSubject.value.token.accessToken);
  return this.currentUserSubject.value;
}

registerUser(newUser: NewUser) {
return this.http.post(`${this.apiServerUrl}/api/auth/signup`, newUser, {
  responseType: 'text',
});
}

loggedIn(): boolean {
  const token = localStorage.getItem('token');
  return true;
}

get isLoggedIn() {
  return this.loginStatus.asObservable();
}

check2FAStatus(username : string) : Observable<any> {
  return this.http.get(`${this.apiServerUrl}/api/auth/two-factor-auth-status/` + username)
}

enable2FA(username : string, status : boolean) : Observable<any> {
  return this.http.put(`${this.apiServerUrl}/api/auth/two-factor-auth/`, {
    username,
    status
  })
}

sendMagicLink(username : string) : Observable<any> {
  return this.http.post(`${this.apiServerUrl}/api/auth/password-less-login/`, username)
}

checkPasswordlessToken(token : string) : Observable<any> {
  return this.http.get(`${this.apiServerUrl}/api/auth/password-less-login/${token}`).pipe(
    map((response: any) => {
      if (response && response.token) {
        this.loginStatus.next(true);
        localStorage.setItem('token', response.token.accessToken);
        localStorage.setItem('currentUser', JSON.stringify(response));
        localStorage.setItem('role' ,response.role)
        localStorage.setItem('username' ,response.username)
        this.currentUserSubject.next(response);
      }
      return this.user;
    })
  );
}

login(model: any): Observable<LoggedUserDto> {
  return this.http.post(`${this.apiServerUrl}/api/auth/login`, model)
  .pipe(
    map((response: any) => {
      if (response) {
        
        this.decoded = this.jwtHelper.decodeToken(response.token.accessToken)
        console.log(this.decoded);
        
        localStorage.setItem('token', response.token.accessToken);
        localStorage.setItem('currentUser', JSON.stringify(response));
        localStorage.setItem('role' , this.decoded.role)
        localStorage.setItem('username', this.decoded.sub)

      
        this.currentUserSubject.next(response);
      }
      return this.user;
    })
  );
}

logout() {
  this.loginStatus.next(false);
  localStorage.removeItem('token');
  localStorage.removeItem('role');
  localStorage.removeItem('currentUser');
  localStorage.removeItem('username');
  this.router.navigate(['/']);
}

sendCode(email: string): Observable<any> {
  return this.http.post<any>(`${this.apiServerUrl}/api/auth/send-code`, email);
}

resetPassword(newPassword: string): Observable<any> {
  return this.http.post<any>(`${this.apiServerUrl}/api/auth/reset-password`, {
    username: localStorage.getItem('usernamee'),
    newPassword: newPassword,
  });
}

checkCode(verCode: string): Observable<any> {
  return this.http.post<any>(`${this.apiServerUrl}/api/auth/check-code`, {
    username: localStorage.getItem('usernamee'),
    code: verCode,
  });
}

getUserInformation() : Observable<any>{
  return this.http.get(`${this.apiServerUrl}/api/user/user-info`);
}

changePassword(data: any) {
  return this.http.put(`${this.apiServerUrl}/api/user/change-password`, data)
}
}
