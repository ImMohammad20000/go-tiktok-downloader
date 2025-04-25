package main

type T struct {
	DefaultScope DefaultScope `json:"__DEFAULT_SCOPE__"`
}

type DefaultScope struct {
	WebappAppContext      WebappAppContext      `json:"webapp.app-context"`
	WebappBizContext      WebappBizContext      `json:"webapp.biz-context"`
	WebappI18NTranslation WebappI18NTranslation `json:"webapp.i18n-translation"`
	SEOAbtest             SEOAbtest             `json:"seo.abtest"`
	WebappVideoDetail     WebappVideoDetail     `json:"webapp.video-detail"`
	WebappAB              WebappAB              `json:"webapp.a-b"`
}

type SEOAbtest struct {
	Canonical  string              `json:"canonical"`
	PageID     string              `json:"pageId"`
	VidList    []interface{}       `json:"vidList"`
	Parameters SEOAbtestParameters `json:"parameters"`
}

type SEOAbtestParameters struct {
	AddInfoCardSEO             AddInfoCardSEO             `json:"add_info_card_seo"`
	AddTranscriptSEO           AddInfoCardSEO             `json:"add_transcript_seo"`
	BotCostOptimize            BotCostOptimize            `json:"bot_cost_optimize"`
	SEOGenericExposureCapacity SEOGenericExposureCapacity `json:"seo_generic_exposure_capacity"`
	UpdateCommentsNumber       AddInfoCardSEO             `json:"update_comments_number"`
	VideoNonTdkPhase2          AddInfoCardSEO             `json:"video_non_tdk_phase2"`
	VideoPageInfoCardBotExp    AddInfoCardSEO             `json:"video_page_info_card_bot_exp"`
	VideoPageSerpCompliance    AddInfoCardSEO             `json:"video_page_serp_compliance"`
	VideoTdkPhase2             AddInfoCardSEO             `json:"video_tdk_phase2"`
}

type AddInfoCardSEO struct {
	Vid string `json:"vid"`
}

type BotCostOptimize struct {
	SimplifyLink      bool `json:"simplify_link"`
	SimplifyVideoDesc bool `json:"simplify_video_desc"`
}

type SEOGenericExposureCapacity struct {
	Count int64  `json:"count"`
	Vid   string `json:"vid"`
}

type WebappAB struct {
	BC string `json:"b_c"`
}

type WebappAppContext struct {
	Language         string        `json:"language"`
	Region           string        `json:"region"`
	AppID            int64         `json:"appId"`
	AppType          string        `json:"appType"`
	Wid              string        `json:"wid"`
	WebIDCreatedTime string        `json:"webIdCreatedTime"`
	OdinID           string        `json:"odinId"`
	Nonce            string        `json:"nonce"`
	BotType          string        `json:"botType"`
	RequestID        string        `json:"requestId"`
	ClusterRegion    string        `json:"clusterRegion"`
	AbTestVersion    AbTestVersion `json:"abTestVersion"`
	CSRFToken        string        `json:"csrfToken"`
	UserAgent        string        `json:"userAgent"`
	EncryptedWebid   string        `json:"encryptedWebid"`
	Host             string        `json:"host"`
}

type AbTestVersion struct {
	VersionName string                  `json:"versionName"`
	Parameters  AbTestVersionParameters `json:"parameters"`
	AbTestApp   AbTestApp               `json:"abTestApp"`
}

type AbTestApp struct {
	Parameters AbTestAppParameters `json:"parameters"`
}

type AbTestAppParameters struct {
	Tiktok PurpleTiktok `json:"tiktok"`
}

type PurpleTiktok struct {
	LongVideoPopupDisplayOptimization bool `json:"long_video_popup_display_optimization"`
}

