import { HomeComponent } from "./home.component";
import { HttpClientTestingModule } from "@angular/common/http/testing";
describe("HomeComponent", () => {
    it("mount", () => {
        cy.mount(HomeComponent);
    });
});

describe('Image Test', () => {
    it('displays the image', () => {
      cy.visit('http://localhost:4200/', { timeout: 10000 });
    })
  });