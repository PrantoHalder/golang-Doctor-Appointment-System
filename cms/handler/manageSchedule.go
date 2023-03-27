package handler

import (
	"fmt"
	"net/http"
)

func (h Handler)MangeSchedule(w http.ResponseWriter, r *http.Request){
	fmt.Println("inside manage schedule ")
}