type AbTestVersionParameters struct {
	MobileConsumptionLimitNonLoggedIn    AddInfoCardSEO      `json:"mobile_consumption_limit_non_logged_in"`
	MobileConsumptionLimitLoggedIn       AddInfoCardSEO      `json:"mobile_consumption_limit_logged_in"`
	CreatorCenterConnect                 AddInfoCardSEO      `json:"creator_center_connect"`
	QrSsoPopup                           AddInfoCardSEO      `json:"qr_sso_popup"`
	WebappSwitchAccount                  AddInfoCardSEO      `json:"webapp_switch_account"`
	MobilePredictiveData                 AddInfoCardSEO      `json:"mobile_predictive_data"`
	ConfirmLogout                        AddInfoCardSEO      `json:"confirm_logout"`
	OneColumnPlayerSize                  AddInfoCardSEO      `json:"one_column_player_size"`
	RemoveBottomBanner                   AddInfoCardSEO      `json:"remove_bottom_banner"`
	UseInboxNoticeCountAPI               AddInfoCardSEO      `json:"use_inbox_notice_count_api"`
	PeriodicLoginPopupInterval           AddInfoCardSEO      `json:"periodic_login_popup_interval"`
	MobileVodkit                         AddInfoCardSEO      `json:"mobile_vodkit"`
	VideoBitrateAdapt                    AddInfoCardSEO      `json:"video_bitrate_adapt"`
	UseFollowV2                          AddInfoCardSEO      `json:"use_follow_v2"`
	SearchVideo                          AddInfoCardSEO      `json:"search_video"`
	NonLoggedInComments                  AddInfoCardSEO      `json:"non_logged_in_comments"`
	LoginModalUIRevamp                   AddInfoCardSEO      `json:"login_modal_ui_revamp"`
	BrowserLoginRedirect                 AddInfoCardSEO      `json:"browser_login_redirect"`
	LastLoginMethod                      AddInfoCardSEO      `json:"last_login_method"`
	MobileConsumptionLimitLogin          AddInfoCardSEO      `json:"mobile_consumption_limit_login"`
	XgplayerPreloadConfig                AddInfoCardSEO      `json:"xgplayer_preload_config"`
	EnableMlModel                        AddInfoCardSEO      `json:"enable_ml_model"`
	XgVolumeTest                         AddInfoCardSEO      `json:"xg_volume_test"`
	ShareButtonPart1Test                 AddInfoCardSEO      `json:"share_button_part1_test"`
	VideoServerpush                      AddInfoCardSEO      `json:"video_serverpush"`
	BrowserModeEncourageLogin            AddInfoCardSEO      `json:"browser_mode_encourage_login"`
	MobileSearchTest                     AddInfoCardSEO      `json:"mobile_search_test"`
	VolumeNormalize                      AddInfoCardSEO      `json:"volume_normalize"`
	WebappLoginEmailPhone                AddInfoCardSEO      `json:"webapp_login_email_phone"`
	SignUpWebappRegionChange             AddInfoCardSEO      `json:"sign_up_webapp_region_change"`
	LoginOptionOrderByMetrics            AddInfoCardSEO      `json:"login_option_order_by_metrics"`
	ZTITest                              ZTITest             `json:"ZTI_test"`
	AbTag                                AbTag               `json:"ab_tag"`
	AbTags                               AbTags              `json:"ab_tags"`
	AddKapEntry                          AddInfoCardSEO      `json:"add_kap_entry"`
	AddProfileLeftBar                    AddInfoCardSEO      `json:"add_profile_left_bar"`
	AddTranscript                        AddInfoCardSEO      `json:"add_transcript"`
	AnalyticsUpgradePhase3               AddInfoCardSEO      `json:"analytics_upgrade_phase3"`
	AutoDarkMode                         AddInfoCardSEO      `json:"auto_dark_mode"`
	AutoScroll                           AddInfoCardSEO      `json:"auto_scroll"`
	AutoscrollReposition                 AddInfoCardSEO      `json:"autoscroll_reposition"`
	BrowseModeAutoplayTest               AddInfoCardSEO      `json:"browse_mode_autoplay_test"`
	BrowserModeCreatorTab3               AddInfoCardSEO      `json:"browser_mode_creator_tab_3"`
	CcPerfPhase1                         CcPerfPhase1        `json:"cc_perf_phase1"`
	ChangeListLengthNew                  AddInfoCardSEO      `json:"change_list_length_new"`
	CleanFixedBottom                     AddInfoCardSEO      `json:"clean_fixed_bottom"`
	CommentRefactorTest                  AddInfoCardSEO      `json:"comment_refactor_test"`
	CreatorCenterConnectGlobal           AddInfoCardSEO      `json:"creator_center_connect_global"`
	CreatorCenterGlobalCommentManagement AddInfoCardSEO      `json:"creator_center_global_comment_management"`
	CreatorCenterGlobalPostManagement    AddInfoCardSEO      `json:"creator_center_global_post_management"`
	CreatorCenterTest                    AddInfoCardSEO      `json:"creator_center_test"`
	DelayGuest                           AddInfoCardSEO      `json:"delay_guest"`
	DesktopUIOpt                         AddInfoCardSEO      `json:"desktop_ui_opt"`
	DesktopUIReply                       AddInfoCardSEO      `json:"desktop_ui_reply"`
	DetailPageCommentsRedesign           AddInfoCardSEO      `json:"detail_page_comments_redesign"`
	DigitalWellbeingWeb                  AddInfoCardSEO      `json:"digital_wellbeing_web"`
	EnableAboutThisAd                    AddInfoCardSEO      `json:"enable_about_this_ad"`
	EnableAds                            AddInfoCardSEO      `json:"enable_ads"`
	EnableAutoscrollMoremenu             AddInfoCardSEO      `json:"enable_autoscroll_moremenu"`
	EnableContinuePlay                   AddInfoCardSEO      `json:"enable_continue_play"`
	EnableEcLcc                          AddInfoCardSEO      `json:"enable_ec_lcc"`
	EnableFbSDK                          AddInfoCardSEO      `json:"enable_fb_sdk"`
	EnableMessageRefactor                AddInfoCardSEO      `json:"enable_message_refactor"`
	EnableMiniPlayer                     AddInfoCardSEO      `json:"enable_mini_player"`
	EnableMiniPlayerNewDesign            AddInfoCardSEO      `json:"enable_mini_player_new_design"`
	EnableMiniPlayerSwitchTabPopup       AddInfoCardSEO      `json:"enable_mini_player_switch_tab_popup"`
	EnableNotInterested                  AddInfoCardSEO      `json:"enable_not_interested"`
	EnablePostTranslation                AddInfoCardSEO      `json:"enable_post_translation"`
	EnableProfilePinnedVideo             AddInfoCardSEO      `json:"enable_profile_pinned_video"`
	EnableVideoDetailMoremenuRefactor    AddInfoCardSEO      `json:"enable_video_detail_moremenu_refactor"`
	EnhanceVideoConsumptionTest          AddInfoCardSEO      `json:"enhance_video_consumption_test"`
	ExchangeRetentionPopup               AddInfoCardSEO      `json:"exchange_retention_popup"`
	ExpandItemTag                        AddInfoCardSEO      `json:"expand_item_tag"`
	ExploreAllTab                        AddInfoCardSEO      `json:"explore_all_tab"`
	ExploreTest                          AddInfoCardSEO      `json:"explore_test"`
	ExploreTrendingTopics                AddInfoCardSEO      `json:"explore_trending_topics"`
	ExploreUIChange                      AddInfoCardSEO      `json:"explore_ui_change"`
	FavoriteTest                         AddInfoCardSEO      `json:"favorite_test"`
	FeedChangeOptimizeFf                 AddInfoCardSEO      `json:"feed_change_optimize_ff"`
	FeedScrollOpt                        AddInfoCardSEO      `json:"feed_scroll_opt"`
	FixTeaSession                        AddInfoCardSEO      `json:"fix_tea_session"`
	FollowingDisplayLive                 AddInfoCardSEO      `json:"following_display_live"`
	ForyouPrefetch                       AddInfoCardSEO      `json:"foryou_prefetch"`
	FriendsTab                           AddInfoCardSEO      `json:"friends_tab"`
	FypBackupV2                          FypBackupV2         `json:"fyp_backup_v2"`
	FypCommentsPanel                     AddInfoCardSEO      `json:"fyp_comments_panel"`
	FypHideMusicInfo                     AddInfoCardSEO      `json:"fyp_hide_music_info"`
	FypProgressBar                       AddInfoCardSEO      `json:"fyp_progress_bar"`
	FypRedesign                          AddInfoCardSEO      `json:"fyp_redesign"`
	FypUpdateGradient                    AddInfoCardSEO      `json:"fyp_update_gradient"`
	GlobalWebFooter                      AddInfoCardSEO      `json:"global_web_footer"`
	GuestModeRedesign                    AddInfoCardSEO      `json:"guest_mode_redesign"`
	GuideUserToNextVideo                 AddInfoCardSEO      `json:"guide_user_to_next_video"`
	IncreaseDetailPageCoverQuantityTest  AddInfoCardSEO      `json:"increase_detail_page_cover_quantity_test"`
	IslandsArchExplore                   AddInfoCardSEO      `json:"islands_arch_explore"`
	IslandsArchPhase2                    AddInfoCardSEO      `json:"islands_arch_phase2"`
	IslandsArchRESTPage                  AddInfoCardSEO      `json:"islands_arch_rest_page"`
	IslandsArchUserProfile               AddInfoCardSEO      `json:"islands_arch_user_profile"`
	IslandsArchVideoDetail               AddInfoCardSEO      `json:"islands_arch_video_detail"`
	KepNewUILogin                        AddInfoCardSEO      `json:"kep_new_ui_login"`
	LiveABRParam                         AddInfoCardSEO      `json:"live_abr_param"`
	LiveABRParamCold                     AddInfoCardSEO      `json:"live_abr_param_cold"`
	LiveABRVersion                       AddInfoCardSEO      `json:"live_abr_version"`
	LiveChatMessageCacheOpt              AddInfoCardSEO      `json:"live_chat_message_cache_opt"`
	LiveContextOptimize                  AddInfoCardSEO      `json:"live_context_optimize"`
	LiveCsrInsertContext                 AddInfoCardSEO      `json:"live_csr_insert_context"`
	LiveEndImprovedMetrics               AddInfoCardSEO      `json:"live_end_improved_metrics"`
	LiveEnterRoomChore                   AddInfoCardSEO      `json:"live_enter_room_chore"`
	LiveEventAggregation                 AddInfoCardSEO      `json:"live_event_aggregation"`
	LiveFeedSave                         AddInfoCardSEO      `json:"live_feed_save"`
	LiveFeedStyle                        AddInfoCardSEO      `json:"live_feed_style"`
	LiveFypRedesignA                     AddInfoCardSEO      `json:"live_fyp_redesign_a"`
	LiveGiftAddAnimation                 AddInfoCardSEO      `json:"live_gift_add_animation"`
	LiveGiftAddAnimationSendLoading      AddInfoCardSEO      `json:"live_gift_add_animation_send_loading"`
	LiveGiftLoggedOutControl             AddInfoCardSEO      `json:"live_gift_logged_out_control"`
	LiveGoliveEntrance                   AddInfoCardSEO      `json:"live_golive_entrance"`
	LiveHeaderDelay                      AddInfoCardSEO      `json:"live_header_delay"`
	LiveI18NReduce                       AddInfoCardSEO      `json:"live_i18n_reduce"`
	LiveLcpPerfOptimize                  AddInfoCardSEO      `json:"live_lcp_perf_optimize"`
	LiveLike                             AddInfoCardSEO      `json:"live_like"`
	LiveLoginReflowBtn                   AddInfoCardSEO      `json:"live_login_reflow_btn"`
	LiveLowLatency                       AddInfoCardSEO      `json:"live_low_latency"`
	LiveMultiStreamSwitch                AddInfoCardSEO      `json:"live_multi_stream_switch"`
	LiveNewDiscover                      AddInfoCardSEO      `json:"live_new_discover"`
	LivePlayerIcon                       AddInfoCardSEO      `json:"live_player_icon"`
	LivePlayerMuteText                   AddInfoCardSEO      `json:"live_player_mute_text"`
	LivePlayerSwitchButton               AddInfoCardSEO      `json:"live_player_switch_button"`
	LivePreviewWeb                       AddInfoCardSEO      `json:"live_preview_web"`
	LiveProShow                          AddInfoCardSEO      `json:"live_pro_show"`
	LiveQuickChatExpand                  AddInfoCardSEO      `json:"live_quick_chat_expand"`
	LiveRechargeByAmount                 AddInfoCardSEO      `json:"live_recharge_by_amount"`
	LiveRechargeCoinNewUIPC              AddInfoCardSEO      `json:"live_recharge_coin_new_ui_pc"`
	LiveRechargeHomescreen               AddInfoCardSEO      `json:"live_recharge_homescreen"`
	LiveRoomAgeRestriction               AddInfoCardSEO      `json:"live_room_age_restriction"`
	LiveRoomCsr                          AddInfoCardSEO      `json:"live_room_csr"`
	LiveRoomMatch                        AddInfoCardSEO      `json:"live_room_match"`
	LiveRoomNonStreaming                 AddInfoCardSEO      `json:"live_room_non_streaming"`
	LiveSearchJump                       AddInfoCardSEO      `json:"live_search_jump"`
	LiveStudioDownloadRefactorPC         AddInfoCardSEO      `json:"live_studio_download_refactor_pc"`
	LiveStudioDownloadType               AddInfoCardSEO      `json:"live_studio_download_type"`
	LiveSubscriptionCashier              AddInfoCardSEO      `json:"live_subscription_cashier"`
	LiveTopViewers                       AddInfoCardSEO      `json:"live_top_viewers"`
	LiveVolumeBalance                    LiveVolumeBalance   `json:"live_volume_balance"`
	LiveWalletPerformancePackup          AddInfoCardSEO      `json:"live_wallet_performance_packup"`
	LiveWebcodecsPlayer                  AddInfoCardSEO      `json:"live_webcodecs_player"`
	NavPhase3                            AddInfoCardSEO      `json:"nav_phase_3"`
	NewLiveIMHooks                       AddInfoCardSEO      `json:"new_Live_im_hooks"`
	NewItemTag                           AddInfoCardSEO      `json:"new_item_tag"`
	NewServerCall                        AddInfoCardSEO      `json:"new_server_call"`
	NonPersonalizedFeedsWeb              AddInfoCardSEO      `json:"non_personalized_feeds_web"`
	OptimiseBrowserMode                  AddInfoCardSEO      `json:"optimise_browser_mode"`
	PCInspiration                        AddInfoCardSEO      `json:"pc_inspiration"`
	PCNonPersonalizedExplore             AddInfoCardSEO      `json:"pc_non_personalized_explore"`
	PCNonPersonalizedSuggestedAccount    AddInfoCardSEO      `json:"pc_non_personalized_suggested_account"`
	PCVideoPlaylistTest                  AddInfoCardSEO      `json:"pc_video_playlist_test"`
	PhotoModeYml                         AddInfoCardSEO      `json:"photo_mode_yml"`
	PhotoTest                            AddInfoCardSEO      `json:"photo_test"`
	PlaybackRefactorDesktop              AddInfoCardSEO      `json:"playback_refactor_desktop"`
	PlayerDegrade                        AddInfoCardSEO      `json:"player_degrade"`
	PlayerErrorOptimize                  AddInfoCardSEO      `json:"player_error_optimize"`
	PlayerRetry                          AddInfoCardSEO      `json:"player_retry"`
	PnsKeywordFiltering                  AddInfoCardSEO      `json:"pns_keyword_filtering"`
	PnsKrConsent                         AddInfoCardSEO      `json:"pns_kr_consent"`
	PnsPopupSDK                          AddInfoCardSEO      `json:"pns_popup_sdk"`
	PreviewCover                         AddInfoCardSEO      `json:"preview_cover"`
	ProfileFollowInfo                    AddInfoCardSEO      `json:"profile_follow_info"`
	PromoteQrCode                        AddInfoCardSEO      `json:"promote_qr_code"`
	ReduceUserItemList                   AddInfoCardSEO      `json:"reduce_user_item_list"`
	RemoveDisclaimer                     AddInfoCardSEO      `json:"remove_disclaimer"`
	RemoveTooltip                        AddInfoCardSEO      `json:"remove_tooltip"`
	ReportItemTag                        AddInfoCardSEO      `json:"report_item_tag"`
	RevampShareMenu                      AddInfoCardSEO      `json:"revamp_share_menu"`
	ReverseExpandItemTag                 AddInfoCardSEO      `json:"reverse_expand_item_tag"`
	Sati                                 Sati                `json:"sati"`
	ScheduledBreaksTeens                 AddInfoCardSEO      `json:"scheduled_breaks_teens"`
	SearchAddLive                        AddInfoCardSEO      `json:"search_add_live"`
	SearchAddNonPersonalizedSwitch       AddInfoCardSEO      `json:"search_add_non_personalized_switch"`
	SearchAddRelatedSearch               AddInfoCardSEO      `json:"search_add_related_search"`
	SearchBarStyleOpt                    AddInfoCardSEO      `json:"search_bar_style_opt"`
	SearchEntryCommentTop                AddInfoCardSEO      `json:"search_entry_comment_top"`
	SearchEntryCommentWord               AddInfoCardSEO      `json:"search_entry_comment_word"`
	SearchEntrySearchBar                 AddInfoCardSEO      `json:"search_entry_search_bar"`
	SearchKeepSugShow                    AddInfoCardSEO      `json:"search_keep_sug_show"`
	SearchPreviewUIChange                AddInfoCardSEO      `json:"search_preview_ui_change"`
	SearchTopAuthorCard                  AddInfoCardSEO      `json:"search_top_author_card"`
	SearchTransferGuesssearch            AddInfoCardSEO      `json:"search_transfer_guesssearch"`
	SearchTransferHistory                AddInfoCardSEO      `json:"search_transfer_history"`
	SearchVideoLab                       AddInfoCardSEO      `json:"search_video_lab"`
	SEOBreadcrumbDetail                  AddInfoCardSEO      `json:"seo_breadcrumb_detail"`
	SEODesktop                           AddInfoCardSEO      `json:"seo_desktop"`
	SEOPreviewUIChange                   AddInfoCardSEO      `json:"seo_preview_ui_change"`
	ShouldHighlightHashtag               AddInfoCardSEO      `json:"should_highlight_hashtag"`
	ShouldRecomReduceIconRisk            AddInfoCardSEO      `json:"should_recom_reduce_icon_risk"`
	ShowAigcLabelWeb                     AddInfoCardSEO      `json:"show_aigc_label_web"`
	SidenavTest                          AddInfoCardSEO      `json:"sidenav_test"`
	SolariaPortraitService               AddInfoCardSEO      `json:"solaria_portrait_service"`
	StudioWebAdvancedVideoPlayer         AddInfoCardSEO      `json:"studio_web_advanced_video_player"`
	StudioWebEhEntrance                  AddInfoCardSEO      `json:"studio_web_eh_entrance"`
	StudioWebEhEntranceV2                AddInfoCardSEO      `json:"studio_web_eh_entrance_v2"`
	Tikcast                              Tikcast             `json:"tikcast"`
	Tiktok                               FluffyTiktok        `json:"tiktok"`
	TiktokWeb                            TiktokWeb           `json:"tiktok_web"`
	TtPlayerLogger                       AddInfoCardSEO      `json:"tt_player_logger"`
	TtPlayerPreload                      TtPlayerPreload     `json:"tt_player_preload"`
	TtehEffectAnchorV1                   AddInfoCardSEO      `json:"tteh_effect_anchor_v1"`
	TtliveBroadcastTopicVersionTwo       AddInfoCardSEO      `json:"ttlive_broadcast_topic_version_two"`
	UILayoutAlignment                    AddInfoCardSEO      `json:"ui_layout_alignment"`
	UseAlignedCopies                     AddInfoCardSEO      `json:"use_aligned_copies"`
	UseErrorBoundary                     AddInfoCardSEO      `json:"use_error_boundary"`
	UseLeftNavigationRefactor            AddInfoCardSEO      `json:"use_left_navigation_refactor"`
	UseNavigationRefactor                AddInfoCardSEO      `json:"use_navigation_refactor"`
	VideoClosedCaption                   AddInfoCardSEO      `json:"video_closed_caption"`
	VideoDetailAuthorCard                AddInfoCardSEO      `json:"video_detail_author_card"`
	VideoDetailAutoPipOpt                AddInfoCardSEO      `json:"video_detail_auto_pip_opt"`
	VideoDetailNavOpt                    AddInfoCardSEO      `json:"video_detail_nav_opt"`
	VideoDetailPageVideoPlay             AddInfoCardSEO      `json:"video_detail_page_video_play"`
	VideoDetailRelatedRefetch            AddInfoCardSEO      `json:"video_detail_related_refetch"`
	VideoDetailResponsiveUI              AddInfoCardSEO      `json:"video_detail_responsive_ui"`
	VideoDetailSearchBar                 AddInfoCardSEO      `json:"video_detail_search_bar"`
	VideoDetailYmlUI                     AddInfoCardSEO      `json:"video_detail_yml_ui"`
	VvAvgPerDayPortrait                  VvAvgPerDayPortrait `json:"vv_avg_per_day_portrait"`
	WebCreationVmok                      AddInfoCardSEO      `json:"web_creation_vmok"`
	WebPlayerRefactor                    AddInfoCardSEO      `json:"web_player_refactor"`
	WebappAutoDarkMode                   AddInfoCardSEO      `json:"webapp_auto_dark_mode"`
	WebappCollectionProfile              AddInfoCardSEO      `json:"webapp_collection_profile"`
	WebappCreatorJustWatched             AddInfoCardSEO      `json:"webapp_creator_just_watched"`
	WebappCreatorPostSort                AddInfoCardSEO      `json:"webapp_creator_post_sort"`
	WebappExploreCategory                AddInfoCardSEO      `json:"webapp_explore_category"`
	WebappExploreNavOrder                AddInfoCardSEO      `json:"webapp_explore_nav_order"`
	WebappExploreVideoInfo               AddInfoCardSEO      `json:"webapp_explore_video_info"`
	WebappHeaderLsEntrance               AddInfoCardSEO      `json:"webapp_header_ls_entrance"`
	WebappInappNotice                    AddInfoCardSEO      `json:"webapp_inapp_notice"`
	WebappJotaiDetail                    AddInfoCardSEO      `json:"webapp_jotai_detail"`
	WebappJotaiForyou                    AddInfoCardSEO      `json:"webapp_jotai_foryou"`
	WebappLiveNavAvatar                  AddInfoCardSEO      `json:"webapp_live_nav_avatar"`
	WebappLiveRecommandSwitch            AddInfoCardSEO      `json:"webapp_live_recommand_switch"`
	WebappMasterOdinID                   AddInfoCardSEO      `json:"webapp_master_odin_id"`
	WebappPreviewCover                   AddInfoCardSEO      `json:"webapp_preview_cover"`
	WebappRecommendLanguage              AddInfoCardSEO      `json:"webapp_recommend_language"`
	WebappRepostAction                   AddInfoCardSEO      `json:"webapp_repost_action"`
	WebappRepostLabel                    AddInfoCardSEO      `json:"webapp_repost_label"`
	WebappRepostNotice                   AddInfoCardSEO      `json:"webapp_repost_notice"`
	WebappRepostTab                      AddInfoCardSEO      `json:"webapp_repost_tab"`
	WebappSEOPhotomodeUserExp            AddInfoCardSEO      `json:"webapp_seo_photomode_user_exp"`
	WebappVideoDetailPageRelatedMask     AddInfoCardSEO      `json:"webapp_video_detail_page_related_mask"`
	Webcast                              Webcast             `json:"webcast"`
	YmlUIOptimize                        AddInfoCardSEO      `json:"yml_ui_optimize"`
}

