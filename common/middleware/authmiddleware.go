package middleware

//
//import (
//	"context"
//	"encoding/json"
//	"fmt"
//	"github.com/zeromicro/go-zero/core/logx"
//	"github.com/zeromicro/go-zero/rest/httpx"
//	"net/http"
//)
//
//type AuthInterceptorMiddleware struct {
//}
//
//func NewAuthInterceptorMiddleware() *AuthInterceptorMiddleware {
//	return &AuthInterceptorMiddleware{}
//}
//
//func (m *AuthInterceptorMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		//跳过登录等不需要验证的接口 ==
//		if r.RequestURI == "/v1/demo" {
//			fmt.Println("跳过接口 ", r.RequestURI)
//			next(w, r)
//			return
//		}
//
//		userID := r.Context().Value("ID").(string)
//		organize := r.Context().Value("Organize").(string)
//		roleID := r.Context().Value("RoleID").(string)
//		departmentsInterface := r.Context().Value("Departments").([]interface{})
//		departmentsTypeInterface := r.Context().Value("DepartmentsType").([]interface{})
//		level2DeptPath := r.Context().Value("Level2DeptPath").(string)
//		levelJsonNum := r.Context().Value("Level").(json.Number)
//
//		var departments []string
//		for _, v := range departmentsInterface {
//			if val, ok := v.(string); ok {
//				departments = append(departments, val)
//			}
//		}
//
//		var departmentsType []uint
//		for _, v := range departmentsTypeInterface {
//			if val, ok := v.(uint); ok {
//				departmentsType = append(departmentsType, val)
//			}
//		}
//		level, err := levelJsonNum.Int64()
//		if err != nil {
//			panic("JWT json.Number 类型转换错误")
//		}
//		user := &schema.User{
//			ID:              userID,
//			Organize:        organize,
//			RoleID:          roleID,
//			Departments:     departments,
//			DepartmentsType: departmentsType,
//			Level2DeptPath:  level2DeptPath,
//			Level:           level,
//		}
//
//		if userID == "" {
//			logx.WithContext(r.Context()).Errorf("无效的Authorization")
//			httpx.WriteJson(w, http.StatusUnauthorized, result.Error(xerr.ErrInvalidToken, "令牌无效", r))
//			return
//		}
//		//具体传递 改成userID
//		ctx := context.WithValue(r.Context(), "userInfo", user)
//		r = r.WithContext(ctx)
//		next(w, r)
//	}
//}
