package quark

type infoResp struct {
	Success bool `json:"success"`
	Data    struct {
		Nickname  string `json:"nickname"`
		AvatarUri string `json:"avatarUri"`
		Mobilekps string `json:"mobilekps"`
		Config    struct {
		} `json:"config"`
	} `json:"data"`
	Code string `json:"code"`
}

type sharePageResp struct {
	Status    int    `json:"status"`
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Timestamp int    `json:"timestamp"`
	Data      struct {
		Subscribed bool   `json:"subscribed"`
		Stoken     string `json:"stoken"`
		ShareType  int    `json:"share_type"`
		Author     struct {
			MemberType string `json:"member_type"`
			AvatarUrl  string `json:"avatar_url"`
			NickName   string `json:"nick_name"`
		} `json:"author"`
		ExpiredType int    `json:"expired_type"`
		ExpiredAt   int64  `json:"expired_at"`
		Title       string `json:"title"`
	} `json:"data"`
	Metadata struct {
		TGroup string `json:"_t_group"`
		GGroup string `json:"_g_group"`
	} `json:"metadata"`
}
type ShareDetailResp struct {
	Status    int    `json:"status"`
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Timestamp int    `json:"timestamp"`
	Data      struct {
		IsOwner int `json:"is_owner"`
		List    []struct {
			Fid                 string  `json:"fid"`
			FileName            string  `json:"file_name"`
			PdirFid             string  `json:"pdir_fid"`
			Category            int     `json:"category"`
			FileType            int     `json:"file_type"`
			Size                int     `json:"size"`
			FormatType          string  `json:"format_type"`
			Status              int     `json:"status"`
			Tags                string  `json:"tags"`
			OwnerUcid           string  `json:"owner_ucid"`
			LCreatedAt          int64   `json:"l_created_at"`
			LUpdatedAt          int64   `json:"l_updated_at"`
			Extra               string  `json:"extra"`
			Source              string  `json:"source"`
			FileSource          string  `json:"file_source"`
			NameSpace           int     `json:"name_space"`
			LShotAt             int64   `json:"l_shot_at"`
			SeriesId            string  `json:"series_id"`
			SourceDisplay       string  `json:"source_display"`
			IncludeItems        int     `json:"include_items,omitempty"`
			SeriesDir           bool    `json:"series_dir"`
			UploadCameraRootDir bool    `json:"upload_camera_root_dir"`
			Fps                 float32 `json:"fps"`
			Like                int     `json:"like"`
			OperatedAt          int64   `json:"operated_at"`
			RiskType            int     `json:"risk_type"`
			BackupSign          int     `json:"backup_sign"`
			FileNameHlStart     int     `json:"file_name_hl_start"`
			FileNameHlEnd       int     `json:"file_name_hl_end"`
			FileStruct          struct {
				FirSource      string `json:"fir_source"`
				SecSource      string `json:"sec_source"`
				ThiSource      string `json:"thi_source"`
				PlatformSource string `json:"platform_source"`
			} `json:"file_struct"`
			Duration   int `json:"duration"`
			EventExtra struct {
			} `json:"event_extra"`
			ScrapeStatus            int    `json:"scrape_status"`
			UpdateViewAt            int64  `json:"update_view_at"`
			LastUpdateAt            int64  `json:"last_update_at"`
			ShareFidToken           string `json:"share_fid_token"`
			Ban                     bool   `json:"ban"`
			CurVersionOrDefault     int    `json:"cur_version_or_default"`
			RawNameSpace            int    `json:"raw_name_space"`
			OfflineSource           bool   `json:"offline_source"`
			BackupSource            bool   `json:"backup_source"`
			SaveAsSource            bool   `json:"save_as_source"`
			OwnerDriveTypeOrDefault int    `json:"owner_drive_type_or_default"`
			Dir                     bool   `json:"dir"`
			File                    bool   `json:"file"`
			CreatedAt               int64  `json:"created_at"`
			UpdatedAt               int64  `json:"updated_at"`
			Extra1                  struct {
			} `json:"_extra"`
			Thumbnail          string `json:"thumbnail,omitempty"`
			BigThumbnail       string `json:"big_thumbnail,omitempty"`
			PreviewUrl         string `json:"preview_url,omitempty"`
			VideoMaxResolution string `json:"video_max_resolution,omitempty"`
			VideoWidth         int    `json:"video_width,omitempty"`
			VideoHeight        int    `json:"video_height,omitempty"`
			VideoRotate        int    `json:"video_rotate,omitempty"`
			ObjCategory        string `json:"obj_category,omitempty"`
			FileNameHl         string `json:"file_name_hl,omitempty"`
			LastPlayInfo       struct {
				Time int `json:"time"`
			} `json:"last_play_info,omitempty"`
		} `json:"list"`
	} `json:"data"`
	Metadata struct {
		Size          int    `json:"_size"`
		Page          int    `json:"_page"`
		Count         int    `json:"_count"`
		Total         int    `json:"_total"`
		CheckFidToken int    `json:"check_fid_token"`
		GGroup        string `json:"_g_group"`
		TGroup        string `json:"_t_group"`
	} `json:"metadata"`
}
type MyNodeResp struct {
	Status    int    `json:"status"`
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Timestamp int    `json:"timestamp"`
	Data      struct {
		LastViewList   []interface{} `json:"last_view_list"`
		RecentFileList []interface{} `json:"recent_file_list"`
		List           []struct {
			Fid                 string `json:"fid"`
			FileName            string `json:"file_name"`
			PdirFid             string `json:"pdir_fid"`
			Category            int    `json:"category"`
			FileType            int    `json:"file_type"`
			Size                int    `json:"size"`
			FormatType          string `json:"format_type"`
			Status              int    `json:"status"`
			Tags                string `json:"tags,omitempty"`
			OwnerUcid           string `json:"owner_ucid"`
			LCreatedAt          int64  `json:"l_created_at"`
			LUpdatedAt          int64  `json:"l_updated_at"`
			Source              string `json:"source"`
			FileSource          string `json:"file_source"`
			NameSpace           int    `json:"name_space"`
			LShotAt             int64  `json:"l_shot_at"`
			SourceDisplay       string `json:"source_display"`
			IncludeItems        int    `json:"include_items,omitempty"`
			SeriesDir           bool   `json:"series_dir"`
			UploadCameraRootDir bool   `json:"upload_camera_root_dir"`
			Fps                 int    `json:"fps"`
			Like                int    `json:"like"`
			OperatedAt          int64  `json:"operated_at"`
			RiskType            int    `json:"risk_type"`
			BackupSign          int    `json:"backup_sign"`
			FileNameHlStart     int    `json:"file_name_hl_start"`
			FileNameHlEnd       int    `json:"file_name_hl_end"`
			Duration            int    `json:"duration"`
			EventExtra          struct {
				IsOpen          bool  `json:"is_open,omitempty"`
				RecentCreatedAt int64 `json:"recent_created_at,omitempty"`
				ViewAt          int64 `json:"view_at,omitempty"`
			} `json:"event_extra"`
			HasSubDirs              bool  `json:"has_sub_dirs,omitempty"`
			ScrapeStatus            int   `json:"scrape_status"`
			UpdateViewAt            int64 `json:"update_view_at"`
			Ban                     bool  `json:"ban"`
			CurVersionOrDefault     int   `json:"cur_version_or_default"`
			RawNameSpace            int   `json:"raw_name_space"`
			BackupSource            bool  `json:"backup_source"`
			OfflineSource           bool  `json:"offline_source"`
			SaveAsSource            bool  `json:"save_as_source"`
			OwnerDriveTypeOrDefault int   `json:"owner_drive_type_or_default"`
			Dir                     bool  `json:"dir"`
			File                    bool  `json:"file"`
			CreatedAt               int64 `json:"created_at"`
			UpdatedAt               int64 `json:"updated_at"`
			Extra                   struct {
			} `json:"_extra"`
			Thumbnail    string `json:"thumbnail,omitempty"`
			BigThumbnail string `json:"big_thumbnail,omitempty"`
			PreviewUrl   string `json:"preview_url,omitempty"`
			ObjCategory  string `json:"obj_category,omitempty"`
			LastPlayInfo struct {
				Time int `json:"time"`
			} `json:"last_play_info,omitempty"`
		} `json:"list"`
	} `json:"data"`
	Metadata struct {
		Size  int    `json:"_size"`
		ReqId string `json:"req_id"`
		Page  int    `json:"_page"`
		Count int    `json:"_count"`
		Total int    `json:"_total"`
	} `json:"metadata"`
}
type fileInfoResp struct {
	FidList      []string `json:"fid_list"`
	FidTokenList []string `json:"fid_token_list"`
}
type saveInfoResp struct {
	Status    int    `json:"status"`
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Timestamp int    `json:"timestamp"`
	Data      struct {
		TaskId string `json:"task_id"`
	} `json:"data"`
	Metadata struct {
		TqGap int `json:"tq_gap"`
	} `json:"metadata"`
}

type TaskInfo struct {
	Status    int    `json:"status"`
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Timestamp int    `json:"timestamp"`
	Data      struct {
		TaskId     string `json:"task_id"`
		EventId    string `json:"event_id"`
		TaskType   int    `json:"task_type"`
		TaskTitle  string `json:"task_title"`
		Status     int    `json:"status"`
		CreatedAt  int64  `json:"created_at"`
		FinishedAt int64  `json:"finished_at"`
		Share      struct {
		} `json:"share"`
		SaveAs struct {
			SearchExit          bool     `json:"search_exit"`
			ToPdirFid           string   `json:"to_pdir_fid"`
			SaveAsSumNum        int      `json:"save_as_sum_num"`
			IsPack              string   `json:"is_pack"`
			SaveAsTopFids       []string `json:"save_as_top_fids"`
			SaveAsSelectTopFids []string `json:"save_as_select_top_fids"`
			ToPdirName          string   `json:"to_pdir_name"`
		} `json:"save_as"`
	} `json:"data"`
	Metadata struct {
		TqGap int `json:"tq_gap"`
	} `json:"metadata"`
}
