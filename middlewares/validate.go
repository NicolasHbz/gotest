package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// ValidationObject is the structure you should send as argument to the Validate() function
// Each attribute must receive a structure following this format: &YourStructure{}
type ValidationObject struct {
	Body   interface{}
	Params interface{}
	Query  interface{}
}

// Validate middleware. Just send the ValidationObject described above as argument
// and it will send an error response to the client if it's request is not valid
func Validate(model ValidationObject) gin.HandlerFunc {
	return func(c *gin.Context) {
		errors := make([]error, 0, 3)

		if errParams := areValidParams(c, model.Params); errParams != nil {
			errors = append(errors, errParams)
		}

		if errBody := isValidBody(c, model.Body); errBody != nil {
			errors = append(errors, errBody)
		}

		if errQuery := isValidQuery(c, model.Query); errQuery != nil {
			errors = append(errors, errQuery)
		}

		if len(errors) > 0 {
			sendError(c, errors)
			return
		}

		c.Next()
	}
}

func sendError(c *gin.Context, err []error) {
	var errMessages []string

	for _, v := range err {
		errMessage := strings.Split(v.Error(), "\n")
		for _, v := range errMessage {
			errMessages = append(errMessages, v)
		}
	}

	c.JSON(http.StatusBadRequest, gin.H{"errors": errMessages})
}

func isValidBody(c *gin.Context, body interface{}) error {
	if body == nil {
		return nil
	} else if err := c.ShouldBindJSON(body); err != nil {
		return err
	}
	return nil
}

func areValidParams(c *gin.Context, params interface{}) error {
	if params == nil {
		return nil
	} else if err := c.ShouldBindUri(params); err != nil {
		return err
	}
	return nil
}

func isValidQuery(c *gin.Context, query interface{}) error {
	if query == nil {
		return nil
	} else if err := c.ShouldBindQuery(query); err != nil {
		return err
	}
	return nil
}
