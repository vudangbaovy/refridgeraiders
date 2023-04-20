import { TestBed, async } from '@angular/core/testing';
import { HttpClientTestingModule, HttpTestingController } from '@angular/common/http/testing';
import { DataService } from '../data.service';
import { ProfileComponent } from '../profile/profile.component';
import { of } from 'rxjs';

describe('ProfileComponent', () => {
  let component: ProfileComponent;
  let httpTestingController: HttpTestingController;
  let dataService: DataService;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      imports: [HttpClientTestingModule],
      providers: [DataService],
      declarations: [ProfileComponent]
    }).compileComponents();
  }));

  beforeEach(() => {
    httpTestingController = TestBed.inject(HttpTestingController);
    dataService = TestBed.inject(DataService);
    component = TestBed.createComponent(ProfileComponent).componentInstance;
  });

  it('should retrieve user profile on init', () => {
    const getUser = {
      user: 'vyvooz',
      password: '123hello'
    };
    const expectedUser = {id: 1, name: 'Vyvooz', email: 'vyvooz@test.com'};

    // Spy on the getUser method of the DataService and return the expected user
    spyOn(dataService, 'getUser').and.returnValue(of(expectedUser));

    component.ngOnInit();

    // Expect the getUser method to have been called with the getUser object
    expect(dataService.getUser).toHaveBeenCalledWith(getUser);

    // Expect the user property of the component to be set to the expected user
    expect(component.user).equal(expectedUser);
  });
});
