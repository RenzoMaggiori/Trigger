package user

import "net/http"

func (h *Handler) GetUsers(w http.ResponseWriter, req *http.Request)          {}
func (h *Handler) GetUserById(w http.ResponseWriter, req *http.Request)       {}
func (h *Handler) GetUserByEmail(w http.ResponseWriter, req *http.Request)    {}
func (h *Handler) AddUser(w http.ResponseWriter, req *http.Request)           {}
func (h *Handler) UpdateUserById(w http.ResponseWriter, req *http.Request)    {}
func (h *Handler) UpdateUserByEmail(w http.ResponseWriter, req *http.Request) {}
func (h *Handler) DeleteUserById(w http.ResponseWriter, req *http.Request)    {}
func (h *Handler) DeleteUserByEmail(w http.ResponseWriter, req *http.Request) {}
