import { Component, Injectable } from '@angular/core';
import { ApiService } from './api.service';
import { Router } from '@angular/router';
import { RecipeService } from '../shared/recipe.service';
@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css'],
})  

export class HomeComponent {
  title = "project";
  searchterm: string = '';
  data: any;

  constructor(private apiService: ApiService, private router: Router, private recipeService: RecipeService) { }

  navigateToNewComponent(recipeIndex: number): void{
    this.recipeService.selectedRecipeIndex = recipeIndex;
    this.recipeService.selectedRecipeData = this.data.hits[recipeIndex];
    this.router.navigate(['/recipe']);
  }
  
    public fetchData(){

      this.apiService.getData(this.searchterm)
      .subscribe((data) => {
        this.data = data;
        console.log(data);
      });

    }
    public buildURLInternal(input: string, app_id: string, app_key: string) {
         let result =  `https://api.edamam.com/search?q=${input}&app_id=${app_id}&app_key=${app_key}`;
         console.log(result);
         return result;
    }

}

export function buildURLInternal(input: string, app_id: string, app_key: string) {}
