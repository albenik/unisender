package unisender_test

import (
	"github.com/alexeyco/unisender"
	"github.com/alexeyco/unisender/api"
	"log"
	"net/http"
	"testing"

	"github.com/alexeyco/unisender/test"
)

func TestUniSender_ApiKey(t *testing.T) {
	apiKeyExpected := test.RandomString(12, 36)
	var apiKeyRequested string

	c := test.NewClient(func(req *http.Request) (res *http.Response, err error) {
		apiKeyRequested = req.FormValue("api_key")

		res = &http.Response{
			StatusCode: http.StatusOK,
		}

		return
	})

	usndr := unisender.New(apiKeyExpected)
	usndr.SetClient(c)

	err := usndr.DeleteList(123).Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if apiKeyExpected != apiKeyRequested {
		t.Fatalf(`API key should be "%s", "%s" given`, apiKeyExpected, apiKeyRequested)
	}
}

func TestUniSender_Format(t *testing.T) {
	apiKeyExpected := test.RandomString(12, 36)

	formatExpected := "json"
	var formatRequested string

	c := test.NewClient(func(req *http.Request) (res *http.Response, err error) {
		formatRequested = req.FormValue("format")

		res = &http.Response{
			StatusCode: http.StatusOK,
		}

		return
	})

	usndr := unisender.New(apiKeyExpected)
	usndr.SetClient(c)

	err := usndr.DeleteList(123).Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if formatExpected != formatRequested {
		t.Fatalf(`Format should be "%s", "%s" given`, formatExpected, formatRequested)
	}
}

func ExampleUniSender_CreateList() {
	usndr := unisender.New("your-api-key")
	usndr.SetLanguage(api.LanguageItalian)

	listID, err := usndr.CreateList("Your list name").
		Execute()

	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("List created; ID=%d", listID)
}

func ExampleUniSender_GetLists() {

}

func ExampleUniSender_UpdateList() {

}

func ExampleUniSender_DeleteList() {

}

func ExampleUniSender_GetContact() {

}

func ExampleUniSender_Subscribe() {

}

func ExampleUniSender_Unsubscribe() {

}

func ExampleUniSender_Exclude() {

}

func ExampleUniSender_ImportContacts() {

}

func ExampleUniSender_ExportContacts() {

}

func ExampleUniSender_IsContactInList() {

}

func ExampleUniSender_GetContactCount() {

}

func ExampleUniSender_GetTotalContactsCount() {

}
