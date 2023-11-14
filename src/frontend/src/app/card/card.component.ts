import { Component, Input } from '@angular/core';
import { Router } from '@angular/router';


@Component({
  selector: 'app-card',
  templateUrl: './card.component.html',
  styleUrls: ['./card.component.css']
})

export class CardComponent {
  @Input()
  title!: string;
  @Input()
  imageSrc!: string;

  @Input()
  textColor: string = 'white';

  changeTextColor(color: string) {
    this.textColor = color;
  }
}

@Component({
  selector: 'app-bottominfo',
  templateUrl: './bottominfo.component.html',
  styleUrls: ['./bottominfo.component.css']
})

export class BottomInfoComponent{
  constructor(
    private router: Router){
  }
  
  navigateToAbout(){
    this.router.navigate(['about']);
  }
}

