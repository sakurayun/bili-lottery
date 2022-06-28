package main

// Detail 动态信息
type Detail struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Ttl     int    `json:"ttl"`
	Data    struct {
		Item struct {
			Basic struct {
				CommentIdStr string `json:"comment_id_str"`
				CommentType  int    `json:"comment_type"`
				LikeIcon     struct {
					ActionUrl string `json:"action_url"`
					EndUrl    string `json:"end_url"`
					Id        int    `json:"id"`
					StartUrl  string `json:"start_url"`
				} `json:"like_icon"`
				RidStr string `json:"rid_str"`
			} `json:"basic"`
			IdStr   string `json:"id_str"`
			Modules struct {
				ModuleAuthor struct {
					Decorate struct {
						CardUrl string `json:"card_url"`
						Fan     struct {
							Color  string `json:"color"`
							IsFan  bool   `json:"is_fan"`
							NumStr string `json:"num_str"`
							Number int    `json:"number"`
						} `json:"fan"`
						Id      int    `json:"id"`
						JumpUrl string `json:"jump_url"`
						Name    string `json:"name"`
						Type    int    `json:"type"`
					} `json:"decorate"`
					Face           string `json:"face"`
					FaceNft        bool   `json:"face_nft"`
					Following      bool   `json:"following"`
					JumpUrl        string `json:"jump_url"`
					Label          string `json:"label"`
					Mid            int    `json:"mid"`
					Name           string `json:"name"`
					OfficialVerify struct {
						Desc string `json:"desc"`
						Type int    `json:"type"`
					} `json:"official_verify"`
					Pendant struct {
						Expire            int    `json:"expire"`
						Image             string `json:"image"`
						ImageEnhance      string `json:"image_enhance"`
						ImageEnhanceFrame string `json:"image_enhance_frame"`
						Name              string `json:"name"`
						Pid               int    `json:"pid"`
					} `json:"pendant"`
					PubAction string `json:"pub_action"`
					PubTime   string `json:"pub_time"`
					PubTs     int    `json:"pub_ts"`
					Type      string `json:"type"`
					Vip       struct {
						AvatarSubscript    int    `json:"avatar_subscript"`
						AvatarSubscriptUrl string `json:"avatar_subscript_url"`
						DueDate            int64  `json:"due_date"`
						Label              struct {
							BgColor               string `json:"bg_color"`
							BgStyle               int    `json:"bg_style"`
							BorderColor           string `json:"border_color"`
							ImgLabelUriHans       string `json:"img_label_uri_hans"`
							ImgLabelUriHansStatic string `json:"img_label_uri_hans_static"`
							ImgLabelUriHant       string `json:"img_label_uri_hant"`
							ImgLabelUriHantStatic string `json:"img_label_uri_hant_static"`
							LabelTheme            string `json:"label_theme"`
							Path                  string `json:"path"`
							Text                  string `json:"text"`
							TextColor             string `json:"text_color"`
							UseImgLabel           bool   `json:"use_img_label"`
						} `json:"label"`
						NicknameColor string `json:"nickname_color"`
						Status        int    `json:"status"`
						ThemeType     int    `json:"theme_type"`
						Type          int    `json:"type"`
					} `json:"vip"`
				} `json:"module_author"`
				ModuleDynamic struct {
					Additional interface{} `json:"additional"`
					Desc       struct {
						RichTextNodes []struct {
							OrigText string `json:"orig_text"`
							Rid      string `json:"rid,omitempty"`
							Text     string `json:"text"`
							Type     string `json:"type"`
						} `json:"rich_text_nodes"`
						Text string `json:"text"`
					} `json:"desc"`
					Major struct {
						Draw struct {
							Id    int `json:"id"`
							Items []struct {
								Height int           `json:"height"`
								Size   float64       `json:"size"`
								Src    string        `json:"src"`
								Tags   []interface{} `json:"tags"`
								Width  int           `json:"width"`
							} `json:"items"`
						} `json:"draw"`
						Type string `json:"type"`
					} `json:"major"`
					Topic interface{} `json:"topic"`
				} `json:"module_dynamic"`
				ModuleMore struct {
					ThreePointItems []struct {
						Label string `json:"label"`
						Type  string `json:"type"`
					} `json:"three_point_items"`
				} `json:"module_more"`
				ModuleStat struct {
					Comment struct {
						Count     int  `json:"count"`
						Forbidden bool `json:"forbidden"`
					} `json:"comment"`
					Forward struct {
						Count     int  `json:"count"`
						Forbidden bool `json:"forbidden"`
					} `json:"forward"`
					Like struct {
						Count     int  `json:"count"`
						Forbidden bool `json:"forbidden"`
						Status    bool `json:"status"`
					} `json:"like"`
				} `json:"module_stat"`
			} `json:"modules"`
			Type    string `json:"type"`
			Visible bool   `json:"visible"`
		} `json:"item"`
	} `json:"data"`
}