type AbTag struct {
	MergeRiskEvent int64 `json:"merge_risk_event"`
}

type AbTags struct {
	PrefetchVisaCryptogram bool `json:"prefetch_visa_cryptogram"`
}

type CcPerfPhase1 struct {
	Vid CcPerfPhase1Vid `json:"vid"`
}

type CcPerfPhase1Vid struct {
	IsCreatorCenterContextEnabled   bool `json:"isCreatorCenterContextEnabled"`
	IsInsightV2Enabled              bool `json:"isInsightV2Enabled"`
	IsOfflineI18NEnabled            bool `json:"isOfflineI18nEnabled"`
	IsParallelIframeEnabled         bool `json:"isParallelIframeEnabled"`
	IsPhase2Enabled                 bool `json:"isPhase2Enabled"`
	IsPrefetchIframeResourceEnabled bool `json:"isPrefetchIframeResourceEnabled"`
	IsServerSideTranslationEnabled  bool `json:"isServerSideTranslationEnabled"`
}

type FypBackupV2 struct {
	Enable string `json:"enable"`
}

type LiveVolumeBalance struct {
	Diff    int64  `json:"diff"`
	MaxDiff int64  `json:"maxDiff"`
	Target  int64  `json:"target"`
	Vid     string `json:"vid"`
}

