package v1

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	l "github.com/note_project/pkg/logger"
	"github.com/note_project/pkg/structures"
)

//@Summary Set Note
//@Description This api for Seting Notes
//@Tags Note
//@Accept json
//@Produce json
//@param exptime query int true "exptime"
//@Param note body NoteStruct true "note body"
//@Success 200 {string} Success
//@Router /v1/notes [post]
func (h handlerV1) SetNoteWithTTL(c *gin.Context) {
	var note NoteStruct

	exptime := c.Query("exptime")

	exp_time, _ := strconv.ParseInt(exptime, 10, 64)

	err := c.ShouldBindJSON(&note)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while getting note", l.Error(err))
		return
	}

	id, err := uuid.NewV4()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while genereting uuid", l.Error(err))
		return
	}
	note.ID = id.String()

	noteRedis, err := json.Marshal(note)

	err = h.redisStorage.SetWithTTL(note.Title, string(noteRedis), exp_time)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while setting note", l.Error(err))
		return
	}
	c.JSON(http.StatusOK, note)

}

//@Summary Create Note
//@Description This api for Creating Note
//@Tags Note
//@Accept json
//@Produce json
//@Param note body NoteStruct true "note body"
//@Success 200 {string} Success
//@Router /v1/createnote [post]
func (h handlerV1) CreateNote(c *gin.Context) {
	var note structures.NoteStruct

	err := c.ShouldBindJSON(&note)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while creating note", l.Error(err))
		return
	}

	noteBody, err := h.postgresStorage.CreateNote(&note)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while creating note", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, noteBody)
}

//@Summary Update Note
//@Description This api for Updating Note
//@Tags Note
//@Accept json
//@Produce json
//@Param note body NoteStruct true "note body"
//@Success 200 {string} Success
//@Router /v1/updatenote [put]
func (h handlerV1) UpdateNote(c *gin.Context) {
	var note structures.NoteStruct

	err := c.ShouldBindJSON(&note)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while updating note", l.Error(err))
		return
	}

	noteBody, err := h.postgresStorage.UpdateNote(&note)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while updating note", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, noteBody)
}

//@Summary Delete Note
//@Description This api for deleting note
//@Tags Note
//@Accept json
//@Produce json
//@Param id path string true "Note ID"
//@Success 200 {string} Success
//@Router /v1/deletenote/{id} [delete]
func (h handlerV1) DeleteNote(c *gin.Context) {
	id := c.Param("id")
	err := h.postgresStorage.DeleteNote(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while deleting note", l.Error(err))
	}

	c.JSON(http.StatusOK, "Deleted")
}