// CommentList 评论列表
type CommentList struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Ttl     int    `json:"ttl"`
	Data    struct {
		Cursor struct {
			AllCount    int    `json:"all_count"`
			IsBegin     bool   `json:"is_begin"`
			Prev        int    `json:"prev"`
			Next        int    `json:"next"`
			IsEnd       bool   `json:"is_end"`
			Mode        int    `json:"mode"`
			ShowType    int    `json:"show_type"`
			SupportMode []int  `json:"support_mode"`
			Name        string `json:"name"`
		} `json:"cursor"`
		Hots    []interface{} `json:"hots"`
		Notice  interface{}   `json:"notice"`
		Replies []struct {
			Rpid      int64  `json:"rpid"`
			Oid       int    `json:"oid"`
			Type      int    `json:"type"`
			Mid       int    `json:"mid"`
			Root      int    `json:"root"`
			Parent    int    `json:"parent"`
			Dialog    int    `json:"dialog"`
			Count     int    `json:"count"`
			Rcount    int    `json:"rcount"`
			State     int    `json:"state"`
			Fansgrade int    `json:"fansgrade"`
			Attr      int    `json:"attr"`
			Ctime     int    `json:"ctime"`
			RpidStr   string `json:"rpid_str"`
			RootStr   string `json:"root_str"`
			ParentStr string `json:"parent_str"`
			Like      int    `json:"like"`
			Action    int    `json:"action"`
			Member    struct {
				Mid            string `json:"mid"`
				Uname          string `json:"uname"`
				Sex            string `json:"sex"`
				Sign           string `json:"sign"`
				Avatar         string `json:"avatar"`
				Rank           string `json:"rank"`
				DisplayRank    string `json:"DisplayRank"`
				FaceNftNew     int    `json:"face_nft_new"`
				IsSeniorMember int    `json:"is_senior_member"`
				LevelInfo      struct {
					CurrentLevel int `json:"current_level"`
					CurrentMin   int `json:"current_min"`
					CurrentExp   int `json:"current_exp"`
					NextExp      int `json:"next_exp"`
				} `json:"level_info"`
				Pendant struct {
					Pid               int    `json:"pid"`
					Name              string `json:"name"`
					Image             string `json:"image"`
					Expire            int    `json:"expire"`
					ImageEnhance      string `json:"image_enhance"`
					ImageEnhanceFrame string `json:"image_enhance_frame"`
				} `json:"pendant"`
				Nameplate struct {
					Nid        int    `json:"nid"`
					Name       string `json:"name"`
					Image      string `json:"image"`
					ImageSmall string `json:"image_small"`
					Level      string `json:"level"`
					Condition  string `json:"condition"`
				} `json:"nameplate"`
				OfficialVerify struct {
					Type int    `json:"type"`
					Desc string `json:"desc"`
				} `json:"official_verify"`
				Vip struct {
					VipType       int    `json:"vipType"`
					VipDueDate    int64  `json:"vipDueDate"`
					DueRemark     string `json:"dueRemark"`
					AccessStatus  int    `json:"accessStatus"`
					VipStatus     int    `json:"vipStatus"`
					VipStatusWarn string `json:"vipStatusWarn"`
					ThemeType     int    `json:"themeType"`
					Label         struct {
						Path                  string `json:"path"`
						Text                  string `json:"text"`
						LabelTheme            string `json:"label_theme"`
						TextColor             string `json:"text_color"`
						BgStyle               int    `json:"bg_style"`
						BgColor               string `json:"bg_color"`
						BorderColor           string `json:"border_color"`
						UseImgLabel           bool   `json:"use_img_label"`
						ImgLabelUriHans       string `json:"img_label_uri_hans"`
						ImgLabelUriHant       string `json:"img_label_uri_hant"`
						ImgLabelUriHansStatic string `json:"img_label_uri_hans_static"`
						ImgLabelUriHantStatic string `json:"img_label_uri_hant_static"`
					} `json:"label"`
					AvatarSubscript int    `json:"avatar_subscript"`
					NicknameColor   string `json:"nickname_color"`
				} `json:"vip"`
				FansDetail  interface{} `json:"fans_detail"`
				Following   int         `json:"following"`
				IsFollowed  int         `json:"is_followed"`
				UserSailing struct {
					Pendant *struct {
						Id                int    `json:"id"`
						Name              string `json:"name"`
						Image             string `json:"image"`
						JumpUrl           string `json:"jump_url"`
						Type              string `json:"type"`
						ImageEnhance      string `json:"image_enhance"`
						ImageEnhanceFrame string `json:"image_enhance_frame"`
					} `json:"pendant"`
					Cardbg *struct {
						Id      int    `json:"id"`
						Name    string `json:"name"`
						Image   string `json:"image"`
						JumpUrl string `json:"jump_url"`
						Fan     struct {
							IsFan   int    `json:"is_fan"`
							Number  int    `json:"number"`
							Color   string `json:"color"`
							Name    string `json:"name"`
							NumDesc string `json:"num_desc"`
						} `json:"fan"`
						Type string `json:"type"`
					} `json:"cardbg"`
					CardbgWithFocus interface{} `json:"cardbg_with_focus"`
				} `json:"user_sailing"`
				IsContractor   bool        `json:"is_contractor"`
				ContractDesc   string      `json:"contract_desc"`
				NftInteraction interface{} `json:"nft_interaction"`
			} `json:"member"`
			Content struct {
				Message string `json:"message"`
				Plat    int    `json:"plat"`
				Device  string `json:"device"`
				Members []struct {
					Mid            string `json:"mid"`
					Uname          string `json:"uname"`
					Sex            string `json:"sex"`
					Sign           string `json:"sign"`
					Avatar         string `json:"avatar"`
					Rank           string `json:"rank"`
					DisplayRank    string `json:"DisplayRank"`
					FaceNftNew     int    `json:"face_nft_new"`
					IsSeniorMember int    `json:"is_senior_member"`
					LevelInfo      struct {
						CurrentLevel int `json:"current_level"`
						CurrentMin   int `json:"current_min"`
						CurrentExp   int `json:"current_exp"`
						NextExp      int `json:"next_exp"`
					} `json:"level_info"`
					Pendant struct {
						Pid               int    `json:"pid"`
						Name              string `json:"name"`
						Image             string `json:"image"`
						Expire            int    `json:"expire"`
						ImageEnhance      string `json:"image_enhance"`
						ImageEnhanceFrame string `json:"image_enhance_frame"`
					} `json:"pendant"`
					Nameplate struct {
						Nid        int    `json:"nid"`
						Name       string `json:"name"`
						Image      string `json:"image"`
						ImageSmall string `json:"image_small"`
						Level      string `json:"level"`
						Condition  string `json:"condition"`
					} `json:"nameplate"`
					OfficialVerify struct {
						Type int    `json:"type"`
						Desc string `json:"desc"`
					} `json:"official_verify"`
					Vip struct {
						VipType       int    `json:"vipType"`
						VipDueDate    int64  `json:"vipDueDate"`
						DueRemark     string `json:"dueRemark"`
						AccessStatus  int    `json:"accessStatus"`
						VipStatus     int    `json:"vipStatus"`
						VipStatusWarn string `json:"vipStatusWarn"`
						ThemeType     int    `json:"themeType"`
						Label         struct {
							Path                  string `json:"path"`
							Text                  string `json:"text"`
							LabelTheme            string `json:"label_theme"`
							TextColor             string `json:"text_color"`
							BgStyle               int    `json:"bg_style"`
							BgColor               string `json:"bg_color"`
							BorderColor           string `json:"border_color"`
							UseImgLabel           bool   `json:"use_img_label"`
							ImgLabelUriHans       string `json:"img_label_uri_hans"`
							ImgLabelUriHant       string `json:"img_label_uri_hant"`
							ImgLabelUriHansStatic string `json:"img_label_uri_hans_static"`
							ImgLabelUriHantStatic string `json:"img_label_uri_hant_static"`
						} `json:"label"`
						AvatarSubscript int    `json:"avatar_subscript"`
						NicknameColor   string `json:"nickname_color"`
					} `json:"vip"`
				} `json:"members"`
				Emote struct {
					星星眼 struct {
						Id        int    `json:"id"`
						PackageId int    `json:"package_id"`
						State     int    `json:"state"`
						Type      int    `json:"type"`
						Attr      int    `json:"attr"`
						Text      string `json:"text"`
						Url       string `json:"url"`
						Meta      struct {
							Size int `json:"size"`
						} `json:"meta"`
						Mtime     int    `json:"mtime"`
						JumpTitle string `json:"jump_title"`
					} `json:"[星星眼],omitempty"`
					Doge struct {
						Id        int    `json:"id"`
						PackageId int    `json:"package_id"`
						State     int    `json:"state"`
						Type      int    `json:"type"`
						Attr      int    `json:"attr"`
						Text      string `json:"text"`
						Url       string `json:"url"`
						Meta      struct {
							Size int `json:"size"`
						} `json:"meta"`
						Mtime     int    `json:"mtime"`
						JumpTitle string `json:"jump_title"`
					} `json:"[doge],omitempty"`
					藏狐 struct {
						Id        int    `json:"id"`
						PackageId int    `json:"package_id"`
						State     int    `json:"state"`
						Type      int    `json:"type"`
						Attr      int    `json:"attr"`
						Text      string `json:"text"`
						Url       string `json:"url"`
						Meta      struct {
							Size int `json:"size"`
						} `json:"meta"`
						Mtime     int    `json:"mtime"`
						JumpTitle string `json:"jump_title"`
					} `json:"[藏狐],omitempty"`
					脱单doge struct {
						Id        int    `json:"id"`
						PackageId int    `json:"package_id"`
						State     int    `json:"state"`
						Type      int    `json:"type"`
						Attr      int    `json:"attr"`
						Text      string `json:"text"`
						Url       string `json:"url"`
						Meta      struct {
							Size int `json:"size"`
						} `json:"meta"`
						Mtime     int    `json:"mtime"`
						JumpTitle string `json:"jump_title"`
					} `json:"[脱单doge],omitempty"`
					高考加油 struct {
						Id        int    `json:"id"`
						PackageId int    `json:"package_id"`
						State     int    `json:"state"`
						Type      int    `json:"type"`
						Attr      int    `json:"attr"`
						Text      string `json:"text"`
						Url       string `json:"url"`
						Meta      struct {
							Size int `json:"size"`
						} `json:"meta"`
						Mtime     int    `json:"mtime"`
						JumpTitle string `json:"jump_title"`
					} `json:"[高考加油],omitempty"`
				} `json:"emote,omitempty"`
				JumpUrl struct {
				} `json:"jump_url"`
				MaxLine int `json:"max_line"`
			} `json:"content"`
			Replies []interface{} `json:"replies"`
			Assist  int           `json:"assist"`
			Folder  struct {
				HasFolded bool   `json:"has_folded"`
				IsFolded  bool   `json:"is_folded"`
				Rule      string `json:"rule"`
			} `json:"folder"`
			UpAction struct {
				Like  bool `json:"like"`
				Reply bool `json:"reply"`
			} `json:"up_action"`
			ShowFollow   bool `json:"show_follow"`
			Invisible    bool `json:"invisible"`
			ReplyControl struct {
			} `json:"reply_control"`
		} `json:"replies"`
		Top struct {
			Admin interface{} `json:"admin"`
			Upper struct {
				Rpid      int64  `json:"rpid"`
				Oid       int    `json:"oid"`
				Type      int    `json:"type"`
				Mid       int    `json:"mid"`
				Root      int    `json:"root"`
				Parent    int    `json:"parent"`
				Dialog    int    `json:"dialog"`
				Count     int    `json:"count"`
				Rcount    int    `json:"rcount"`
				State     int    `json:"state"`
				Fansgrade int    `json:"fansgrade"`
				Attr      int    `json:"attr"`
				Ctime     int    `json:"ctime"`
				RpidStr   string `json:"rpid_str"`
				RootStr   string `json:"root_str"`
				ParentStr string `json:"parent_str"`
				Like      int    `json:"like"`
				Action    int    `json:"action"`
				Member    struct {
					Mid            string `json:"mid"`
					Uname          string `json:"uname"`
					Sex            string `json:"sex"`
					Sign           string `json:"sign"`
					Avatar         string `json:"avatar"`
					Rank           string `json:"rank"`
					DisplayRank    string `json:"DisplayRank"`
					FaceNftNew     int    `json:"face_nft_new"`
					IsSeniorMember int    `json:"is_senior_member"`
					LevelInfo      struct {
						CurrentLevel int `json:"current_level"`
						CurrentMin   int `json:"current_min"`
						CurrentExp   int `json:"current_exp"`
						NextExp      int `json:"next_exp"`
					} `json:"level_info"`
					Pendant struct {
						Pid               int    `json:"pid"`
						Name              string `json:"name"`
						Image             string `json:"image"`
						Expire            int    `json:"expire"`
						ImageEnhance      string `json:"image_enhance"`
						ImageEnhanceFrame string `json:"image_enhance_frame"`
					} `json:"pendant"`
					Nameplate struct {
						Nid        int    `json:"nid"`
						Name       string `json:"name"`
						Image      string `json:"image"`
						ImageSmall string `json:"image_small"`
						Level      string `json:"level"`
						Condition  string `json:"condition"`
					} `json:"nameplate"`
					OfficialVerify struct {
						Type int    `json:"type"`
						Desc string `json:"desc"`
					} `json:"official_verify"`
					Vip struct {
						VipType       int    `json:"vipType"`
						VipDueDate    int64  `json:"vipDueDate"`
						DueRemark     string `json:"dueRemark"`
						AccessStatus  int    `json:"accessStatus"`
						VipStatus     int    `json:"vipStatus"`
						VipStatusWarn string `json:"vipStatusWarn"`
						ThemeType     int    `json:"themeType"`
						Label         struct {
							Path                  string `json:"path"`
							Text                  string `json:"text"`
							LabelTheme            string `json:"label_theme"`
							TextColor             string `json:"text_color"`
							BgStyle               int    `json:"bg_style"`
							BgColor               string `json:"bg_color"`
							BorderColor           string `json:"border_color"`
							UseImgLabel           bool   `json:"use_img_label"`
							ImgLabelUriHans       string `json:"img_label_uri_hans"`
							ImgLabelUriHant       string `json:"img_label_uri_hant"`
							ImgLabelUriHansStatic string `json:"img_label_uri_hans_static"`
							ImgLabelUriHantStatic string `json:"img_label_uri_hant_static"`
						} `json:"label"`
						AvatarSubscript int    `json:"avatar_subscript"`
						NicknameColor   string `json:"nickname_color"`
					} `json:"vip"`
					FansDetail struct {
						Uid              int    `json:"uid"`
						MedalId          int    `json:"medal_id"`
						MedalName        string `json:"medal_name"`
						Score            int    `json:"score"`
						Level            int    `json:"level"`
						Intimacy         int    `json:"intimacy"`
						MasterStatus     int    `json:"master_status"`
						IsReceive        int    `json:"is_receive"`
						MedalColor       int    `json:"medal_color"`
						MedalColorEnd    int    `json:"medal_color_end"`
						MedalColorBorder int64  `json:"medal_color_border"`
						MedalColorName   int64  `json:"medal_color_name"`
						MedalColorLevel  int64  `json:"medal_color_level"`
						GuardLevel       int    `json:"guard_level"`
						GuardIcon        string `json:"guard_icon"`
						HonorIcon        string `json:"honor_icon"`
					} `json:"fans_detail"`
					Following   int `json:"following"`
					IsFollowed  int `json:"is_followed"`
					UserSailing struct {
						Pendant         interface{} `json:"pendant"`
						Cardbg          interface{} `json:"cardbg"`
						CardbgWithFocus interface{} `json:"cardbg_with_focus"`
					} `json:"user_sailing"`
					IsContractor   bool        `json:"is_contractor"`
					ContractDesc   string      `json:"contract_desc"`
					NftInteraction interface{} `json:"nft_interaction"`
				} `json:"member"`
				Content struct {
					Message string        `json:"message"`
					Plat    int           `json:"plat"`
					Device  string        `json:"device"`
					Members []interface{} `json:"members"`
					JumpUrl struct {
						HttpsWwwBilibiliComVideoBV1A34Y1L7W6 struct {
							Title          string `json:"title"`
							State          int    `json:"state"`
							PrefixIcon     string `json:"prefix_icon"`
							AppUrlSchema   string `json:"app_url_schema"`
							AppName        string `json:"app_name"`
							AppPackageName string `json:"app_package_name"`
							ClickReport    string `json:"click_report"`
							IsHalfScreen   bool   `json:"is_half_screen"`
							ExposureReport string `json:"exposure_report"`
							Underline      bool   `json:"underline"`
							MatchOnce      bool   `json:"match_once"`
							PcUrl          string `json:"pc_url"`
							IconPosition   int    `json:"icon_position"`
						} `json:"https://www.bilibili.com/video/BV1A34y1L7w6/"`
					} `json:"jump_url"`
					MaxLine int `json:"max_line"`
				} `json:"content"`
				Replies []interface{} `json:"replies"`
				Assist  int           `json:"assist"`
				Folder  struct {
					HasFolded bool   `json:"has_folded"`
					IsFolded  bool   `json:"is_folded"`
					Rule      string `json:"rule"`
				} `json:"folder"`
				UpAction struct {
					Like  bool `json:"like"`
					Reply bool `json:"reply"`
				} `json:"up_action"`
				ShowFollow   bool `json:"show_follow"`
				Invisible    bool `json:"invisible"`
				ReplyControl struct {
				} `json:"reply_control"`
			} `json:"upper"`
			Vote interface{} `json:"vote"`
		} `json:"top"`
		TopReplies []struct {
			Rpid      int64  `json:"rpid"`
			Oid       int    `json:"oid"`
			Type      int    `json:"type"`
			Mid       int    `json:"mid"`
			Root      int    `json:"root"`
			Parent    int    `json:"parent"`
			Dialog    int    `json:"dialog"`
			Count     int    `json:"count"`
			Rcount    int    `json:"rcount"`
			State     int    `json:"state"`
			Fansgrade int    `json:"fansgrade"`
			Attr      int    `json:"attr"`
			Ctime     int    `json:"ctime"`
			RpidStr   string `json:"rpid_str"`
			RootStr   string `json:"root_str"`
			ParentStr string `json:"parent_str"`
			Like      int    `json:"like"`
			Action    int    `json:"action"`
			Member    struct {
				Mid            string `json:"mid"`
				Uname          string `json:"uname"`
				Sex            string `json:"sex"`
				Sign           string `json:"sign"`
				Avatar         string `json:"avatar"`
				Rank           string `json:"rank"`
				DisplayRank    string `json:"DisplayRank"`
				FaceNftNew     int    `json:"face_nft_new"`
				IsSeniorMember int    `json:"is_senior_member"`
				LevelInfo      struct {
					CurrentLevel int `json:"current_level"`
					CurrentMin   int `json:"current_min"`
					CurrentExp   int `json:"current_exp"`
					NextExp      int `json:"next_exp"`
				} `json:"level_info"`
				Pendant struct {
					Pid               int    `json:"pid"`
					Name              string `json:"name"`
					Image             string `json:"image"`
					Expire            int    `json:"expire"`
					ImageEnhance      string `json:"image_enhance"`
					ImageEnhanceFrame string `json:"image_enhance_frame"`
				} `json:"pendant"`
				Nameplate struct {
					Nid        int    `json:"nid"`
					Name       string `json:"name"`
					Image      string `json:"image"`
					ImageSmall string `json:"image_small"`
					Level      string `json:"level"`
					Condition  string `json:"condition"`
				} `json:"nameplate"`
				OfficialVerify struct {
					Type int    `json:"type"`
					Desc string `json:"desc"`
				} `json:"official_verify"`
				Vip struct {
					VipType       int    `json:"vipType"`
					VipDueDate    int64  `json:"vipDueDate"`
					DueRemark     string `json:"dueRemark"`
					AccessStatus  int    `json:"accessStatus"`
					VipStatus     int    `json:"vipStatus"`
					VipStatusWarn string `json:"vipStatusWarn"`
					ThemeType     int    `json:"themeType"`
					Label         struct {
						Path                  string `json:"path"`
						Text                  string `json:"text"`
						LabelTheme            string `json:"label_theme"`
						TextColor             string `json:"text_color"`
						BgStyle               int    `json:"bg_style"`
						BgColor               string `json:"bg_color"`
						BorderColor           string `json:"border_color"`
						UseImgLabel           bool   `json:"use_img_label"`
						ImgLabelUriHans       string `json:"img_label_uri_hans"`
						ImgLabelUriHant       string `json:"img_label_uri_hant"`
						ImgLabelUriHansStatic string `json:"img_label_uri_hans_static"`
						ImgLabelUriHantStatic string `json:"img_label_uri_hant_static"`
					} `json:"label"`
					AvatarSubscript int    `json:"avatar_subscript"`
					NicknameColor   string `json:"nickname_color"`
				} `json:"vip"`
				FansDetail struct {
					Uid              int    `json:"uid"`
					MedalId          int    `json:"medal_id"`
					MedalName        string `json:"medal_name"`
					Score            int    `json:"score"`
					Level            int    `json:"level"`
					Intimacy         int    `json:"intimacy"`
					MasterStatus     int    `json:"master_status"`
					IsReceive        int    `json:"is_receive"`
					MedalColor       int    `json:"medal_color"`
					MedalColorEnd    int    `json:"medal_color_end"`
					MedalColorBorder int64  `json:"medal_color_border"`
					MedalColorName   int64  `json:"medal_color_name"`
					MedalColorLevel  int64  `json:"medal_color_level"`
					GuardLevel       int    `json:"guard_level"`
					GuardIcon        string `json:"guard_icon"`
					HonorIcon        string `json:"honor_icon"`
				} `json:"fans_detail"`
				Following   int `json:"following"`
				IsFollowed  int `json:"is_followed"`
				UserSailing struct {
					Pendant         interface{} `json:"pendant"`
					Cardbg          interface{} `json:"cardbg"`
					CardbgWithFocus interface{} `json:"cardbg_with_focus"`
				} `json:"user_sailing"`
				IsContractor   bool        `json:"is_contractor"`
				ContractDesc   string      `json:"contract_desc"`
				NftInteraction interface{} `json:"nft_interaction"`
			} `json:"member"`
			Content struct {
				Message string        `json:"message"`
				Plat    int           `json:"plat"`
				Device  string        `json:"device"`
				Members []interface{} `json:"members"`
				JumpUrl struct {
					HttpsWwwBilibiliComVideoBV1A34Y1L7W6 struct {
						Title          string `json:"title"`
						State          int    `json:"state"`
						PrefixIcon     string `json:"prefix_icon"`
						AppUrlSchema   string `json:"app_url_schema"`
						AppName        string `json:"app_name"`
						AppPackageName string `json:"app_package_name"`
						ClickReport    string `json:"click_report"`
						IsHalfScreen   bool   `json:"is_half_screen"`
						ExposureReport string `json:"exposure_report"`
						Underline      bool   `json:"underline"`
						MatchOnce      bool   `json:"match_once"`
						PcUrl          string `json:"pc_url"`
						IconPosition   int    `json:"icon_position"`
					} `json:"https://www.bilibili.com/video/BV1A34y1L7w6/"`
				} `json:"jump_url"`
				MaxLine int `json:"max_line"`
			} `json:"content"`
			Replies []interface{} `json:"replies"`
			Assist  int           `json:"assist"`
			Folder  struct {
				HasFolded bool   `json:"has_folded"`
				IsFolded  bool   `json:"is_folded"`
				Rule      string `json:"rule"`
			} `json:"folder"`
			UpAction struct {
				Like  bool `json:"like"`
				Reply bool `json:"reply"`
			} `json:"up_action"`
			ShowFollow   bool `json:"show_follow"`
			Invisible    bool `json:"invisible"`
			ReplyControl struct {
			} `json:"reply_control"`
		} `json:"top_replies"`
		Folder struct {
			HasFolded bool   `json:"has_folded"`
			IsFolded  bool   `json:"is_folded"`
			Rule      string `json:"rule"`
		} `json:"folder"`
		UpSelection struct {
			PendingCount int `json:"pending_count"`
			IgnoreCount  int `json:"ignore_count"`
		} `json:"up_selection"`
		Cm struct {
		} `json:"cm"`
		CmInfo struct {
			Ads struct {
				Field1 []struct {
					Id           int    `json:"id"`
					ContractId   string `json:"contract_id"`
					PosNum       int    `json:"pos_num"`
					Name         string `json:"name"`
					Pic          string `json:"pic"`
					Litpic       string `json:"litpic"`
					Url          string `json:"url"`
					Style        int    `json:"style"`
					Agency       string `json:"agency"`
					Label        string `json:"label"`
					Intro        string `json:"intro"`
					CreativeType int    `json:"creative_type"`
					RequestId    string `json:"request_id"`
					SrcId        int    `json:"src_id"`
					Area         int    `json:"area"`
					IsAdLoc      bool   `json:"is_ad_loc"`
					AdCb         string `json:"ad_cb"`
					Title        string `json:"title"`
					ServerType   int    `json:"server_type"`
					CmMark       int    `json:"cm_mark"`
					Stime        int    `json:"stime"`
					Mid          string `json:"mid"`
					ActivityType int    `json:"activity_type"`
					Epid         int    `json:"epid"`
					SubTitle     string `json:"sub_title"`
					AdDesc       string `json:"ad_desc"`
					AdverName    string `json:"adver_name"`
					NullFrame    bool   `json:"null_frame"`
					PicMainColor string `json:"pic_main_color"`
				} `json:"4763"`
			} `json:"ads"`
		} `json:"cm_info"`
		Effects struct {
			Preloading string `json:"preloading"`
		} `json:"effects"`
		Assist    int `json:"assist"`
		Blacklist int `json:"blacklist"`
		Vote      int `json:"vote"`
		Config    struct {
			Showtopic  int  `json:"showtopic"`
			ShowUpFlag bool `json:"show_up_flag"`
			ReadOnly   bool `json:"read_only"`
		} `json:"config"`
		Upper struct {
			Mid int `json:"mid"`
		} `json:"upper"`
		Control struct {
			InputDisable          bool   `json:"input_disable"`
			RootInputText         string `json:"root_input_text"`
			ChildInputText        string `json:"child_input_text"`
			GiveupInputText       string `json:"giveup_input_text"`
			BgText                string `json:"bg_text"`
			WebSelection          bool   `json:"web_selection"`
			AnswerGuideText       string `json:"answer_guide_text"`
			AnswerGuideIconUrl    string `json:"answer_guide_icon_url"`
			AnswerGuideIosUrl     string `json:"answer_guide_ios_url"`
			AnswerGuideAndroidUrl string `json:"answer_guide_android_url"`
			ShowType              int    `json:"show_type"`
			ShowText              string `json:"show_text"`
			DisableJumpEmote      bool   `json:"disable_jump_emote"`
		} `json:"control"`
		Note      int         `json:"note"`
		Callbacks interface{} `json:"callbacks"`
	} `json:"data"`
}

