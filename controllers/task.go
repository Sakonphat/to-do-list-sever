package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sever/controllers/requests"
	"sever/models"
)

func CreateTask(c *gin.Context)  {

	user, exists := getUser(c)
	if !exists {
		return
	}

	request, requestErr := requests.GetCreateTaskRequest(c)
	if requestErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code" : http.StatusInternalServerError,
			"success" : false,
			"message" : requestErr.Error(),
		})
		return
	}

	validateErr := requests.ValidateCreateTaskRequest(request)
	if validateErr != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code" : http.StatusUnprocessableEntity,
			"success" : false,
			"message" : validateErr,
		})
		return
	}

	todo := models.Task{
		UserId: user.ID,
		Title: request.Title,
		Description: request.Description,
	}

	queryErr := models.CreateTask(&todo)
	if queryErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code" : http.StatusInternalServerError,
			"success" : false,
			"message" : queryErr.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code" : http.StatusOK,
		"success" : true,
		"message" : "Successful task creation.",
	})
	return
}

func GetTask(c *gin.Context)  {

	user, exists := getUser(c)
	if !exists {
		return
	}

	uuid := c.Param("uuid")
	
	task := models.Task{}
	queryErr := models.GetATaskByUuid(&task, uuid)
	if queryErr != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code" : http.StatusNotFound,
			"success" : false,
			"message" : queryErr.Error(),
		})
		return
	}

	if user.ID != task.UserId {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code" : http.StatusUnprocessableEntity,
			"success" : false,
			"message" : "Permission Denied.",
		})
		return
	}

	data := map[string]interface{}{
		"uuid" : task.Uuid,
		"title" : task.Title,
		"description" : task.Description,
		"created_at" : task.CreatedAt,
		"updated_at" : task.UpdatedAt,
	}

	c.JSON(http.StatusOK, gin.H{
		"code" : http.StatusOK,
		"success" : true,
		"message" : "Success.",
		"data" : data,
	})
	return
}

//List all todos
func GetAllTask(c *gin.Context) {

	user, exists := getUser(c)
	if !exists {
		return
	}

	var tasks []models.Task
	if err := models.GetAllTask(&tasks, user.ID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code" : http.StatusInternalServerError,
			"success" : true,
			"message" : err,
			"data" : nil,
		})
		return
	}

	data := makeData(&tasks)
	if data == nil {
		data = make([]map[string]interface{}, 0)
	}

	c.JSON(http.StatusOK, gin.H{
		"code" : http.StatusOK,
		"success" : true,
		"message" : "Success.",
		"data" : data,
	})
	return
}

func EditTask(c *gin.Context)  {

	user, exists := getUser(c)
	if !exists {
		return
	}

	request, requestErr := requests.GetEditTaskRequest(c)
	if requestErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code" : http.StatusInternalServerError,
			"success" : false,
			"message" : requestErr.Error(),
		})
		return
	}

	validateErr := requests.ValidateEditTaskRequest(request)
	if validateErr != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code" : http.StatusUnprocessableEntity,
			"success" : false,
			"message" : validateErr,
		})
		return
	}

	task := models.Task{}
	queryErr := models.GetATaskByUuid(&task, request.Uuid)
	if queryErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code" : http.StatusInternalServerError,
			"success" : false,
			"message" : queryErr.Error(),
		})
		return
	}

	if user.ID != task.UserId {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code" : http.StatusUnprocessableEntity,
			"success" : false,
			"message" : "Permission Denied.",
		})
		return
	}

	updateData := map[string]interface{}{
		"title" : request.Title,
		"description" : request.Description,
	}

	editErr := models.UpdateTaskByMany(&task, updateData)
	if editErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code" : http.StatusInternalServerError,
			"success" : false,
			"message" : editErr.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code" : http.StatusOK,
		"success" : true,
		"message" : "Success.",
	})
	return

}

func CompleteTask(c *gin.Context)  {

	user, exists := getUser(c)
	if !exists {
		return
	}

	uuid := c.Param("uuid")

	task := models.Task{}
	queryErr := models.GetATaskByUuid(&task, uuid)
	if queryErr != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code" : http.StatusNotFound,
			"success" : false,
			"message" : queryErr.Error(),
		})
		return
	}

	if user.ID != task.UserId {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code" : http.StatusUnprocessableEntity,
			"success" : false,
			"message" : "Permission Denied.",
		})
		return
	}

	updateErr := models.UpdateCompletedTask(&task)
	if updateErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code" : http.StatusInternalServerError,
			"success" : false,
			"message" : updateErr.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code" : http.StatusOK,
		"success" : true,
		"message" : "Success.",
	})
	return

}

func UndoTask(c *gin.Context)  {

	user, exists := getUser(c)
	if !exists {
		return
	}

	uuid := c.Param("uuid")

	task := models.Task{}
	queryErr := models.GetATaskByUuid(&task, uuid)
	if queryErr != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code" : http.StatusNotFound,
			"success" : false,
			"message" : queryErr.Error(),
		})
		return
	}

	if user.ID != task.UserId {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code" : http.StatusUnprocessableEntity,
			"success" : false,
			"message" : "Permission Denied.",
		})
		return
	}

	updateErr := models.UpdateUndoTask(&task)
	if updateErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code" : http.StatusInternalServerError,
			"success" : false,
			"message" : updateErr.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code" : http.StatusOK,
		"success" : true,
		"message" : "Success.",
	})
	return

}

func DeleteTask(c *gin.Context)  {

	user, exists := getUser(c)
	if !exists {
		return
	}

	uuid := c.Param("uuid")

	task := models.Task{}
	queryErr := models.GetATaskByUuid(&task, uuid)
	if queryErr != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code" : http.StatusNotFound,
			"success" : false,
			"message" : queryErr.Error(),
		})
		return
	}

	if user.ID != task.UserId {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code" : http.StatusUnprocessableEntity,
			"success" : false,
			"message" : "Permission Denied.",
		})
		return
	}

	updateErr := models.DeleteATask(&task)
	if updateErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code" : http.StatusInternalServerError,
			"success" : false,
			"message" : updateErr.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code" : http.StatusOK,
		"success" : true,
		"message" : "Success.",
	})
	return

}

func DeleteAllTask(c *gin.Context)  {

	user, exists := getUser(c)
	if !exists {
		return
	}

	task := models.Task{}
	deleteErr := models.DeleteTasks(&task, user.ID)
	if deleteErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code" : http.StatusInternalServerError,
			"success" : false,
			"message" : deleteErr.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code" : http.StatusOK,
		"success" : true,
		"message" : "Success.",
	})
	return

}

func getUser(c *gin.Context) (*models.User, bool) {

	value, exists := c.Get("user")
	if !exists {
		return nil, false
	}

	return value.(*models.User), true
}

func makeData(tasks *[]models.Task) []map[string]interface{} {

	var data []map[string]interface{}

	for _, task := range *tasks {
		temp := map[string]interface{}{
			"uuid" : task.Uuid,
			"title" : task.Title,
			"description" : task.Description,
			"is_completed" : task.IsCompleted,
		}
		data = append(data, temp)
	}

	return data
}