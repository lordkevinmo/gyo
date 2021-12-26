package render

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRender_Page(t *testing.T) {
	r, err := http.NewRequest(http.MethodGet, "/random-url", nil)
	if err != nil {
		t.Error(err)
	}

	w := httptest.NewRecorder()

	testRenderer.RootPath = "./testData"

	shouldThrowErrorWhenTheTemplateEngineIsSetWrongly(w, r, t)

	shouldNotReturnErrorWhenRendererIsGoTemplate(w, r, t)
	shouldReturnErrorWhenRendererIsGoTemplateButTheFileDoesNotExists(w, r, t)

	shouldNotReturnErrorWhenRendererIsJetTemplate(w, r, t)
	shouldReturnErrorWhenRendererIsJetTemplateButTheFileDoesNotExists(w, r, t)
}

func shouldThrowErrorWhenTheTemplateEngineIsSetWrongly(w http.ResponseWriter, r *http.Request, t *testing.T) {
	testRenderer.Renderer = ""
	err := testRenderer.Page(w, r, "home", nil, nil)
	if err == nil {
		t.Error("No Error returned while rendering with invalid renderer specified", err)
	}
}

func shouldNotReturnErrorWhenRendererIsGoTemplate(w http.ResponseWriter, r *http.Request, t *testing.T) {
	testRenderer.Renderer = "go"
	err := testRenderer.Page(w, r, "home", nil, nil)
	if err != nil {
		t.Error("Error rendering Go template Page", err)
	}
}

func shouldReturnErrorWhenRendererIsGoTemplateButTheFileDoesNotExists(w http.ResponseWriter, r *http.Request, t *testing.T) {
	testRenderer.Renderer = "go"
	err := testRenderer.Page(w, r, "no-file", nil, nil)
	if err == nil {
		t.Error("Error rendering non-existent Go template page", err)
	}
}

func shouldNotReturnErrorWhenRendererIsJetTemplate(w http.ResponseWriter, r *http.Request, t *testing.T) {
	testRenderer.Renderer = "jet"
	err := testRenderer.Page(w, r, "home", nil, nil)
	if err != nil {
		t.Error("Error rendering Jet template Page", err)
	}
}

func shouldReturnErrorWhenRendererIsJetTemplateButTheFileDoesNotExists(w http.ResponseWriter, r *http.Request, t *testing.T) {
	testRenderer.Renderer = "jet"
	err := testRenderer.Page(w, r, "no-file", nil, nil)
	if err == nil {
		t.Error("Error rendering non-existent Jet template page", err)
	}
}
