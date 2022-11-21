package main

import (
	"errors"
	"net/http"
	"privy/internal/data"
	"privy/internal/dto"
	"privy/internal/validator"
)

func (app *application) GetAllCakes(w http.ResponseWriter, r *http.Request) {
	cakes, err := app.repos.Cakes.GetAll()
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	err = app.writeJSON(w, 200, envelope{"cakes": cakes}, nil)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) GetCakeByID(w http.ResponseWriter, r *http.Request) {
	id, err := app.getIDParam(r)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	res, err := app.repos.Cakes.GetOneByID(id)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	err = app.writeJSON(w, 200, envelope{"cake": res}, nil)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) CreateCake(w http.ResponseWriter, r *http.Request) {
	var cakeInsert dto.CakeInsert
	err := app.readJSON(w, r, &cakeInsert)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	cake := &data.Cake{
		Title:       cakeInsert.Title,
		Description: cakeInsert.Description,
		Rating:      cakeInsert.Rating,
		Image:       cakeInsert.Image,
	}

	v := validator.New()

	data.ValidateCake(v, cake)
	if !v.Valid() {
		app.errorJSON(w, errors.New("invalid input"), http.StatusUnprocessableEntity)
		return
	}

	err = app.repos.Cakes.Insert(cake)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	err = app.writeJSON(w, http.StatusCreated, envelope{"message": "user is created succesfully"}, nil)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) UpdateCakeByID(w http.ResponseWriter, r *http.Request) {
	id, err := app.getIDParam(r)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	var cakeUpdate dto.CakeUpdate
	err = app.readJSON(w, r, &cakeUpdate)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	current, err := app.repos.Cakes.GetOneByID(id)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	if cakeUpdate.Title != nil {
		current.Title = *cakeUpdate.Title
	}
	if cakeUpdate.Description != nil {
		current.Description = *cakeUpdate.Description
	}
	if cakeUpdate.Rating != nil {
		current.Rating = *cakeUpdate.Rating
	}
	if cakeUpdate.Image != nil {
		current.Image = *cakeUpdate.Image
	}

	v := validator.New()

	data.ValidateCake(v, current)
	if !v.Valid() {
		app.errorJSON(w, errors.New("invalid input"), http.StatusUnprocessableEntity)
		return
	}

	err = app.repos.Cakes.UpdateByID(current)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"message": "cake updated sucesfully"}, nil)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) DeleteCake(w http.ResponseWriter, r *http.Request) {
	id, err := app.getIDParam(r)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	err = app.repos.Cakes.DeleteByID(id)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"message": "cake deleted sucesfully"}, nil)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}
