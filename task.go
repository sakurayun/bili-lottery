package main

import (
	"crypto/rand"
	"fmt"
	"github.com/go-co-op/gocron"
	"log"
	"math/big"
	"strconv"
	"time"
)

func taskStart() {
	getDynamicInfo()

	// 定时任务
	if config.Select.EndTime != 0 && time.Now().Unix() < int64(config.Select.EndTime) {
		s := gocron.NewScheduler(time.UTC)
		_, err := s.Every(config.Select.Time).Seconds().Do(userTask)
		checkErr(err)
		go s.StartBlocking()
		log.Printf("抓取中, 抽奖结束时间: %v",
			time.Unix(0, int64(config.Select.EndTime)*int64(time.Second)).
				Format("2006-01-02 15:04:05"))
		waitCronStop(s)
	} else if config.Select.Member == 0 || time.Now().Unix() >= int64(config.Select.EndTime) {
		userTask()
	}

	if config.Weights.Fan != 0 && config.Weights.Friend != 0 {
		getRelation()

	}

	weightCalculate()
}

// 等待任务结束
func waitCronStop(s *gocron.Scheduler) {
	for {
		task := time.NewTimer(100 * time.Millisecond)
		now := time.Now()
		if now.Unix() >= int64(config.Select.EndTime) {
			s.Remove(s)
		}
		<-task.C
	}
}

func userTask() {
	if config.Weights.Forward != 0 {
		getForwardList()
	}

	if config.Weights.Comment != 0 {
		getCommentList()
	}

	if config.Weights.Like != 0 {
		getLikeList()
	}
}

// 登录验证
func verifyLogin() bool {
	headers := map[string]string{
		"dnt":        "1",
		"user-agent": ua,
		"origin":     "https://www.bilibili.com",
		"refer":      "https://www.bilibili.com/",
	}

	nav := &Nav{}
	_, err := client.R().
		SetResult(nav).
		SetHeaders(headers).
		Get("/web-interface/nav")

	checkErr(err)

	if nav.Code == -101 {
		log.Println("帐号未登录，请检查cookies.")
		return false
	}

	uname := nav.Data.Uname

	log.Printf("登录成功, 当前帐号: %v.", uname)
	return true
}

// 获取动态基本信息
func getDynamicInfo() {
	detail := &Detail{}

	params := map[string]string{
		"id": config.Dynamic.ID,
	}

	_, err := api.R().
		SetResult(detail).
		SetQueryParams(params).
		Get("/polymer/web-dynamic/v1/detail")

	checkErr(err)
	if detail.Code != 0 {
		log.Println(err)
	}

	config.Dynamic.Type = detail.Data.Item.Basic.CommentType
	config.Dynamic.Oid = detail.Data.Item.Basic.CommentIdStr

	err = writeConf(filePath)
	checkErr(err)

	log.Printf("动态类型: %v, 动态OID: %v.", config.Dynamic.Type, config.Dynamic.Oid)
}

// 获取转发列表
func getForwardList() {
	offset := 0
	forward := &ForwardList{}

	for {
		params := map[string]string{
			"dynamic_id": config.Dynamic.ID,
			"offset":     strconv.Itoa(offset),
		}

		_, err := apiLive.R().
			SetResult(forward).
			SetQueryParams(params).
			Get("dynamic_repost/v1/dynamic_repost/view_repost")

		checkErr(err)

		if len(forward.Data.Comments) != 0 {
			for _, comment := range forward.Data.Comments {
				// 确认是否为新用户, 如果是旧用户，那么 result 就不为 0 了，所以要进一步确认
				result1 := searchDB("Mid", comment.Uid)
				if result1 == 0 {
					insertDB("Mid", comment.Uid)
					updateDB("Forward", 1, comment.Uid)
					addWeight(comment.Uid, config.Weights.Forward)
				} else {
					result2 := searchDB("Forward", comment.Uid)
					if result2 == 0 {
						updateDB("Forward", 1, comment.Uid)
						addWeight(comment.Uid, config.Weights.Forward)
					}
				}
				// 确认是否为大会员
				if comment.Detail.Desc.UserProfile.Vip.VipType != 0 {
					checkVIP(comment.Uid)
				}
			}
		} else {
			break
		}

		// 翻页
		if forward.Data.HasMore == false {
			break
		} else {
			if len(forward.Data.Comments) == 0 {
				break
			}
			offset += 20
		}
	}

}