// LikeList 点赞列表
type LikeList struct {
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	Message string `json:"message"`
	Data    struct {
		ItemLikes []struct {
			Uid      int    `json:"uid"`
			Time     int    `json:"time"`
			FaceUrl  string `json:"face_url"`
			Uname    string `json:"uname"`
			UserInfo struct {
				Uid            int    `json:"uid"`
				Uname          string `json:"uname"`
				Face           string `json:"face"`
				Rank           string `json:"rank"`
				OfficialVerify struct {
					Type int    `json:"type"`
					Desc string `json:"desc"`
				} `json:"official_verify"`
				Vip struct {
					VipType    int   `json:"vipType"`
					VipDueDate int64 `json:"vipDueDate"`
					VipStatus  int   `json:"vipStatus"`
					ThemeType  int   `json:"themeType"`
					Label      struct {
						Path        string `json:"path"`
						Text        string `json:"text"`
						LabelTheme  string `json:"label_theme"`
						TextColor   string `json:"text_color"`
						BgStyle     int    `json:"bg_style"`
						BgColor     string `json:"bg_color"`
						BorderColor string `json:"border_color"`
					} `json:"label"`
					AvatarSubscript    int    `json:"avatar_subscript"`
					NicknameColor      string `json:"nickname_color"`
					Role               int    `json:"role"`
					AvatarSubscriptUrl string `json:"avatar_subscript_url"`
				} `json:"vip"`
				Pendant struct {
					Pid               int    `json:"pid"`
					Name              string `json:"name"`
					Image             string `json:"image"`
					Expire            int    `json:"expire"`
					ImageEnhance      string `json:"image_enhance"`
					ImageEnhanceFrame string `json:"image_enhance_frame"`
				} `json:"pendant"`
				Sign      string `json:"sign"`
				LevelInfo struct {
					CurrentLevel int `json:"current_level"`
				} `json:"level_info"`
			} `json:"user_info"`
			Attend int `json:"attend"`
		} `json:"item_likes"`
		HasMore    int `json:"has_more"`
		TotalCount int `json:"total_count"`
		Gt         int `json:"_gt_"`
	} `json:"data"`
}

