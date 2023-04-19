import { Component, OnInit } from '@angular/core';
import { ApiService } from '../home/api.service';
import { RecipeService } from '../shared/recipe.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-recipe',
  templateUrl: './recipe.component.html',
  styleUrls: ['./recipe.component.css']
})
export class RecipeComponent implements OnInit{

  public selectedRecipeData: any;
  public selectedRecipeIndex: number;

  constructor(private apiService: ApiService, private recipeService: RecipeService) {
     this.selectedRecipeIndex = this.recipeService.selectedRecipeIndex;
   }

   
  ngOnInit() {
    this.selectedRecipeIndex = this.recipeService.selectedRecipeIndex;
    this.selectedRecipeData = this.recipeService.selectedRecipeData;
    console.log(this.selectedRecipeData);
  }
}
