package startnewgame

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/matthiashermsen/kaboom/api/response"
	"github.com/matthiashermsen/kaboom/storage"
)

func Handle(store storage.Store, logger *slog.Logger) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		var requestBody RequestBody

		decoder := json.NewDecoder(request.Body) // TODO was alles auslagern?
		err := decoder.Decode(&requestBody)

		if err != nil {
			responseWriter.WriteHeader(http.StatusBadRequest)
			response.WriteRequestBodyInvalidJSONResponse(responseWriter, logger)

			return
		}

		err = store.CreateNewBoard(requestBody.AmountOfRows, requestBody.AmountOfColumns, requestBody.AmountOfMines)

		if err != nil {
			apiResponse := response.NewFailureAPIResponse("!!! COMMAND FAILED !!!", "TODO") // !!! TODO !!!
			response.WriteJSONResponse(responseWriter, apiResponse, logger)

			return
		}

		// !!! TODO -- FOR TESTING PURPOSES -- !!!

		board, err := store.GetBoard()

		if err != nil {
			responseWriter.WriteHeader(http.StatusInternalServerError)
			// TODO ...

			return
		}

		encodedBoard, err := json.Marshal(board)

		if err != nil {
			responseWriter.WriteHeader(http.StatusInternalServerError)
			// TODO ...

			return
		}

		stringifiedBoard := string(encodedBoard)

		apiResponse := response.NewSuccessAPIResponse(stringifiedBoard)
		response.WriteJSONResponse(responseWriter, apiResponse, logger)
	}
}