// ForwardList 转发列表
type ForwardList struct {
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	Message string `json:"message"`
	Data    struct {
		HasMore  bool `json:"has_more"`
		Comments []struct {
			Uid     int    `json:"uid"`
			Comment string `json:"comment"`
			Ts      int    `json:"ts"`
			FaceUrl string `json:"face_url"`
			Uname   string `json:"uname"`
			RpDynId int64  `json:"rp_dyn_id"`
			Ctrl    string `json:"ctrl"`
			Detail  struct {
				Desc struct {
					Uid         int   `json:"uid"`
					Type        int   `json:"type"`
					Rid         int64 `json:"rid"`
					Acl         int   `json:"acl"`
					View        int   `json:"view"`
					Repost      int   `json:"repost"`
					Comment     int   `json:"comment"`
					Like        int   `json:"like"`
					IsLiked     int   `json:"is_liked"`
					DynamicId   int64 `json:"dynamic_id"`
					Timestamp   int   `json:"timestamp"`
					PreDyId     int64 `json:"pre_dy_id"`
					OrigDyId    int64 `json:"orig_dy_id"`
					OrigType    int   `json:"orig_type"`
					UserProfile struct {
						Info struct {
							Uid     int    `json:"uid"`
							Uname   string `json:"uname"`
							Face    string `json:"face"`
							FaceNft int    `json:"face_nft"`
						} `json:"info"`
						Card struct {
							OfficialVerify struct {
								Type int    `json:"type"`
								Desc string `json:"desc"`
							} `json:"official_verify"`
						} `json:"card"`
						Vip struct {
							VipType    int   `json:"vipType"`
							VipDueDate int64 `json:"vipDueDate"`
							VipStatus  int   `json:"vipStatus"`
							ThemeType  int   `json:"themeType"`
							Label      struct {
								Path        string `json:"path"`
								Text        string `json:"text"`
								LabelTheme  string `json:"label_theme"`
								TextColor   string `json:"text_color"`
								BgStyle     int    `json:"bg_style"`
								BgColor     string `json:"bg_color"`
								BorderColor string `json:"border_color"`
							} `json:"label"`
							AvatarSubscript    int    `json:"avatar_subscript"`
							NicknameColor      string `json:"nickname_color"`
							Role               int    `json:"role"`
							AvatarSubscriptUrl string `json:"avatar_subscript_url"`
						} `json:"vip"`
						Pendant struct {
							Pid               int    `json:"pid"`
							Name              string `json:"name"`
							Image             string `json:"image"`
							Expire            int    `json:"expire"`
							ImageEnhance      string `json:"image_enhance"`
							ImageEnhanceFrame string `json:"image_enhance_frame"`
						} `json:"pendant"`
						Rank      string `json:"rank"`
						Sign      string `json:"sign"`
						LevelInfo struct {
							CurrentLevel int `json:"current_level"`
						} `json:"level_info"`
					} `json:"user_profile"`
					UidType      int    `json:"uid_type"`
					Stype        int    `json:"stype"`
					RType        int    `json:"r_type"`
					InnerId      int    `json:"inner_id"`
					Status       int    `json:"status"`
					DynamicIdStr string `json:"dynamic_id_str"`
					PreDyIdStr   string `json:"pre_dy_id_str"`
					OrigDyIdStr  string `json:"orig_dy_id_str"`
					RidStr       string `json:"rid_str"`
					Origin       struct {
						Uid         int   `json:"uid"`
						Type        int   `json:"type"`
						Rid         int   `json:"rid"`
						Acl         int   `json:"acl"`
						View        int   `json:"view"`
						Repost      int   `json:"repost"`
						Comment     int   `json:"comment"`
						Like        int   `json:"like"`
						IsLiked     int   `json:"is_liked"`
						DynamicId   int64 `json:"dynamic_id"`
						Timestamp   int   `json:"timestamp"`
						PreDyId     int   `json:"pre_dy_id"`
						OrigDyId    int   `json:"orig_dy_id"`
						UserProfile struct {
							Info struct {
								Uid     int    `json:"uid"`
								Uname   string `json:"uname"`
								Face    string `json:"face"`
								FaceNft int    `json:"face_nft"`
							} `json:"info"`
							Card struct {
								OfficialVerify struct {
									Type int    `json:"type"`
									Desc string `json:"desc"`
								} `json:"official_verify"`
							} `json:"card"`
							Vip struct {
								VipType    int   `json:"vipType"`
								VipDueDate int64 `json:"vipDueDate"`
								VipStatus  int   `json:"vipStatus"`
								ThemeType  int   `json:"themeType"`
								Label      struct {
									Path        string `json:"path"`
									Text        string `json:"text"`
									LabelTheme  string `json:"label_theme"`
									TextColor   string `json:"text_color"`
									BgStyle     int    `json:"bg_style"`
									BgColor     string `json:"bg_color"`
									BorderColor string `json:"border_color"`
								} `json:"label"`
								AvatarSubscript    int    `json:"avatar_subscript"`
								NicknameColor      string `json:"nickname_color"`
								Role               int    `json:"role"`
								AvatarSubscriptUrl string `json:"avatar_subscript_url"`
							} `json:"vip"`
							Pendant struct {
								Pid               int    `json:"pid"`
								Name              string `json:"name"`
								Image             string `json:"image"`
								Expire            int    `json:"expire"`
								ImageEnhance      string `json:"image_enhance"`
								ImageEnhanceFrame string `json:"image_enhance_frame"`
							} `json:"pendant"`
							Rank      string `json:"rank"`
							Sign      string `json:"sign"`
							LevelInfo struct {
								CurrentLevel int `json:"current_level"`
							} `json:"level_info"`
						} `json:"user_profile"`
						UidType      int    `json:"uid_type"`
						Stype        int    `json:"stype"`
						RType        int    `json:"r_type"`
						InnerId      int    `json:"inner_id"`
						Status       int    `json:"status"`
						DynamicIdStr string `json:"dynamic_id_str"`
						PreDyIdStr   string `json:"pre_dy_id_str"`
						OrigDyIdStr  string `json:"orig_dy_id_str"`
						RidStr       string `json:"rid_str"`
					} `json:"origin"`
				} `json:"desc"`
				Card       string `json:"card"`
				ExtendJson string `json:"extend_json"`
				Display    struct {
					Origin struct {
						Relation struct {
							Status     int `json:"status"`
							IsFollow   int `json:"is_follow"`
							IsFollowed int `json:"is_followed"`
						} `json:"relation"`
					} `json:"origin"`
					Relation struct {
						Status     int `json:"status"`
						IsFollow   int `json:"is_follow"`
						IsFollowed int `json:"is_followed"`
					} `json:"relation"`
				} `json:"display"`
			} `json:"detail"`
		} `json:"comments"`
		TotalCount int `json:"total_count"`
		Gt         int `json:"_gt_"`
	} `json:"data"`
}