type Sati struct {
	EnableGlobalReservePrice bool            `json:"enable_global_reserve_price"`
	GlobalReservePrice       int64           `json:"global_reserve_price"`
	IdtTiktokPCFeed          IdtTiktokPCFeed `json:"idt_tiktok_pc_feed"`
}

type IdtTiktokPCFeed struct {
	FrameworkMinAdGap            int64 `json:"framework_min_ad_gap"`
	EnablePadsGapFusion          bool  `json:"enable_pads_gap_fusion"`
	EnableUavGapFusion           bool  `json:"enable_uav_gap_fusion"`
	EnableThresholdPadsGapFusion bool  `json:"enable_threshold_pads_gap_fusion"`
	ThresholdPadsGapFusionWeight int64 `json:"threshold_pads_gap_fusion_weight"`
	ThresholdPadsGapFusionBias   int64 `json:"threshold_pads_gap_fusion_bias"`
}

type Tikcast struct {
	EnableEcLcc        bool `json:"enable_ec_lcc"`
	LiveRestrictedMode bool `json:"live_restricted_mode"`
}

type FluffyTiktok struct {
	DisableEffectFilter        int64        `json:"disable_effect_filter"`
	LiveRestrictedMode         bool         `json:"live_restricted_mode"`
	PrivateAccountPromptForU18 int64        `json:"private_account_prompt_for_u18"`
	SearchEngine               SearchEngine `json:"search_engine"`
	UseAlignedCopies           int64        `json:"use_aligned_copies"`
}

type SearchEngine struct {
	DebugUnlimitedMusicRetrievalCopyrightCodes []string `json:"debug_unlimited_music_retrieval_copyright_codes"`
	EnableTiktokStudioUnlimitedMusicSearch     int64    `json:"enable_tiktok_studio_unlimited_music_search"`
	EnableUnlimitedMusicRetrieval              int64    `json:"enable_unlimited_music_retrieval"`
	IesMTMusicDisplayUnlimitedMusicTag         int64    `json:"ies_mt_music_display_unlimited_music_tag"`
}

type TiktokWeb struct {
	AsyncPost                          int64  `json:"async_post"`
	CapcutEntryGroup                   int64  `json:"capcut_entry_group"`
	CopyrightPrecheck                  string `json:"copyright_precheck"`
	EnableCloudDraft                   string `json:"enable_cloud_draft"`
	EnableEditPost                     string `json:"enable_edit_post"`
	EnableLocalDraft                   string `json:"enable_local_draft"`
	EnableNewPlaylist                  string `json:"enable_new_playlist"`
	EnableNewPostAPI                   string `json:"enable_new_post_api"`
	EnableWeb60_MinFlag                bool   `json:"enable_web_60_min_flag"`
	HashtagHistory                     int64  `json:"hashtag_history"`
	InteractionOptimization            int64  `json:"interaction_optimization"`
	MaxDividedNum                      string `json:"max_divided_num"`
	OptV1NewEntranceV3                 int64  `json:"opt_v1_new_entrance_v3"`
	OptV1NewUploadUI                   int64  `json:"opt_v1_new_upload_ui"`
	PostPollingVersion                 int64  `json:"post_polling_version"`
	ShowAigcToggle                     int64  `json:"show_aigc_toggle"`
	ShowCreatorAcademyPCEntrance       int64  `json:"show_creator_academy_pc_entrance"`
	TTSProductAnchor                   int64  `json:"tts_product_anchor"`
	UnlimitedMusicLibrary              string `json:"unlimited_music_library"`
	UploadFileStepOptimization         int64  `json:"upload_file_step_optimization"`
	UseVideoPreviewTranscodeMultiChunk bool   `json:"use_video_preview_transcode_multi_chunk"`
	VideoSplitCountLimit               int64  `json:"video_split_count_limit"`
	WebCreationCoverTool               int64  `json:"web_creation_cover_tool"`
	WebCreationLocalVideoPreview       int64  `json:"web_creation_local_video_preview"`
	WebCreationPoi                     string `json:"web_creation_poi"`
	WebCreationSupportEdit             int64  `json:"web_creation_support_edit"`
}