// 获取评论列表
func getCommentList() {
	//var dynamicID string

	next := 0
	comments := &CommentList{}

	//switch config.Dynamic.Type {
	//
	//}

	for {
		params := map[string]string{
			"jsonp": "jsonp",
			"next":  strconv.Itoa(next),
			"type":  strconv.Itoa(config.Dynamic.Type),
			"oid":   config.Dynamic.Oid,
			"mode":  "2",
			"plat":  "1",
			"ps":    "30",
		}

		_, err := api.R().
			SetResult(comments).
			SetQueryParams(params).
			Get("/v2/reply/main")

		checkErr(err)

		if len(comments.Data.Replies) != 0 {
			for _, comment := range comments.Data.Replies {
				// 确认是否为新用户, 如果是旧用户，那么 result 就不为 0 了，所以要进一步确认
				result1 := searchDB("Mid", comment.Mid)
				if result1 == 0 {
					insertDB("Mid", comment.Mid)
					updateDB("Comment", 1, comment.Mid)
					addWeight(comment.Mid, config.Weights.Comment)
				} else {
					result2 := searchDB("Comment", comment.Mid)
					if result2 == 0 {
						updateDB("Comment", 1, comment.Mid)
						addWeight(comment.Mid, config.Weights.Comment)
					}
				}
				// 确认是否为大会员
				if comment.Member.Vip.VipType != 0 {
					checkVIP(comment.Mid)
				}
			}
		} else {
			break
		}

		// 翻页
		if comments.Data.Cursor.Next == 1 {
			break
		} else {
			next = comments.Data.Cursor.Next
		}
	}
}

// 获取点赞列表
func getLikeList() {
	pn := 0
	likes := &LikeList{}

	for {
		params := map[string]string{
			"dynamic_id": config.Dynamic.ID,
			"pn":         strconv.Itoa(pn),
			"ps":         "50",
		}

		_, err := apiVc.R().
			SetResult(likes).
			SetQueryParams(params).
			Get("/dynamic_like/v1/dynamic_like/spec_item_likes")

		checkErr(err)

		if len(likes.Data.ItemLikes) != 0 {
			for _, like := range likes.Data.ItemLikes {
				// 确认是否为新用户, 如果是旧用户，那么 result 就不为 0 了，所以要进一步确认
				result1 := searchDB("Mid", like.Uid)
				if result1 == 0 {
					insertDB("Mid", like.Uid)
					updateDB("Like", 1, like.Uid)
					addWeight(like.Uid, config.Weights.Like)
				} else {
					result2 := searchDB("Like", like.Uid)
					if result2 == 0 {
						updateDB("Like", 1, like.Uid)
						addWeight(like.Uid, config.Weights.Like)
					}
				}
				// 确认是否为大会员
				if like.UserInfo.Vip.VipType != 0 {
					checkVIP(like.Uid)
				}
			}
		} else {
			break
		}

		// 翻页
		if likes.Data.HasMore == 0 {
			break
		} else {
			pn++
		}
	}
}

