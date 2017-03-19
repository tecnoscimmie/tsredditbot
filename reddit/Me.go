package reddit

// Me is what reddit replies when poking /api/v1/me
type Me struct {
	IsEmployee bool `json:"is_employee"`
	Features   struct {
		NewUserOnboarding struct {
			Owner        string `json:"owner"`
			Variant      string `json:"variant"`
			ExperimentID int    `json:"experiment_id"`
		} `json:"new_user_onboarding"`
		LiveHappeningNow    bool `json:"live_happening_now"`
		AdserverReporting   bool `json:"adserver_reporting"`
		LegacySearchPref    bool `json:"legacy_search_pref"`
		MobileWebTargeting  bool `json:"mobile_web_targeting"`
		AdsAuction          bool `json:"ads_auction"`
		RulesModalOnComment struct {
			Owner        string `json:"owner"`
			Variant      string `json:"variant"`
			ExperimentID int    `json:"experiment_id"`
		} `json:"rules_modal_on_comment"`
		AdzerkDoNotTrack      bool `json:"adzerk_do_not_track"`
		WhitelistedPms        bool `json:"whitelisted_pms"`
		ShowRulesOnSubmitPage struct {
			Owner        string `json:"owner"`
			Variant      string `json:"variant"`
			ExperimentID int    `json:"experiment_id"`
		} `json:"show_rules_on_submit_page"`
		StickyComments       bool `json:"sticky_comments"`
		UpgradeCookies       bool `json:"upgrade_cookies"`
		AdsPrefs             bool `json:"ads_prefs"`
		BlockUserByReport    bool `json:"block_user_by_report"`
		AdsAutoRefund        bool `json:"ads_auto_refund"`
		OrangeredsAsEmails   bool `json:"orangereds_as_emails"`
		ExpandoEvents        bool `json:"expando_events"`
		EuCookiePolicy       bool `json:"eu_cookie_policy"`
		ProgrammaticAds      bool `json:"programmatic_ads"`
		ForceHTTPS           bool `json:"force_https"`
		ActivityServiceWrite bool `json:"activity_service_write"`
		PokemongoContent     struct {
			Owner        string `json:"owner"`
			Variant      string `json:"variant"`
			ExperimentID int    `json:"experiment_id"`
		} `json:"pokemongo_content"`
		DoNotTrack                    bool `json:"do_not_track"`
		ReddituploadsRedirect         bool `json:"reddituploads_redirect"`
		OutboundClicktracking         bool `json:"outbound_clicktracking"`
		ImageUploads                  bool `json:"image_uploads"`
		MwebXpromoRequireLoginAndroid struct {
			Owner        string `json:"owner"`
			Variant      string `json:"variant"`
			ExperimentID int    `json:"experiment_id"`
		} `json:"mweb_xpromo_require_login_android"`
		NewLoggedinCachePolicy                bool `json:"new_loggedin_cache_policy"`
		HTTPSRedirect                         bool `json:"https_redirect"`
		MwebXpromoInterstitialCommentsIos     bool `json:"mweb_xpromo_interstitial_comments_ios"`
		LiveOrangereds                        bool `json:"live_orangereds"`
		MwebXpromoListingClickEveryTime       bool `json:"mweb_xpromo_listing_click_every_time"`
		UtmCommentLinks                       bool `json:"utm_comment_links"`
		GiveHstsGrants                        bool `json:"give_hsts_grants"`
		PauseAds                              bool `json:"pause_ads"`
		ShowRecommendedLink                   bool `json:"show_recommended_link"`
		MobileNativeBanner                    bool `json:"mobile_native_banner"`
		MwebXpromoInterstitialCommentsAndroid bool `json:"mweb_xpromo_interstitial_comments_android"`
		ScreenviewEvents                      bool `json:"screenview_events"`
		NewReportDialog                       bool `json:"new_report_dialog"`
		MoatTracking                          bool `json:"moat_tracking"`
		SubredditRules                        bool `json:"subreddit_rules"`
		RulesModalOnSubmit                    struct {
			Owner        string `json:"owner"`
			Variant      string `json:"variant"`
			ExperimentID int    `json:"experiment_id"`
		} `json:"rules_modal_on_submit"`
		AdzerkReporting2    bool `json:"adzerk_reporting_2"`
		InboxPush           bool `json:"inbox_push"`
		AdsAutoExtend       bool `json:"ads_auto_extend"`
		InterestTargeting   bool `json:"interest_targeting"`
		PostEmbed           bool `json:"post_embed"`
		MobileSettings      bool `json:"mobile_settings"`
		ScrollEvents        bool `json:"scroll_events"`
		AdblockTest         bool `json:"adblock_test"`
		ActivityServiceRead bool `json:"activity_service_read"`
	} `json:"features"`
	IsSuspended             bool        `json:"is_suspended"`
	GoldExpiration          interface{} `json:"gold_expiration"`
	ID                      string      `json:"id"`
	SuspensionExpirationUtc interface{} `json:"suspension_expiration_utc"`
	NewModmailExists        bool        `json:"new_modmail_exists"`
	Over18                  bool        `json:"over_18"`
	IsGold                  bool        `json:"is_gold"`
	IsMod                   bool        `json:"is_mod"`
	HasVerifiedEmail        bool        `json:"has_verified_email"`
	Email                   string      `json:"email"`
	HasModMail              bool        `json:"has_mod_mail"`
	OauthClientID           string      `json:"oauth_client_id"`
	HideFromRobots          bool        `json:"hide_from_robots"`
	LinkKarma               int         `json:"link_karma"`
	InboxCount              int         `json:"inbox_count"`
	HasMail                 bool        `json:"has_mail"`
	Name                    string      `json:"name"`
	Created                 float64     `json:"created"`
	GoldCreddits            int         `json:"gold_creddits"`
	CreatedUtc              float64     `json:"created_utc"`
	InBeta                  bool        `json:"in_beta"`
	CommentKarma            int         `json:"comment_karma"`
}
