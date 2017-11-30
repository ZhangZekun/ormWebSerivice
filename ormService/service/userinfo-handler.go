package service

import (
    "net/http"
    "strconv"

    "webService/ormService/entities"

    "github.com/unrolled/render"
)

func postUserInfoHandler(formatter *render.Render) http.HandlerFunc {

    return func(w http.ResponseWriter, req *http.Request) {
        req.ParseForm()
        if len(req.Form["username"][0]) == 0 {
            formatter.JSON(w, http.StatusBadRequest, struct{ ErrorIndo string }{"Bad Input!"})
            return
        }

        u := &entities.UserInfo{UserName: req.Form["username"][0]}
        u.DepartName = req.Form["departname"][0]
        _, err2 := entities.Enginea.Insert(u)
        entities.CheckErr(err2)
        formatter.JSON(w, http.StatusOK, u)
    }
}

func getUserInfoHandler(formatter *render.Render) http.HandlerFunc {

    return func(w http.ResponseWriter, req *http.Request) {
        req.ParseForm()
        if len(req.Form["userid"][0]) != 0 {
            i, _ := strconv.ParseInt(req.Form["userid"][0], 10, 32)
            UserInfoGetFromId := &entities.UserInfo{UID:int(i)}
            _, err3 := entities.Enginea.Get(UserInfoGetFromId)
            entities.CheckErr(err3)
            formatter.JSON(w, http.StatusOK, UserInfoGetFromId)
            return
        }
        ulist := make([]entities.UserInfo, 0)
        err := entities.Enginea.Find(&ulist)
        entities.CheckErr(err)
        formatter.JSON(w, http.StatusOK, ulist)
    }
}