// FollowerList 粉丝列表
type FollowerList struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Ttl     int    `json:"ttl"`
	Data    struct {
		List []struct {
			Mid          int         `json:"mid"`
			Attribute    int         `json:"attribute"`
			Mtime        int         `json:"mtime"`
			Tag          interface{} `json:"tag"`
			Special      int         `json:"special"`
			ContractInfo struct {
				IsContractor bool `json:"is_contractor"`
				Ts           int  `json:"ts"`
				IsContract   bool `json:"is_contract"`
				UserAttr     int  `json:"user_attr"`
			} `json:"contract_info"`
			Uname          string `json:"uname"`
			Face           string `json:"face"`
			Sign           string `json:"sign"`
			FaceNft        int    `json:"face_nft"`
			OfficialVerify struct {
				Type int    `json:"type"`
				Desc string `json:"desc"`
			} `json:"official_verify"`
			Vip struct {
				VipType       int    `json:"vipType"`
				VipDueDate    int64  `json:"vipDueDate"`
				DueRemark     string `json:"dueRemark"`
				AccessStatus  int    `json:"accessStatus"`
				VipStatus     int    `json:"vipStatus"`
				VipStatusWarn string `json:"vipStatusWarn"`
				ThemeType     int    `json:"themeType"`
				Label         struct {
					Path        string `json:"path"`
					Text        string `json:"text"`
					LabelTheme  string `json:"label_theme"`
					TextColor   string `json:"text_color"`
					BgStyle     int    `json:"bg_style"`
					BgColor     string `json:"bg_color"`
					BorderColor string `json:"border_color"`
				} `json:"label"`
				AvatarSubscript    int    `json:"avatar_subscript"`
				NicknameColor      string `json:"nickname_color"`
				AvatarSubscriptUrl string `json:"avatar_subscript_url"`
			} `json:"vip"`
		} `json:"list"`
		ReVersion int `json:"re_version"`
		Total     int `json:"total"`
	} `json:"data"`
}

