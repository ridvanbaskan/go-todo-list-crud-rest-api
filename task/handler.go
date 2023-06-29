package task

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"math"
	"net/http"
	"strconv"
	"strings"
	"todo-list/db"
	"todo-list/pagination"
	user "todo-list/user"
)

type Handler struct{}

func NewTaskHandler() *Handler {
	return &Handler{}
}

func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var userID = r.Context().Value("userID").(uint)

	var currentUser user.User
	userRes := db.Db.First(&currentUser, userID)

	fmt.Println(currentUser)
	if userRes.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("User not found")
		return
	}

	var task Task
	_ = json.NewDecoder(r.Body).Decode(&task)

	result := db.Db.Create(&task)

	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(result.Error)
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func (h *Handler) GetTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	taskID, err := strconv.Atoi(params["id"])

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid task ID")
		return
	}

	var task Task
	result := db.Db.First(&task, taskID)

	if result.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Task not found")
		return
	}

	json.NewEncoder(w).Encode(task)
}

func (h *Handler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	taskID, err := strconv.Atoi(params["id"])

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid task ID")
		return
	}

	var task Task

	result := db.Db.Delete(&task, taskID)

	if result.Error != nil || result.RowsAffected == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Task not found")
	}

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode("Task Deleted Successfully")
}

func (h *Handler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	taskID, err := strconv.Atoi(params["id"])

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid task ID")
		return
	}

	var task Task
	result := db.Db.First(&task, taskID)

	if result.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Task not found")
		return
	}

	_ = json.NewDecoder(r.Body).Decode(&task)

	db.Db.Save(&task)

	json.NewEncoder(w).Encode(task)
}

func (h *Handler) MarkTaskAs(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	taskID, err := strconv.Atoi(params["id"])

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid Task ID")
		return
	}

	var task Task
	result := db.Db.First(&task, taskID)

	if result.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Task not found")
		return
	}

	var completed bool
	_ = json.NewDecoder(r.Body).Decode(&completed)

	task.Completed = completed
	db.Db.Save(&task)

	json.NewEncoder(w).Encode(task)

}

func (h *Handler) GetTasks(w http.ResponseWriter, r *http.Request) {
	dueDate := r.URL.Query().Get("dueDate")
	priority := r.URL.Query().Get("priority")
	sortBy := r.URL.Query().Get("sortBy")

	page, offset, limit := pagination.GetPaginationOffsetAndLimit(r)

	query := db.Db.Model(&Task{})

	if dueDate != "" {
		query = query.Where("due_date = ?", dueDate)
	}

	if priority != "" {
		query = query.Where("priority = ?", priority)
	}

	var sortField, sortOrder string

	if strings.Contains(sortBy, ".") {
		split := strings.Split(sortBy, ".")
		sortField = split[0]
		sortOrder = split[1]
	} else {
		sortField = sortBy
		sortOrder = "asc"
	}

	switch sortField {
	case "dueDate":
		query = query.Order("due_date " + sortOrder)
	case "priority":
		query = query.Order("priority " + sortOrder)

	}

	var totalRecords int64
	query.Count(&totalRecords)

	var tasks []Task

	result := query.Offset(offset).Limit(limit).Find(&tasks)

	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(result.Error)
		return
	}

	totalPages := int(math.Ceil(float64(totalRecords) / float64(limit)))

	response := pagination.Pagination[Task]{
		TotalRecords: totalRecords,
		PageSize:     limit,
		CurrentPage:  page,
		TotalPages:   totalPages,
		Data:         tasks,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

}
