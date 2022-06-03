package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"github.com/Eshy10/go-bookstore/pkg/models"
	"github.com/Eshy10/go-bookstore/pkg/utils"
)

var Newbook models.Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
newBooks := models.GetAllBooks()
res, _ := json.Marshal(newBooks)
w.Header().Set("Content-Type", "application/json")
w.WriteHeader(http.StatusOK)
w.Write(res)
}