type TtPlayerPreload struct {
	Vid TtPlayerPreloadVid `json:"vid"`
}

type TtPlayerPreloadVid struct {
	MaxQueueCount          int64  `json:"maxQueueCount"`
	MinBufferLength        int64  `json:"minBufferLength"`
	PreloadMaxCacheCount   int64  `json:"preloadMaxCacheCount"`
	PreloadSize            int64  `json:"preloadSize"`
	PreloadTime            int64  `json:"preloadTime"`
	SegmentMinDuration     int64  `json:"segmentMinDuration"`
	StartPreloadControl    bool   `json:"startPreloadControl"`
	StartPreloadMinBuffer  int64  `json:"startPreloadMinBuffer"`
	StartPreloadMinPosTime int64  `json:"startPreloadMinPosTime"`
	Vid                    string `json:"vid"`
}

type VvAvgPerDayPortrait struct {
	Vid VvAvgPerDayPortraitVid `json:"vid"`
}

type VvAvgPerDayPortraitVid struct {
	Levels []Level `json:"levels"`
	Vid    string  `json:"vid"`
}

type Level struct {
	Target    int64 `json:"target"`
	Threshold int64 `json:"threshold"`
}

type Webcast struct {
	EnableEcLcc                 bool              `json:"enable_ec_lcc"`
	LikeIconOptimize            bool              `json:"like_icon_optimize"`
	LiveRestrictedMode          bool              `json:"live_restricted_mode"`
	PiClipEu                    int64             `json:"pi_clip_eu"`
	PiClipRow                   int64             `json:"pi_clip_row"`
	PiClipUs                    int64             `json:"pi_clip_us"`
	WebDrawerShowExplore        bool              `json:"web_drawer_show_explore"`
	WebFollowGuideStrategyGroup int64             `json:"web_follow_guide_strategy_group"`
	WebMinimalPackage           WebMinimalPackage `json:"web_minimal_package"`
}

type WebMinimalPackage struct {
	Baseline               string        `json:"baseline"`
	FilterPackages         []interface{} `json:"filter_packages"`
	ForceRecommendPackages []string      `json:"force_recommend_packages"`
}

type ZTITest struct {
	ConsumerPathList []string `json:"consumer_path_list"`
	Vid              string   `json:"vid"`
}

type WebappBizContext struct {
	CookieConsent              CookieConsent              `json:"cookieConsent"`
	OS                         string                     `json:"os"`
	IsMobile                   bool                       `json:"isMobile"`
	IsAndroid                  bool                       `json:"isAndroid"`
	IsIOS                      bool                       `json:"isIOS"`
	JumpType                   string                     `json:"jumpType"`
	NavList                    []NavList                  `json:"navList"`
	KapLinks                   []KapLink                  `json:"kapLinks"`
	Config                     PurpleConfig               `json:"config"`
	Domains                    Domains                    `json:"domains"`
	DownloadLink               DownloadLink               `json:"downloadLink"`
	DeviceLimitRegisterExpired bool                       `json:"deviceLimitRegisterExpired"`
	Subdivisions               []string                   `json:"subdivisions"`
	Geo                        []string                   `json:"geo"`
	GeoCity                    GeoCity                    `json:"geoCity"`
	IsGoogleBot                bool                       `json:"isGoogleBot"`
	IsBingBot                  bool                       `json:"isBingBot"`
	IsBot                      bool                       `json:"isBot"`
	IsSearchEngineBot          bool                       `json:"isSearchEngineBot"`
	IsTTP                      bool                       `json:"isTTP"`
	DateFmtLocale              DateFmtLocale              `json:"dateFmtLocale"`
	VideoPlayerConfig          VideoPlayerConfig          `json:"videoPlayerConfig"`
	PlaybackNormalizePath      PlaybackNormalizePath      `json:"playbackNormalizePath"`
	BitrateConfig              BitrateConfigElement       `json:"bitrateConfig"`
	SearchVideoForLoggedin     bool                       `json:"searchVideoForLoggedin"`
	StudioDownloadEntrance     StudioDownloadEntrance     `json:"studioDownloadEntrance"`
	LiveSuggestConfig          LiveSuggestConfig          `json:"liveSuggestConfig"`
	LiveAnchorEntrance         LiveAnchorEntrance         `json:"liveAnchorEntrance"`
	LiveStudioEnable           bool                       `json:"liveStudioEnable"`
	XgplayerInitHost           XgplayerInitHost           `json:"xgplayerInitHost"`
	VideoOrder                 WebappBizContextVideoOrder `json:"videoOrder"`
	SearchLiveForLoggedin      bool                       `json:"searchLiveForLoggedin"`
	CanUseQuery                bool                       `json:"canUseQuery"`
	BitrateSelectorConfigs     BitrateSelectorConfigs     `json:"bitrateSelectorConfigs"`
	Idc                        string                     `json:"idc"`
	Vregion                    string                     `json:"vregion"`
	Vgeo                       string                     `json:"vgeo"`
	VideoCoverSettings         VideoCoverSettings         `json:"videoCoverSettings"`
	HevcRobustness             HevcRobustness             `json:"hevcRobustness"`
	APIKeys                    APIKeys                    `json:"apiKeys"`
	SolariaPortrait            SolariaPortrait            `json:"solariaPortrait"`
	ClaConfig                  ClaConfig                  `json:"claConfig"`
	UpliftModelInfo            string                     `json:"upliftModelInfo"`
	InterestList               string                     `json:"interestList"`
}

type APIKeys struct {
	Firebase string `json:"firebase"`
}

type BitrateConfigElement struct {
	BitrateLower                 int64              `json:"bitrateLower"`
	BitrateRange                 []int64            `json:"bitrateRange"`
	BitrateUpper                 int64              `json:"bitrateUpper"`
	Mode                         string             `json:"mode"`
	ParamBf                      float64            `json:"paramBf"`
	ParamBp                      float64            `json:"paramBp"`
	ParamLower                   float64            `json:"paramLower"`
	ParamUpper                   float64            `json:"paramUpper"`
	ParamUpperBl                 float64            `json:"paramUpperBl"`
	ParamVl1                     float64            `json:"paramVl1"`
	ParamVl2                     int64              `json:"paramVl2"`
	ParamVlLower                 float64            `json:"paramVlLower"`
	ParamVlUpper                 float64            `json:"paramVlUpper"`
	SlidingWindowCountThreshold  int64              `json:"slidingWindowCountThreshold"`
	SlidingWindowExtraction      string             `json:"slidingWindowExtraction"`
	SlidingWindowType            string             `json:"slidingWindowType"`
	SlidingWindowWeight          string             `json:"slidingWindowWeight"`
	SlidingWindowWeightThreshold int64              `json:"slidingWindowWeightThreshold"`
	QualityFilter                *SolariaPortrait   `json:"quality_filter,omitempty"`
	WhiteList                    []interface{}      `json:"white_list,omitempty"`
	AutoBitrateParams            *AutoBitrateParams `json:"autoBitrateParams,omitempty"`
	DefaultBitrate               *int64             `json:"defaultBitrate,omitempty"`
}

type AutoBitrateParams struct {
	ParamA     int64   `json:"paramA"`
	ParamB     float64 `json:"paramB"`
	ParamC     float64 `json:"paramC"`
	ParamD     int64   `json:"paramD"`
	MinBitrate int64   `json:"minBitrate"`
}

type SolariaPortrait struct {
}