// Config 配置文件
type Config struct {
	Version string `toml:"Version"`
	Weights struct {
		Forward int `toml:"Forward"`
		Comment int `toml:"Comment"`
		Like    int `toml:"Like"`
		Fan     int `toml:"Fan"`
		Friend  int `toml:"Friend"`
		Vip     int `toml:"VIP"`
	} `toml:"Weights"`
	Select struct {
		Level   int `toml:"Level"`
		Member  int `toml:"Member"`
		Time    int `toml:"Time"`
		EndTime int `toml:"EndTime"`
	} `toml:"Select"`
	Dynamic struct {
		ID   string `toml:"Id"`
		Oid  string `toml:"Oid"`
		Type int    `toml:"Type"`
	} `toml:"Dynamic"`
	Cookies struct {
		Sessdata        string `toml:"SESSDATA"`
		BiliJct         string `toml:"bili_jct"`
		DedeUserID      string `toml:"DedeUserID"`
		DedeUserIDCkMd5 string `toml:"DedeUserID__ckMd5"`
	} `toml:"Cookies"`
	ProxyURL struct {
		API     string `toml:"api"`
		APIVc   string `toml:"api_vc"`
		APILive string `toml:"api_live"`
	} `toml:"ProxyURL"`
}

// GetLoginUrl 获取登陆链接
type GetLoginUrl struct {
	Code   int  `json:"code"`
	Status bool `json:"status"`
	Ts     int  `json:"ts"`
	Data   struct {
		Url      string `json:"url"`
		OauthKey string `json:"oauthKey"`
	} `json:"data"`
}

