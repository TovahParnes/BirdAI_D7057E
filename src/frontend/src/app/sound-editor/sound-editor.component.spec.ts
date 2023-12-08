import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SoundEditorComponent } from './sound-editor.component';

describe('SoundEditorComponent', () => {
  let component: SoundEditorComponent;
  let fixture: ComponentFixture<SoundEditorComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [SoundEditorComponent]
    });
    fixture = TestBed.createComponent(SoundEditorComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