type BitrateSelectorConfigs struct {
	Configs []BitrateConfigElement `json:"configs"`
}

type ClaConfig struct {
	TranslationLanguageList   []LanguageList    `json:"translationLanguageList"`
	DntLanguageList           []LanguageList    `json:"dntLanguageList"`
	ISOToSubtitleLanguageCode map[string]string `json:"isoToSubtitleLanguageCode"`
	SubtitleToISOLanguageCode map[string]string `json:"subtitleToIsoLanguageCode"`
}

type LanguageList struct {
	Code      string `json:"code"`
	EnName    string `json:"en_name"`
	LocalName string `json:"local_name"`
}

type PurpleConfig struct {
	FeatureFlags           FeatureFlags           `json:"featureFlags"`
	DesktopAppDownloadLink DesktopAppDownloadLink `json:"desktopAppDownloadLink"`
	SignUpOpen             bool                   `json:"signUpOpen"`
	CookieBanner           CookieBanner           `json:"cookieBanner"`
	IsGrayFilter           bool                   `json:"isGrayFilter"`
	NickNameControlDay     string                 `json:"nickNameControlDay"`
	DesktopAppSurveyLink   DesktopAppSurveyLink   `json:"desktopAppSurveyLink"`
	DesktopWebSurveyLink   DesktopWebSurveyLink   `json:"desktopWebSurveyLink"`
	OnDeviceMLConfig       OnDeviceMLConfig       `json:"onDeviceMLConfig"`
	ExploreCategoryList    ExploreCategoryList    `json:"exploreCategoryList"`
}

type CookieBanner struct {
	LoadDynamically             bool     `json:"load_dynamically"`
	DeclineBtnStagedRolloutArea []string `json:"decline_btn_staged_rollout_area"`
	Resource                    Resource `json:"resource"`
	I18N                        I18N     `json:"i18n"`
}

type I18N struct {
	CookieBannerTitle                            string `json:"cookieBannerTitle"`
	CookieBannerTitleNew                         string `json:"cookieBannerTitleNew"`
	CookieBannerSubTitle                         string `json:"cookieBannerSubTitle"`
	CookieBannerSubTitleNew                      string `json:"cookieBannerSubTitleNew"`
	CookieBannerSubTitleV2                       string `json:"cookieBannerSubTitleV2"`
	CookieBannerBtnManage                        string `json:"cookieBannerBtnManage"`
	CookieBannerBtnAccept                        string `json:"cookieBannerBtnAccept"`
	CookieBannerBtnDecline                       string `json:"cookieBannerBtnDecline"`
	CookiesBannerDetails                         string `json:"cookiesBannerDetails"`
	CookiesBannerCookiesPolicy                   string `json:"cookiesBannerCookiesPolicy"`
	CookiesBannerAccept                          string `json:"cookiesBannerAccept"`
	WebDoNotSellSettingsSavedToast               string `json:"webDoNotSellSettingsSavedToast"`
	CookieSettingManageYourCookieTitle           string `json:"cookieSettingManageYourCookieTitle"`
	CookieSettingSave                            string `json:"cookieSettingSave"`
	CookieSettingAnalyticsAndMarketing           string `json:"cookieSettingAnalyticsAndMarketing"`
	CookieSettingNecessary                       string `json:"cookieSettingNecessary"`
	CookieSettingNecessarySubtitle               string `json:"cookieSettingNecessarySubtitle"`
	CookieSettingNecessaryV2                     string `json:"cookieSettingNecessaryV2"`
	CookieSettingNecessarySubtitleV2             string `json:"cookieSettingNecessarySubtitleV2"`
	CookieSettingAnalyticsAndMarketingSubtitle   string `json:"cookieSettingAnalyticsAndMarketingSubtitle"`
	CookieSettingAnalyticsAndMarketingSubtitleV2 string `json:"cookieSettingAnalyticsAndMarketingSubtitleV2"`
	CookieManageTip                              string `json:"cookieManageTip"`
}

type Resource struct {
	Prefix   string   `json:"prefix"`
	Themes   []string `json:"themes"`
	Esm      string   `json:"esm"`
	Nomodule string   `json:"nomodule"`
	Version  string   `json:"version"`
}

type DesktopAppDownloadLink struct {
	MAC string `json:"mac"`
	Win string `json:"win"`
}

type DesktopAppSurveyLink struct {
	Default string `json:"default"`
	Vn      string `json:"vn"`
}

type DesktopWebSurveyLink struct {
	New string `json:"new"`
	Old string `json:"old"`
}

type ExploreCategoryList struct {
	V0 []V0 `json:"v0"`
	V1 []V0 `json:"v1"`
	V2 []V0 `json:"v2"`
}

