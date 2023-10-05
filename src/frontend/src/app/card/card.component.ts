import { Component, Input } from '@angular/core';


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

}

@Component({
  selector: 'app-card2',
  templateUrl: './card2.component.html',
  styleUrls: ['./card2.component.css']
})
export class Card2Component {
  @Input()
  title!: string;
  @Input()
  imageSrc!: string;
  @Input()
  date!: string;
}
