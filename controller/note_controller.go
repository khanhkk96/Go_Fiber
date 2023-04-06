package controller

import (
	"go-fiber-crud/data/request"
	"go-fiber-crud/data/response"
	"go-fiber-crud/helper"
	"go-fiber-crud/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type NoteController struct {
	noteService service.NoteService
}

func NewNoteController(service service.NoteService) *NoteController {
	return &NoteController{
		noteService: service,
	}
}

func (controller *NoteController) Create(ctx *fiber.Ctx) error {
	createNoteRequest := request.CreateNoteRequest{}
	err := ctx.BodyParser(&createNoteRequest)
	helper.ErrorPanic(err)

	controller.noteService.Create(createNoteRequest)

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully created notes data!",
		Data:    nil,
	}
	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}

func (controller *NoteController) Update(ctx *fiber.Ctx) error {
	updateNoteRequest := request.UpdateNoteRequest{}
	err := ctx.BodyParser(updateNoteRequest)
	helper.ErrorPanic(err)

	noteId := ctx.Params("noteId")
	id, err := strconv.Atoi(noteId)
	helper.ErrorPanic(err)

	updateNoteRequest.Id = id
	controller.noteService.Update(updateNoteRequest)

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully updated notes data!",
		Data:    nil,
	}

	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}

func (controller *NoteController) Delete(ctx *fiber.Ctx) error {
	noteId := ctx.Params("noteId")
	id, err := strconv.Atoi(noteId)
	helper.ErrorPanic(err)

	controller.noteService.Delete(id)

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully deleted note data!",
		Data:    nil,
	}
	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}

func (controller *NoteController) FindById(ctx *fiber.Ctx) error {
	noteId := ctx.Params("noteId")
	id, err := strconv.Atoi(noteId)
	helper.ErrorPanic(err)

	noteReponse := controller.noteService.FindById(id)

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully get note data by id!",
		Data:    noteReponse,
	}
	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}

func (controller *NoteController) FindAll(ctx *fiber.Ctx) error {
	noteReponse := controller.noteService.FindAll()
	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully get all notes!",
		Data:    noteReponse,
	}
	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}