// GetLoginInfo 二维码状态
type GetLoginInfo struct {
	Status bool `json:"status"`
	//Data    string `json:"data"`
	Message string `json:"message"`
}

// Nav 登录验证
type Nav struct {
	Code int `json:"code"`
	Data struct {
		Wallet struct {
			BcoinBalance float64 `json:"bcoin_balance"`
		} `json:"wallet"`
		Uname string `json:"uname"`
	} `json:"data"`
}

// Member 用户信息
type Member struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Ttl     int    `json:"ttl"`
	Data    struct {
		Mid         int    `json:"mid"`
		Name        string `json:"name"`
		Sex         string `json:"sex"`
		Face        string `json:"face"`
		FaceNft     int    `json:"face_nft"`
		FaceNftType int    `json:"face_nft_type"`
		Sign        string `json:"sign"`
		Rank        int    `json:"rank"`
		Level       int    `json:"level"`
		Jointime    int    `json:"jointime"`
		Moral       int    `json:"moral"`
		Silence     int    `json:"silence"`
		Coins       int    `json:"coins"`
		FansBadge   bool   `json:"fans_badge"`
		FansMedal   struct {
			Show  bool        `json:"show"`
			Wear  bool        `json:"wear"`
			Medal interface{} `json:"medal"`
		} `json:"fans_medal"`
		Official struct {
			Role  int    `json:"role"`
			Title string `json:"title"`
			Desc  string `json:"desc"`
			Type  int    `json:"type"`
		} `json:"official"`
		Vip struct {
			Type       int   `json:"type"`
			Status     int   `json:"status"`
			DueDate    int64 `json:"due_date"`
			VipPayType int   `json:"vip_pay_type"`
			ThemeType  int   `json:"theme_type"`
			Label      struct {
				Path        string `json:"path"`
				Text        string `json:"text"`
				LabelTheme  string `json:"label_theme"`
				TextColor   string `json:"text_color"`
				BgStyle     int    `json:"bg_style"`
				BgColor     string `json:"bg_color"`
				BorderColor string `json:"border_color"`
			} `json:"label"`
			AvatarSubscript    int    `json:"avatar_subscript"`
			NicknameColor      string `json:"nickname_color"`
			Role               int    `json:"role"`
			AvatarSubscriptUrl string `json:"avatar_subscript_url"`
		} `json:"vip"`
		Pendant struct {
			Pid               int    `json:"pid"`
			Name              string `json:"name"`
			Image             string `json:"image"`
			Expire            int    `json:"expire"`
			ImageEnhance      string `json:"image_enhance"`
			ImageEnhanceFrame string `json:"image_enhance_frame"`
		} `json:"pendant"`
		Nameplate struct {
			Nid        int    `json:"nid"`
			Name       string `json:"name"`
			Image      string `json:"image"`
			ImageSmall string `json:"image_small"`
			Level      string `json:"level"`
			Condition  string `json:"condition"`
		} `json:"nameplate"`
		UserHonourInfo struct {
			Mid    int           `json:"mid"`
			Colour interface{}   `json:"colour"`
			Tags   []interface{} `json:"tags"`
		} `json:"user_honour_info"`
		IsFollowed bool   `json:"is_followed"`
		TopPhoto   string `json:"top_photo"`
		Theme      struct {
		} `json:"theme"`
		SysNotice struct {
		} `json:"sys_notice"`
		LiveRoom struct {
			RoomStatus    int    `json:"roomStatus"`
			LiveStatus    int    `json:"liveStatus"`
			Url           string `json:"url"`
			Title         string `json:"title"`
			Cover         string `json:"cover"`
			Roomid        int    `json:"roomid"`
			RoundStatus   int    `json:"roundStatus"`
			BroadcastType int    `json:"broadcast_type"`
			WatchedShow   struct {
				Switch       bool   `json:"switch"`
				Num          int    `json:"num"`
				TextSmall    string `json:"text_small"`
				TextLarge    string `json:"text_large"`
				Icon         string `json:"icon"`
				IconLocation string `json:"icon_location"`
				IconWeb      string `json:"icon_web"`
			} `json:"watched_show"`
		} `json:"live_room"`
		Birthday   string      `json:"birthday"`
		School     interface{} `json:"school"`
		Profession struct {
			Name       string `json:"name"`
			Department string `json:"department"`
			Title      string `json:"title"`
			IsShow     int    `json:"is_show"`
		} `json:"profession"`
		Tags   interface{} `json:"tags"`
		Series struct {
			UserUpgradeStatus int  `json:"user_upgrade_status"`
			ShowUpgradeWindow bool `json:"show_upgrade_window"`
		} `json:"series"`
		IsSeniorMember int `json:"is_senior_member"`
	} `json:"data"`
}
