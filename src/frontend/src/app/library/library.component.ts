import {Component} from '@angular/core';
import {SocialAuthService, GoogleLoginProvider} from '@abacritt/angularx-social-login';
import {Router} from '@angular/router';
import { AppComponent } from '../app.component';
import { Card2Component } from '../card/card.component';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Component({
  selector: 'app-library',
  templateUrl: './library.component.html',
  styleUrls: ['./library.component.css']
})

export class LibraryComponent {

  list1: any[] = [];
  list2: any[] = [];
  private jsonUrl = 'assets/data.json';

  // cardlist = [
  //   {title: 'Duck',imageSrc: 'assets/duck.jpg', date:'2023-10-05'},
  //   {title: 'Budgie',imageSrc: 'assets/undulat.jpg', date:'2023-10-04'},
  // ]

  constructor(
    private router: Router, 
    public mainApp: AppComponent,
    public socialAuthService: SocialAuthService,
    private http: HttpClient) {
  }

  navigateToSpecies(imageId: string, imageName: string, imageDate: string): void {
    this.router.navigate(['species-page'], {
      queryParams: {
        imageId: encodeURIComponent(imageId),
        imageName: encodeURIComponent(imageName),
        imageDate: encodeURIComponent(imageDate),
      }
      });
  }

  getData(): Observable<any[]> {
    return this.http.get<any[]>(this.jsonUrl);
  }

  ngOnInit(): void {
    this.getData().subscribe((response) => {
      const data = response;
      this.list1 = data.find((item) => 'list1' in item)?.list1 || [];
    });
  }

}
