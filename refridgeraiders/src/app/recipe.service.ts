import { Injectable } from '@angular/core';


@Injectable({
  providedIn: 'root'
})
export class RecipeService {
  public selectedRecipeIndex: number = -1;
  selectedRecipeData: any;

  constructor() { }
}
