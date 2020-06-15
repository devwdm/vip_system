package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"vipSys/dao"
	"vipSys/model"
)

func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 核心处理方式
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
		c.Header("Access-Control-Allow-Methods", "GET, OPTIONS, POST, PUT, DELETE")
		c.Set("content-type", "application/json")
		c.Next()
	}
}

func main() {

	fmt.Println("开始启动了")
	router := gin.Default()
	router.Use(CorsMiddleware())
	//	登录
	router.POST("/login", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password") // 可设置默认值
		admin, err := dao.CheckAdminNameAndPassword(username, password)
		if err == nil {
			if admin.ID > 0 {
				dao.UpdateLastLogin(admin.ID)
				c.JSON(http.StatusOK, gin.H{"code": "0",
					"adminInfo": admin})
			} else {
				c.JSON(http.StatusOK, gin.H{"code": "400"})
			}
		} else {
			c.JSON(http.StatusOK, gin.H{"code": "100"})
		}
	})
	//	新增会员
	router.POST("/addVipInfo", func(c *gin.Context) {
		name := c.PostForm("name")
		iSex, _ := strconv.Atoi(c.PostForm("sex"))
		mobile := c.PostForm("mobile")
		iBelong, _ := strconv.Atoi(c.PostForm("belong"))
		vip := model.VipInfo{
			Name:   name,
			Sex:    iSex,
			Mobile: mobile,
			Belong: iBelong,
		}
		rowCount, err := dao.AddVip(&vip)
		if err == nil && rowCount > 0 {
			c.JSON(http.StatusOK, gin.H{"code": "0", "rowCount": rowCount})
		} else {
			c.JSON(http.StatusOK, gin.H{"code": "100", "rowCount": 0})
		}

	})
	//	会员列表
	router.POST("/vipInfo", func(c *gin.Context) {
		page_no := c.PostForm("page_no")
		if page_no == "" || page_no == "0" {
			page_no = "1"
		}
		page_size := c.PostForm("page_size")
		if page_size == "" || page_size == "0" {
			page_size = "30"
		}
		//将页码页面大小转换为int64类型
		iPageNo, _ := strconv.ParseInt(page_no, 10, 64)
		iPageSize, _ := strconv.ParseInt(page_size, 10, 64)
		//调用vipdao获取分页的会员
		page, err := dao.GetPageVipInfo(iPageNo, iPageSize)
		if err == nil {
			c.JSON(http.StatusOK, gin.H{"code": "0",
				"page_no":       page.PageNo,
				"page_size":     page.PageSize,
				"total_page_no": page.TotalPageNo,
				"total_record":  page.TotalRecord,
				"vip_infos":     page.VipInfos})
		} else {
			c.JSON(http.StatusOK, gin.H{"code": "100"})
		}
	})
	//	搜索会员
	router.POST("/vipInfoBySearch", func(c *gin.Context) {
		page_no := c.PostForm("page_no")
		if page_no == "" || page_no == "0" {
			page_no = "1"
		}
		page_size := c.PostForm("page_size")
		if page_size == "" || page_size == "0" {
			page_size = "30"
		}
		id := c.PostForm("id")
		if id == "" {
			id = "0"
		}
		name := c.PostForm("name")
		mobile := c.PostForm("mobile")
		//将页码页面大小转换为int64类型
		iPageNo, _ := strconv.ParseInt(page_no, 10, 64)
		iPageSize, _ := strconv.ParseInt(page_size, 10, 64)
		iId, _ := strconv.ParseInt(id, 10, 64)
		page, err := dao.GetPageVipInfoByWhere(iId, name, mobile, iPageNo, iPageSize)
		if err == nil {
			c.JSON(http.StatusOK, gin.H{"code": "0",
				"page_no":       page.PageNo,
				"page_size":     page.PageSize,
				"total_page_no": page.TotalPageNo,
				"total_record":  page.TotalRecord,
				"vip_infos":     page.VipInfos})
		} else {
			c.JSON(http.StatusOK, gin.H{"code": "100"})
		}
	})
	//	获取单个会员
	router.POST("/vipInfoOne", func(c *gin.Context) {
		id := c.PostForm("id")
		if id != "" {
			iId, _ := strconv.ParseInt(id, 10, 32)
			vipInfo, err := dao.GetOneVipInfoById(iId)
			if err == nil {
				c.JSON(http.StatusOK, gin.H{"code": "0",
					"vip_info": vipInfo,
				})
			} else {
				c.JSON(http.StatusOK, gin.H{"code": "100"})
			}

		} else {
			c.JSON(http.StatusOK, gin.H{"code": "300"})
		}

	})
	//	修改会员信息
	router.POST("/updataVipInfo", func(c *gin.Context) {
		vipId, _ := strconv.Atoi(c.PostForm("vipId"))
		adminId, _ := strconv.Atoi(c.PostForm("adminId"))
		name := c.PostForm("name")
		iSex, _ := strconv.Atoi(c.PostForm("sex"))
		mobile := c.PostForm("mobile")
		iStatus, _ := strconv.Atoi(c.PostForm("status"))
		vip := model.VipInfo{
			ID:     vipId,
			Name:   name,
			Sex:    iSex,
			Mobile: mobile,
			Status: iStatus,
		}
		rowCount, err := dao.UpdateVip(&vip)
		if err == nil && rowCount > 0 {
			dao.AddLog(adminId, 2, "修改了["+string(vipId)+"]的信息")
			c.JSON(http.StatusOK, gin.H{"code": "0", "rowCount": rowCount})
		} else {
			c.JSON(http.StatusOK, gin.H{"code": "100", "rowCount": 0})
		}
	})
	//	删除会员
	router.POST("/deleteVipInfo", func(c *gin.Context) {
		vipId, _ := strconv.Atoi(c.PostForm("vipId"))
		adminId, _ := strconv.Atoi(c.PostForm("adminId"))
		rowCount, err := dao.DelVipById(vipId)
		if err == nil && rowCount > 0 {
			dao.AddLog(adminId, 3, "用户["+string(adminId)+"]删除会员["+string(vipId)+"]")
			c.JSON(http.StatusOK, gin.H{"code": "0", "rowCount": rowCount})
		} else {
			c.JSON(http.StatusOK, gin.H{"code": "100", "rowCount": 0})
		}

	})
	//	增加积分
	router.POST("/addJiFen", func(c *gin.Context) {
		vipId, _ := strconv.Atoi(c.PostForm("vipId"))
		adminId, _ := strconv.Atoi(c.PostForm("adminId"))
		jiFen, _ := strconv.ParseFloat(c.PostForm("jiFen"), 2)
		rowCount, err := dao.AddJiFen(vipId, jiFen)
		if err == nil && rowCount > 0 {
			//业务日志
			dao.AddJiFenRecord(vipId, adminId, 1, jiFen)
			c.JSON(http.StatusOK, gin.H{"code": "0", "rowCount": rowCount})
		} else {
			c.JSON(http.StatusOK, gin.H{"code": "100", "rowCount": 0})
		}

	})
	//	减少积分
	router.POST("/reduceJiFen", func(c *gin.Context) {
		vipId, _ := strconv.Atoi(c.PostForm("vipId"))
		adminId, _ := strconv.Atoi(c.PostForm("adminId"))
		jiFen, _ := strconv.ParseFloat(c.PostForm("jiFen"), 2)
		rowCount, err := dao.ReduceJiFen(vipId, jiFen)
		if err == nil && rowCount > 0 {
			//业务日志
			dao.AddJiFenRecord(vipId, adminId, 2, jiFen)
			c.JSON(http.StatusOK, gin.H{"code": "0", "rowCount": rowCount})
		} else {
			c.JSON(http.StatusOK, gin.H{"code": "100", "rowCount": 0})
		}

	})
	//	用户列表
	router.POST("/adminInfo", func(c *gin.Context) {
		page_no := c.PostForm("page_no")
		if page_no == "" || page_no == "0" {
			page_no = "1"
		}
		page_size := c.PostForm("page_size")
		fmt.Println(page_size)
		if page_size == "" || page_size == "0" {
			page_size = "30"
		}
		//将页码页面大小转换为int64类型
		iPageNo, _ := strconv.ParseInt(page_no, 10, 64)
		iPageSize, _ := strconv.ParseInt(page_size, 10, 64)
		page, err := dao.GetPageAdminInfo(iPageNo, iPageSize)
		if err == nil {
			c.JSON(http.StatusOK, gin.H{"code": "0",
				"page_no":       page.PageNo,
				"page_size":     page.PageSize,
				"total_page_no": page.TotalPageNo,
				"total_record":  page.TotalRecord,
				"admin_infos":   page.AdminInfos})
		} else {
			c.JSON(http.StatusOK, gin.H{"code": "100"})
		}
	})
	//	新增用户
	router.POST("/addAdminInfo", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		iSex, _ := strconv.Atoi(c.PostForm("sex"))
		mobile := c.PostForm("mobile")
		if password == "" && username != "" && mobile != "" {
			password = "123456"
			adminInfo := model.AdminInfo{
				Username: username,
				Password: password,
				Sex:      iSex,
				Mobile:   mobile,
			}
			rowCount, err := dao.AddAdmin(&adminInfo)
			if err == nil && rowCount > 0 {
				c.JSON(http.StatusOK, gin.H{"code": "0", "rowCount": rowCount})
			} else {
				c.JSON(http.StatusOK, gin.H{"code": "100", "rowCount": 0})
			}
		} else {
			c.JSON(http.StatusOK, gin.H{"code": "300", "rowCount": 0})
		}

	})
	//	删除用户
	router.POST("/deleteAdminInfo", func(c *gin.Context) {
		delId, _ := strconv.Atoi(c.PostForm("delId"))
		adminId, _ := strconv.Atoi(c.PostForm("adminId"))
		rowCount, err := dao.DelAdminById(delId)
		if err == nil && rowCount > 0 {
			dao.AddLog(adminId, 3, "用户["+string(adminId)+"]被删除")
			c.JSON(http.StatusOK, gin.H{"code": "0", "rowCount": rowCount})
		} else {
			c.JSON(http.StatusOK, gin.H{"code": "100", "rowCount": 0})
		}
	})
	//	用户修改密码
	router.POST("/updatepwd", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.PostForm("id"))
		newPassword := c.PostForm("new_password")
		oldPassword := c.PostForm("old_password")
		if id != 0 && newPassword != "" && newPassword != "" {
			if dao.CheckAdminPassword(oldPassword, id) {
				rowCount, err := dao.UpdateAdminPassword(newPassword, id)
				if err == nil && rowCount > 0 {
					c.JSON(http.StatusOK, gin.H{"code": "0", "rowCount": rowCount})
				} else {
					c.JSON(http.StatusOK, gin.H{"code": "100", "rowCount": 0})
				}
			} else {
				c.JSON(http.StatusOK, gin.H{"code": "400", "rowCount": 0})
			}
		} else {
			c.JSON(http.StatusOK, gin.H{"code": "300", "rowCount": 0})
		}
	})
	// 验证登录用户名或手机号是否存在
	router.POST("/checkUsernameAndMobile", func(c *gin.Context) {
		username := c.PostForm("username")
		result := dao.CheckUsernameAndMobile(username)
		c.JSON(http.StatusOK, gin.H{"code": "0", "result": result})
	})
	// 验证用户名是否存在
	router.POST("/checkUsername", func(c *gin.Context) {
		username := c.PostForm("username")
		result := dao.CheckAdminUsername(username)
		c.JSON(http.StatusOK, gin.H{"code": "0", "result": result})
	})
	// 验证用户手机号是否存在
	router.POST("/checkUserMobile", func(c *gin.Context) {
		mobile := c.PostForm("mobile")
		result := dao.CheckAdminMobile(mobile)
		c.JSON(http.StatusOK, gin.H{"code": "0", "result": result})
	})
	// 验证会员名是否存在
	router.POST("/checkVipName", func(c *gin.Context) {
		username := c.PostForm("name")
		result := dao.CheckVipName(username)
		c.JSON(http.StatusOK, gin.H{"code": "0", "result": result})
	})
	// 验证用户手机号是否存在
	router.POST("/checkVipMobile", func(c *gin.Context) {
		mobile := c.PostForm("mobile")
		result := dao.CheckVipMobile(mobile)
		c.JSON(http.StatusOK, gin.H{"code": "0", "result": result})
	})

	////设置处理静态资源，如css和js文件
	//http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("views/static"))))
	////直接去html页面
	//http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("views/pages"))))
	////系统根目录
	//http.HandleFunc("/", controller.Root)
	////系统登录	//登录
	//http.HandleFunc("/login", controller.Login)
	////退出登录
	//http.HandleFunc("/logout", controller.Logout)
	////会员列表
	//http.HandleFunc("/vipinfo", controller.GetPageVipInfo)
	//http.HandleFunc("/vipinfoPage", controller.GetAllVipInfo)
	//http.HandleFunc("/getPageVipInfoBySearch", controller.GetPageVipInfoBySearch)
	////新增会员
	////http.HandleFunc("/addinfo", controller.AddVip)
	////登录成功首页
	//http.HandleFunc("/index", controller.Index)
	////用户信息
	//http.HandleFunc("/adminfo", controller.GetAllAdminInfo)
	////修改密码页面解析
	//http.HandleFunc("/updatepwdhtml", controller.UpdatepwdHtml)
	////修改密码
	//http.HandleFunc("/updatepwd", controller.Updatepwd)
	////通过Ajax请求验证旧密码是否正确
	//http.HandleFunc("/checkAdminOldPwd", controller.CheckAdminOldPwd)
	////通过Ajax请求验证用户名或手机号是否存在
	//http.HandleFunc("/checkUsernameAndMobile", controller.CheckUsernameAndMobile)

	fmt.Println("一切正常开始监听80端口")
	http.ListenAndServe(":80", router)
}
