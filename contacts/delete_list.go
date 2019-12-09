package contacts

import (
	"strconv"

	"github.com/alexeyco/unisender/api"
)

// DeleteListRequest request to delete a list.
type DeleteListRequest struct {
	request *api.Request
}

// Execute sends request to UniSender API and returns result.
func (r *DeleteListRequest) Execute() error {
	return r.request.Execute("deleteList", nil)
}

// DeleteList returns request to delete a list.
//
// See https://www.unisender.com/en/support/api/partners/deletelist/
func DeleteList(request *api.Request, listID int64) *DeleteListRequest {
	request.Add("list_id", strconv.FormatInt(listID, 10))

	return &DeleteListRequest{
		request: request,
	}
}
