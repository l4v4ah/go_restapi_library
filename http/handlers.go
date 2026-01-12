package http

import (
	"encoding/json"
	"fmt"
	"library/lib"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type HTTPHandlers struct {
	library *lib.Lib
}

func NewHTTPHAndlers(l *lib.Lib) *HTTPHandlers {
	return &HTTPHandlers{
		library: l,
	}
}

// pattern: /library
// method: POST
// info: JSON
func (h *HTTPHandlers) HandleCreateNewBook(w http.ResponseWriter, r *http.Request) {
	var bDTO BookDTO
	if err := json.NewDecoder(r.Body).Decode(&bDTO); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("errL:", err)
		return
	}
	b := lib.NewBook(bDTO.Name, bDTO.Author, bDTO.Pages)
	if err := h.library.AddBook(b); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("errL:", err)
		return
	}
	fmt.Println("asd:", b)
	w.WriteHeader(http.StatusCreated)
	byteBook, err := json.MarshalIndent(b, "", "    ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("errL:", err)
		return
	}
	w.Write(byteBook)
}

// pattern: /library/{book} + json (true or false)
// method: PATCH
// info: JSON
func (h *HTTPHandlers) HandleCompleteBook(w http.ResponseWriter, r *http.Request) {
	completeDto := CompleteDTO{}
	if err := json.NewDecoder(r.Body).Decode(&completeDto); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("errL:", err)
		return
	}
	bookIdstr := mux.Vars(r)["book"]

	bookId, err := strconv.Atoi(bookIdstr)
	if err != nil {
		fmt.Println("4zh:", err)
		return
	}
	book, err := h.library.GetBook(bookId)
	if err != nil {
		fmt.Println("eblan net takoi knigi:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !completeDto.Complete {
		if book.Completed {
			h.library.UncompleteBook(bookId)
			book.Uncomplete()
		} else {
			fmt.Println("eblan итак камплит:", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	} else {
		if !book.Completed {
			h.library.CompleteBook(bookId)
			book.Complete()
		} else {
			fmt.Println("eblan итак анкамплитi:", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	b, err := json.MarshalIndent(book, "", "    ")
	if err != nil {
		fmt.Println("err:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if _, err := w.Write(b); err != nil {
		fmt.Println("err:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

// pattern: /library/{book}
// method: GET
// info: JSON
func (h *HTTPHandlers) HandleGetBook(w http.ResponseWriter, r *http.Request) {
	bookIdstr := mux.Vars(r)["book"]

	bookId, err := strconv.Atoi(bookIdstr)
	if err != nil {
		fmt.Println("4zh:", err)
		return
	}
	book, err := h.library.GetBook(bookId)
	if err != nil {
		fmt.Println("eblan net takoi knigi:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	b, err := json.MarshalIndent(book, "", "    ")
	if err != nil {
		fmt.Println("lol:", err)
		return
	}
	if _, err := w.Write(b); err != nil {
		fmt.Println("pizdez:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

// pattern: /library/
// method: GET
// info: JSON
func (h *HTTPHandlers) HandleGetAllBooks(w http.ResponseWriter, r *http.Request) {
	tmp := h.library.ListBook()
	b, err := json.MarshalIndent(tmp, "", "    ")
	if err != nil {
		fmt.Println("err:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if _, err := w.Write(b); err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println(tmp)

}

// pattern: /library/{book}
// method: DELETE
// info: -
func (h *HTTPHandlers) HandleDeleteBook(w http.ResponseWriter, r *http.Request) {
	bookIdstr := mux.Vars(r)["book"]

	bookId, err := strconv.Atoi(bookIdstr)
	if err != nil {
		fmt.Println("4zh:", err)
		return
	}
	if err := h.library.DeleteBook(bookId); err != nil {
		fmt.Println("dayn ne to ydalil:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// pattern: /library?author=(author)
// method: GET
// info: JSON
func (h *HTTPHandlers) HandleGetAuthorBooks(w http.ResponseWriter, r *http.Request) {
	author := r.URL.Query().Get("author")
	if author == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("eblan")
		return
	}
	tmp := h.library.ListAuthorBooks(author)
	b, err := json.MarshalIndent(tmp, "", "    ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("err:", err)
		return
	}
	if _, err := w.Write(b); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("err:", err)
		return
	}
}

// pattern: /library?complete=(true/false)
// method: GET
// info: JSON
func (h *HTTPHandlers) HandleGetCompletedBooks(w http.ResponseWriter, r *http.Request) {
	isCompStr := r.URL.Query().Get("complete")
	var isComp bool
	switch isCompStr {
	case "true":
		isComp = true
	case "false":
		isComp = false
	default:
		fmt.Println("квери как у дауна")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if isComp {
		tmp := h.library.ListCompletedBooks()
		b, err := json.MarshalIndent(tmp, "", "    ")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println("err:", err)
			return
		}
		if _, err := w.Write(b); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println("err:", err)
			return
		}
	} else {
		tmp := h.library.ListUncompletedBooks()
		b, err := json.MarshalIndent(tmp, "", "    ")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println("err:", err)
			return
		}
		if _, err := w.Write(b); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println("err:", err)
			return
		}
	}

}