type V0 struct {
	Text string `json:"text"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type FeatureFlags struct {
	FeatureBar                        bool     `json:"feature_bar"`
	BusinessAccountOpen               bool     `json:"business_account_open"`
	FeatureTt4BAds                    bool     `json:"feature_tt4b_ads"`
	SupportMultilineDesc              bool     `json:"support_multiline_desc"`
	PCVideoPlaylist                   bool     `json:"pc_video_playlist"`
	FeatureMobileUIOptStage2          bool     `json:"feature_mobile_ui_opt_stage2"`
	AddRecipeCard                     bool     `json:"add_recipe_card"`
	CollapseSEOHeader                 bool     `json:"collapse_seo_header"`
	CollapseSEOHeaderMobile           bool     `json:"collapse_seo_header_mobile"`
	SEOEnableNewPoiPage               bool     `json:"seo_enable_new_poi_page"`
	EnablePrivacyCenter               bool     `json:"enable_privacy_center"`
	HashtagViewcount                  bool     `json:"hashtag_viewcount"`
	ShouldShowEffectDetailPage        bool     `json:"should_show_effect_detail_page"`
	KepRemoveDescKeywords             bool     `json:"kep_remove_desc_keywords"`
	HashtagCanonicalURL               bool     `json:"hashtag_canonical_url"`
	MusicCanonicalURL                 bool     `json:"music_canonical_url"`
	UserCanonicalURL                  bool     `json:"user_canonical_url"`
	FindCardRefactor                  bool     `json:"find_card_refactor"`
	ShapeLoggedinDisabled             bool     `json:"shape_loggedin_disabled"`
	KepNewGrid                        bool     `json:"kep_new_grid"`
	FeedbackProjectUSwitch            bool     `json:"feedback_project_u_switch"`
	FeedbackProjectUFAQWhitelist      []string `json:"feedback_project_u_faq_whitelist"`
	FeedbackProjectUSubmitEntrySwitch bool     `json:"feedback_project_u_submit_entry_switch"`
	FeedbackProjectUSubmitPageSwitch  bool     `json:"feedback_project_u_submit_page_switch"`
}

type OnDeviceMLConfig struct {
	CommentsPreload CommentsPreload `json:"commentsPreload"`
	PlayerPreload   PlayerPreload   `json:"playerPreload"`
}

type CommentsPreload struct {
	CommentPreloadHighThreshold CommentPreloadThreshold `json:"commentPreloadHighThreshold"`
	CommentPreloadLowThreshold  CommentPreloadThreshold `json:"commentPreloadLowThreshold"`
}

type CommentPreloadThreshold struct {
	EnablePreload bool      `json:"enable_preload"`
	PreloadMl     PreloadMl `json:"preload_ml"`
}

type PreloadMl struct {
	Scene        string       `json:"scene"`
	Delay        int64        `json:"delay"`
	SkipCount    int64        `json:"skip_count"`
	RunGap       int64        `json:"run_gap"`
	IgnoreCount  int64        `json:"ignore_count"`
	Package      string       `json:"package"`
	Features     string       `json:"features"`
	Output       []Output     `json:"output"`
	EngineConfig EngineConfig `json:"engine_config"`
}

type EngineConfig struct {
	Inputs  []string `json:"inputs"`
	Outputs []string `json:"outputs"`
}

type Output struct {
	Op     string    `json:"op"`
	Args   []float64 `json:"args"`
	Labels []string  `json:"labels"`
}

type PlayerPreload struct {
	PlayerPreloadStrategy PlayerPreloadStrategy `json:"playerPreloadStrategy"`
}

type PlayerPreloadStrategy struct {
	EnablePreload bool      `json:"enable_preload"`
	PreloadMl     PreloadMl `json:"preload_ml"`
	SlideStrategy []int64   `json:"slideStrategy"`
}

type CookieConsent struct {
	Ga  bool `json:"ga"`
	AF  bool `json:"af"`
	Fbp bool `json:"fbp"`
	Lip bool `json:"lip"`
}

type DateFmtLocale struct {
	Name          string   `json:"name"`
	Months        []string `json:"months"`
	MonthsShort   []string `json:"monthsShort"`
	Weekdays      []string `json:"weekdays"`
	WeekdaysMin   []string `json:"weekdaysMin"`
	WeekdaysShort []string `json:"weekdaysShort"`
	Formats       Formats  `json:"formats"`
	Past          Abbr     `json:"past"`
	Future        Abbr     `json:"future"`
	Abbr          Abbr     `json:"abbr"`
	JustNow       string   `json:"justNow"`
	Yesterday     string   `json:"yesterday"`
	Today         string   `json:"today"`
	Tomorrow      string   `json:"tomorrow"`
	WeekStart     int64    `json:"weekStart"`
}

type Abbr struct {
	Y     M `json:"y"`
	M     M `json:"M"`
	W     M `json:"w"`
	D     M `json:"d"`
	H     M `json:"h"`
	AbbrM M `json:"m"`
	S     M `json:"s"`
}

type M struct {
	One   string `json:"one"`
	Other string `json:"other"`
}

type Formats struct {
	FormatsLt    string `json:"lt"`
	FormatsLts   string `json:"lts"`
	Lt           string `json:"LT"`
	Lts          string `json:"LTS"`
	L            string `json:"L"`
	Ll           string `json:"LL"`
	LlD          string `json:"LL-D"`
	LlY          string `json:"LL-Y"`
	Lll          string `json:"LLL"`
	LllY         string `json:"LLL-Y"`
	Llll         string `json:"LLLL"`
	LlllY        string `json:"LLLL-Y"`
	FormatsL     string `json:"l"`
	LY           string `json:"l-Y"`
	FormatsLl    string `json:"ll"`
	FormatsLlY   string `json:"ll-Y"`
	FormatsLll   string `json:"lll"`
	FormatsLllY  string `json:"lll-Y"`
	FormatsLlll  string `json:"llll"`
	FormatsLlllY string `json:"llll-Y"`
	LlYW         string `json:"ll-Y+w"`
	LTTo         string `json:"LT+to"`
	LLLTo        string `json:"LLL+to"`
}

type Domains struct {
	Kind           string `json:"kind"`
	CAPTCHA        string `json:"captcha"`
	IMAPI          string `json:"imApi"`
	IMFrontier     string `json:"imFrontier"`
	MTAPI          string `json:"mTApi"`
	RootAPI        string `json:"rootApi"`
	SECSDK         string `json:"secSDK"`
	Slardar        string `json:"slardar"`
	Starling       string `json:"starling"`
	Tea            string `json:"tea"`
	TeaChannel     string `json:"teaChannel"`
	TeaChannelType string `json:"teaChannelType"`
	LibraWebSDK    string `json:"libraWebSDK"`
	WebcastAPI     string `json:"webcastApi"`
	WebcastRootAPI string `json:"webcastRootApi"`
	Tcc            string `json:"tcc"`
	LocationAPI    string `json:"locationApi"`
}

type DownloadLink struct {
	Microsoft Amazon `json:"microsoft"`
	Apple     Amazon `json:"apple"`
	Amazon    Amazon `json:"amazon"`
	Google    Amazon `json:"google"`
}

type Amazon struct {
	Visible bool   `json:"visible"`
	Normal  string `json:"normal"`
}

type GeoCity struct {
	City                 string                `json:"City"`
	Subdivisions         string                `json:"Subdivisions"`
	OriginalSubdivisions []OriginalSubdivision `json:"OriginalSubdivisions"`
	SubdivisionsArr      []string              `json:"SubdivisionsArr"`
}

type OriginalSubdivision struct {
	GeoNameID string `json:"GeoNameID"`
	ASCIName  string `json:"ASCIName"`
	Name      string `json:"Name"`
	LocalID   string `json:"LocalID"`
}

type HevcRobustness struct {
	UseHevcRobustTest bool     `json:"useHevcRobustTest"`
	ForceRobustTest   []string `json:"forceRobustTest"`
}

type KapLink struct {
	Title    string         `json:"title"`
	Children []KapLinkChild `json:"children"`
}

type KapLinkChild struct {
	Lang  []string `json:"lang"`
	Links []Link   `json:"links"`
}

type Link struct {
	Title string `json:"title"`
	Href  string `json:"href"`
}

type LiveAnchorEntrance struct {
	LiveCenter bool `json:"liveCenter"`
	CreatorHub bool `json:"creatorHub"`
	LiveStudio bool `json:"liveStudio"`
}

type LiveSuggestConfig struct {
	IsBlockedArea bool `json:"isBlockedArea"`
	IsRiskArea    bool `json:"isRiskArea"`
}

type NavList struct {
	Title    string         `json:"title"`
	Children []NavListChild `json:"children"`
}

type NavListChild struct {
	Title string  `json:"title"`
	Href  string  `json:"href"`
	Key   *string `json:"key,omitempty"`
}

type PlaybackNormalizePath struct {
	Path []string `json:"path"`
}

type StudioDownloadEntrance struct {
	Regions             []string `json:"regions"`
	UserRegions         []string `json:"userRegions"`
	AllRegions          bool     `json:"allRegions"`
	UserBlockRegions    []string `json:"userBlockRegions"`
	UserBlockGeoNameIDs []string `json:"userBlockGeoNameIDs"`
}

type VideoCoverSettings struct {
	Format       int64  `json:"format"`
	AcceptHeader string `json:"acceptHeader"`
	SsrCount     int64  `json:"_ssrCount"`
}

type WebappBizContextVideoOrder struct {
	VideoOrder []VideoOrderElement `json:"videoOrder"`
}

type VideoOrderElement struct {
	Property string  `json:"property"`
	Detail   []int64 `json:"detail,omitempty"`
	Order    *string `json:"order,omitempty"`
}

type VideoPlayerConfig struct {
	Fallback bool `json:"fallback"`
}

type XgplayerInitHost struct {
	Group1 []string `json:"group1"`
	Group2 []string `json:"group2"`
}

type WebappI18NTranslation struct {
	Webapp map[string]string `json:"Webapp"`
}

type WebappVideoDetail struct {
	ItemInfo   ItemInfo  `json:"itemInfo"`
	ShareMeta  ShareMeta `json:"shareMeta"`
	StatusCode int64     `json:"statusCode"`
	StatusMsg  string    `json:"statusMsg"`
}

type ItemInfo struct {
	ItemStruct ItemStruct `json:"itemStruct"`
}

type ItemStruct struct {
	ID                         string        `json:"id"`
	Desc                       string        `json:"desc"`
	CreateTime                 string        `json:"createTime"`
	ScheduleTime               int64         `json:"scheduleTime"`
	Video                      Video         `json:"video"`
	Author                     Author        `json:"author"`
	Music                      Music         `json:"music"`
	Challenges                 []interface{} `json:"challenges"`
	Stats                      Stats         `json:"stats"`
	StatsV2                    StatsV2       `json:"statsV2"`
	WarnInfo                   []interface{} `json:"warnInfo"`
	OriginalItem               bool          `json:"originalItem"`
	OfficalItem                bool          `json:"officalItem"`
	TextExtra                  []interface{} `json:"textExtra"`
	Secret                     bool          `json:"secret"`
	ForFriend                  bool          `json:"forFriend"`
	Digged                     bool          `json:"digged"`
	ItemCommentStatus          int64         `json:"itemCommentStatus"`
	TakeDown                   int64         `json:"takeDown"`
	EffectStickers             []interface{} `json:"effectStickers"`
	AuthorStats                AuthorStats   `json:"authorStats"`
	PrivateItem                bool          `json:"privateItem"`
	DuetEnabled                bool          `json:"duetEnabled"`
	StitchEnabled              bool          `json:"stitchEnabled"`
	StickersOnItem             []interface{} `json:"stickersOnItem"`
	ShareEnabled               bool          `json:"shareEnabled"`
	Comments                   []interface{} `json:"comments"`
	DuetDisplay                int64         `json:"duetDisplay"`
	StitchDisplay              int64         `json:"stitchDisplay"`
	IndexEnabled               bool          `json:"indexEnabled"`
	DiversificationLabels      []string      `json:"diversificationLabels"`
	LocationCreated            string        `json:"locationCreated"`
	SuggestedWords             []interface{} `json:"suggestedWords"`
	Contents                   []Content     `json:"contents"`
	DiversificationID          int64         `json:"diversificationId"`
	Collected                  bool          `json:"collected"`
	ChannelTags                []interface{} `json:"channelTags"`
	ItemControl                ItemControl   `json:"item_control"`
	IsAigc                     bool          `json:"IsAigc"`
	AIGCDescription            string        `json:"AIGCDescription"`
	BackendSourceEventTracking string        `json:"backendSourceEventTracking"`
	CategoryType               int64         `json:"CategoryType"`
	TextLanguage               string        `json:"textLanguage"`
	TextTranslatable           bool          `json:"textTranslatable"`
}

type Author struct {
	ID                   string `json:"id"`
	ShortID              string `json:"shortId"`
	UniqueID             string `json:"uniqueId"`
	Nickname             string `json:"nickname"`
	AvatarLarger         string `json:"avatarLarger"`
	AvatarMedium         string `json:"avatarMedium"`
	AvatarThumb          string `json:"avatarThumb"`
	Signature            string `json:"signature"`
	CreateTime           int64  `json:"createTime"`
	Verified             bool   `json:"verified"`
	SECUid               string `json:"secUid"`
	Ftc                  bool   `json:"ftc"`
	Relation             int64  `json:"relation"`
	OpenFavorite         bool   `json:"openFavorite"`
	CommentSetting       int64  `json:"commentSetting"`
	DuetSetting          int64  `json:"duetSetting"`
	StitchSetting        int64  `json:"stitchSetting"`
	PrivateAccount       bool   `json:"privateAccount"`
	Secret               bool   `json:"secret"`
	IsADVirtual          bool   `json:"isADVirtual"`
	RoomID               string `json:"roomId"`
	UniqueIDModifyTime   int64  `json:"uniqueIdModifyTime"`
	TtSeller             bool   `json:"ttSeller"`
	DownloadSetting      int64  `json:"downloadSetting"`
	RecommendReason      string `json:"recommendReason"`
	NowInvitationCardURL string `json:"nowInvitationCardUrl"`
	NickNameModifyTime   int64  `json:"nickNameModifyTime"`
	IsEmbedBanned        bool   `json:"isEmbedBanned"`
	CanExpPlaylist       bool   `json:"canExpPlaylist"`
	SuggestAccountBind   bool   `json:"suggestAccountBind"`
}

type AuthorStats struct {
	FollowerCount  int64 `json:"followerCount"`
	FollowingCount int64 `json:"followingCount"`
	Heart          int64 `json:"heart"`
	HeartCount     int64 `json:"heartCount"`
	VideoCount     int64 `json:"videoCount"`
	DiggCount      int64 `json:"diggCount"`
	FriendCount    int64 `json:"friendCount"`
}

type Content struct {
	Desc      string        `json:"desc"`
	TextExtra []interface{} `json:"textExtra"`
}

type ItemControl struct {
	CanRepost bool `json:"can_repost"`
}

type Music struct {
	ID                 string          `json:"id"`
	Title              string          `json:"title"`
	PlayURL            string          `json:"playUrl"`
	CoverLarge         string          `json:"coverLarge"`
	CoverMedium        string          `json:"coverMedium"`
	CoverThumb         string          `json:"coverThumb"`
	AuthorName         string          `json:"authorName"`
	Original           bool            `json:"original"`
	Duration           int64           `json:"duration"`
	ScheduleSearchTime int64           `json:"scheduleSearchTime"`
	Collected          bool            `json:"collected"`
	PreciseDuration    PreciseDuration `json:"preciseDuration"`
	IsCopyrighted      bool            `json:"isCopyrighted"`
}

type PreciseDuration struct {
	PreciseDuration         float64 `json:"preciseDuration"`
	PreciseShootDuration    float64 `json:"preciseShootDuration"`
	PreciseAuditionDuration float64 `json:"preciseAuditionDuration"`
	PreciseVideoDuration    float64 `json:"preciseVideoDuration"`
}

type Stats struct {
	DiggCount    int64  `json:"diggCount"`
	ShareCount   int64  `json:"shareCount"`
	CommentCount int64  `json:"commentCount"`
	PlayCount    int64  `json:"playCount"`
	CollectCount string `json:"collectCount"`
}

type StatsV2 struct {
	DiggCount    string `json:"diggCount"`
	ShareCount   string `json:"shareCount"`
	CommentCount string `json:"commentCount"`
	PlayCount    string `json:"playCount"`
	CollectCount string `json:"collectCount"`
	RepostCount  string `json:"repostCount"`
}

type Video struct {
	ID            string            `json:"id"`
	Height        int64             `json:"height"`
	Width         int64             `json:"width"`
	Duration      int64             `json:"duration"`
	Ratio         string            `json:"ratio"`
	Cover         string            `json:"cover"`
	OriginCover   string            `json:"originCover"`
	DynamicCover  string            `json:"dynamicCover"`
	PlayAddr      string            `json:"playAddr"`
	DownloadAddr  string            `json:"downloadAddr"`
	ShareCover    []string          `json:"shareCover"`
	ReflowCover   string            `json:"reflowCover"`
	Bitrate       int64             `json:"bitrate"`
	EncodedType   string            `json:"encodedType"`
	Format        string            `json:"format"`
	VideoQuality  string            `json:"videoQuality"`
	EncodeUserTag string            `json:"encodeUserTag"`
	CodecType     string            `json:"codecType"`
	Definition    string            `json:"definition"`
	SubtitleInfos []SubtitleInfo    `json:"subtitleInfos"`
	ZoomCover     map[string]string `json:"zoomCover"`
	VolumeInfo    VolumeInfo        `json:"volumeInfo"`
	BitrateInfo   []BitrateInfo     `json:"bitrateInfo"`
	VQScore       string            `json:"VQScore"`
	ClaInfo       ClaInfo           `json:"claInfo"`
}

type BitrateInfo struct {
	Bitrate     int64    `json:"Bitrate"`
	QualityType int64    `json:"QualityType"`
	GearName    string   `json:"GearName"`
	PlayAddr    PlayAddr `json:"PlayAddr"`
	CodecType   string   `json:"CodecType"`
	Mvmaf       string   `json:"MVMAF"`
}

type PlayAddr struct {
	DataSize string   `json:"DataSize"`
	Width    int64    `json:"Width"`
	Height   int64    `json:"Height"`
	URI      string   `json:"Uri"`
	URLList  []string `json:"UrlList"`
	URLKey   string   `json:"UrlKey"`
	FileHash string   `json:"FileHash"`
	FileCS   string   `json:"FileCs"`
}

type ClaInfo struct {
	HasOriginalAudio  bool          `json:"hasOriginalAudio"`
	EnableAutoCaption bool          `json:"enableAutoCaption"`
	CaptionInfos      []interface{} `json:"captionInfos"`
	NoCaptionReason   int64         `json:"noCaptionReason"`
}

type SubtitleInfo struct {
	URLExpire        string `json:"UrlExpire"`
	Size             string `json:"Size"`
	LanguageID       string `json:"LanguageID"`
	LanguageCodeName string `json:"LanguageCodeName"`
	URL              string `json:"Url"`
	Format           Format `json:"Format"`
	Version          string `json:"Version"`
	Source           Source `json:"Source"`
}

type VolumeInfo struct {
	Loudness float64 `json:"Loudness"`
	Peak     float64 `json:"Peak"`
}

type ShareMeta struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

type Format string

const (
	Webvtt Format = "webvtt"
)

type Source string

const (
	ASR Source = "ASR"
	MT  Source = "MT"
)
