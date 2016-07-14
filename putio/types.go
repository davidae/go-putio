package putio

import "fmt"

// File represents a Put.io file.
type File struct {
	ID                int    `json:"id"`
	Filename          string `json:"name"`
	Filesize          int64  `json:"size"`
	ContentType       string `json:"content_type"`
	CreatedAt         string `json:"created_at"`
	FirstAccessedAt   string `json:"first_accessed_at"`
	ParentID          int    `json:"parent_id"`
	Screenshot        string `json:"screenshot"`
	OpensubtitlesHash string `json:"opensubtitles_hash"`
	IsMP4Available    bool   `json:"is_mp4_available"`
	Icon              string `json:"icon"`
	CRC32             string `json:"crc32"`
	IsShared          bool   `json:"is_shared"`
}

// FileList represents a list of files of a Put.io directory.
type FileList struct {
	Files  []File `json:"file"`
	Parent File   `json:"parent"`
}

type Upload struct {
	File     *File     `json:"file"`
	Transfer *Transfer `json:"transfer"`
}

type Search struct {
	Files []File `json:"files"`
	Next  string `json:"next"`
}

type Transfer struct {
	Availability       string `json:"availability"`
	CallbackURL        string `json:"callback_url"`
	CreatedAt          string `json:"created_at"`
	CreatedTorrent     bool   `json:"created_torrent"`
	ClientIP           string `json:"client_ip"`
	CurrentRatio       string `json:"current_ratio"`
	DownloadSpeed      int    `json:"down_speed"`
	Downloaded         int    `json:"downloaded"`
	DownloadID         int    `json:"download_id"`
	ErrorMessage       string `json:"error_message"`
	EstimatedTime      string `json:"estimated_time"`
	Extract            bool   `json:"extract"`
	FileID             int    `json:"file_id"`
	FinishedAt         string `json:"finished_at"`
	ID                 int    `json:"id"`
	IsPrivate          bool   `json:"is_private"`
	MagnetURI          string `json:"magneturi"`
	Name               string `json:"name"`
	PeersConnected     int    `json:"peers_connected"`
	PeersGettingFromUs int    `json:"peers_getting_from_us"`
	PeersSendingToUs   int    `json:"peers_sending_to_us"`
	PercentDone        int    `json:"percent_done"`
	SaveParentID       int    `json:"save_parent_id"`
	SecondsSeeding     int    `json:"seconds_seeding"`
	Size               int    `json:"size"`
	Source             string `json:"source"`
	Status             string `json:"status"`
	StatusMessage      string `json:"status_message"`
	SubscriptionID     int    `json:"subscription_id"`
	TorrentLink        string `json:"torrent_link"`
	TrackerMessage     string `json:"tracker_message"`
	Trackers           string `json:"tracker"`
	Type               string `json:"type"`
	UploadSpeed        int    `json:"up_speed"`
	Uploaded           int    `json:"uploaded"`
}

type Info struct {
	AccountActive           bool   `json:"account_active"`
	AvatarURL               string `json:"avatar_url"`
	DaysUntilFilesDeletion  int    `json:"days_until_files_deletion"`
	DefaultSubtitleLanguage string `json:"default_subtitle_language"`
	Disk                    struct {
		Avail int `json:"avail"`
		Size  int `json:"size"`
		Used  int `json:"used"`
	} `json:"disk"`
	HasVoucher                int      `json:"has_voucher"`
	Mail                      string   `json:"mail"`
	PassiveAccount            bool     `json:"passive_account"`
	PlanExpirationDate        string   `json:"plan_expiration_date"`
	Settings                  Settings `json:"settings"`
	SimultaneousDownloadLimit int      `json:"simultaneous_download_limit"`
	SubtitleLanguages         []string `json:"subtitle_languages"`
	UserID                    int      `json:"user_id"`
	Username                  string   `json:"username"`
}

type Settings struct {
	CallbackURL             string      `json:"callback_url"`
	DefaultDownloadFolder   int         `json:"default_download_folder"`
	DefaultSubtitleLanguage string      `json:"default_subtitle_language"`
	DownloadFolderUnset     bool        `json:"download_folder_unset"`
	IsInvisible             bool        `json:"is_invisible"`
	Nextepisode             bool        `json:"nextepisode"`
	PrivateDownloadHostIP   interface{} `json:"private_download_host_ip"`
	PushoverToken           string      `json:"pushover_token"`
	Routing                 string      `json:"routing"`
	Sorting                 string      `json:"sorting"`
	SSLEnabled              bool        `json:"ssl_enabled"`
	StartFrom               bool        `json:"start_from"`
	SubtitleLanguages       []string    `json:"subtitle_languages"`
}

type Friend struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	AvatarURL string `json:"avatar_url"`
}

// FIXME: missing_files field is missin
type Zip struct {
	ID        int    `json:"id"`
	CreatedAt string `json:"created_at"`

	Size   int    `json:"size"`
	Status string `json:"status"`
	URL    string `json:"url"`
}

// errorResponse represents a common error message that Put.io v2 API sends on
// error.
type errorResponse struct {
	ErrorMessage string `json:"error_message"`
	ErrorType    string `json:"error_type"`
	ErrorURI     string `json:"error_uri"`
	Status       string `json:"status"`
	StatusCode   int    `json:"status_code"`
}

func (e errorResponse) Error() string {
	return fmt.Sprintf("StatusCode: %v ErrorType: %v ErrorMsg: %v", e.StatusCode, e.ErrorType, e.ErrorMessage)
}