// 获取用户关系
func getRelation() {
	pn := 1

	headers := map[string]string{
		"dnt":        "1",
		"user-agent": ua,
		"refer":      fmt.Sprintf("https://space.bilibili.com/%v/fans/fans", config.Cookies.DedeUserID),
	}

	for {
		params := map[string]string{
			"vmid":       config.Cookies.DedeUserID,
			"pn":         strconv.Itoa(pn),
			"ps":         "50",
			"order":      "desc",
			"order_type": "attention",
			"jsonp":      "jsonp",
		}

		followers := &FollowerList{}

		_, err := client.R().
			SetResult(followers).
			SetHeaders(headers).
			SetQueryParams(params).
			Get("/relation/followers")

		checkErr(err)

		if len(followers.Data.List) != 0 {
			for _, follower := range followers.Data.List {

				result1 := searchDB("Mid", follower.Mid)
				if result1 == 0 {
					insertDB("Mid", follower.Mid)
					if follower.Attribute == 0 {
						updateDB("Relation", 1, follower.Mid)
						addWeight(follower.Mid, config.Weights.Fan)
					} else if follower.Attribute == 6 {
						updateDB("Relation", 2, follower.Mid)
						addWeight(follower.Mid, config.Weights.Friend)
					}
				} else {
					result2 := searchDB("Relation", follower.Mid)
					if result2 == 0 {
						if follower.Attribute == 0 {
							updateDB("Relation", 1, follower.Mid)
							addWeight(follower.Mid, config.Weights.Fan)
						} else if follower.Attribute == 6 {
							updateDB("Relation", 2, follower.Mid)
							addWeight(follower.Mid, config.Weights.Friend)
						}
					} else {
						// 关系转为互关
						if follower.Attribute == 6 && result2 == 1 {
							updateDB("Relation", 2, follower.Mid)
							addWeight(follower.Mid, config.Weights.Friend-config.Weights.Fan)
							// 取消互关
						} else if follower.Attribute == 0 && result2 == 2 {
							updateDB("Relation", 1, follower.Mid)
							addWeight(follower.Mid, config.Weights.Fan-config.Weights.Friend)
						}
					}
				}

				// 确认是否为大会员
				if follower.Vip.VipType != 0 {
					checkVIP(follower.Mid)
				}
			}
		} else {
			break
		}

		// 翻页
		if pn*50 >= followers.Data.Total {
			break
		} else {
			pn++
		}
	}

}

// 是否大会员
func checkVIP(mid int) {
	result := searchDB("VIP", mid)
	if result == 0 {
		updateDB("VIP", 1, mid)
		addWeight(mid, config.Weights.Vip)
	}
}

// 增加权重
func addWeight(uid, w int) {
	if w == 0 {
		return
	}
	result := searchDB("Weights", uid)
	weight := result + w
	updateDB("Weights", weight, uid)
}

// 计算权重
func weightCalculate() {
	var (
		mid, weights, uid int
		weightList        []int
	)

	searchWeights := fmt.Sprintf("SELECT Mid, Weights FROM %v", tableName)

	rows, err := db.Query(searchWeights)
	checkErr(err)

	for rows.Next() {
		err = rows.Scan(&mid, &weights)
		for i := 0; i < weights; i++ {
			weightList = append(weightList, mid)
		}
	}

Loop:
	n, err := rand.Int(rand.Reader, big.NewInt(int64(len(weightList))))
	if config.Select.Member == 0 {
		uid = weightList[int(n.Int64())]
	} else {
		uid = config.Select.Member
	}
	if !luckyMember(uid) {
		goto Loop
	}

}

// 中奖用户信息
func luckyMember(uid int) bool {
	headers := map[string]string{
		"dnt":        "1",
		"user-agent": ua,
		"origin":     "https://space.bilibili.com",
		"refer":      fmt.Sprintf("https://space.bilibili.com/%v/dynamic", uid),
	}

	params := map[string]string{
		"mid":   strconv.Itoa(uid),
		"jsonp": "jsonp",
	}

	member := &Member{}

	_, err := api.R().
		SetHeaders(headers).
		SetQueryParams(params).
		SetResult(member).
		Get("/space/acc/info")

	checkErr(err)

	if member.Data.Level <= config.Select.Level && config.Select.Level != 0 && config.Select.Member == 0 {
		log.Printf("不符合所选条件, 等级 <= %v, 重抽中...", config.Select.Level)
		time.Sleep(1 * time.Second)
		return false
	}

	lucky := fmt.Sprintf("中奖用户: %v, UID: %v.", member.Data.Name, uid)
	log.Printf(lucky)

	return true
}
