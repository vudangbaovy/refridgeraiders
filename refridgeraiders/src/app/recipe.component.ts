import { Component, OnInit } from '@angular/core';
import { ApiService } from './home/api.service';
import { RecipeService } from './recipe.service';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-recipe',
  templateUrl: './recipe.component.html',
  styleUrls: ['./recipe.component.css']
})
export class RecipeComponent implements OnInit{

  public selectedRecipeData: any;
  public selectedRecipeIndex: number;
  
  constructor(private apiService: ApiService, private recipeService: RecipeService, private http: HttpClient) {
     this.selectedRecipeIndex = this.recipeService.selectedRecipeIndex;
   }

   
  ngOnInit() {
    this.selectedRecipeIndex = this.recipeService.selectedRecipeIndex;
    this.selectedRecipeData = this.recipeService.selectedRecipeData;
    console.log(this.selectedRecipeData);
  }
   sendStringToBackend(data: string) {
    const url = 'http://localhost:3000/bookmark';
    return this.http.post(url, { data });
  }
  
  updateStringInBackend(data: string) {
    const url = `http://localhost:3000/bookmark/${data}`;
    return this.http.put(url, { data });
  }
  
  deleteStringFromBackend(data: string) {
    const url = `http://localhost:3000/bookmark/${data}`;
    return this.http.delete(url);
  }
  
  
}
