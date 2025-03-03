/**
 * Copyright 2014 @ 56x.net.
 * name :
 * author : jarryliu
 * date : 2013-12-09 09:49
 * description :
 * history :
 */

package member

const (
	// DefaultRelateUser 默认操作用户
	DefaultRelateUser int64 = 0
)
const (
	StateStopped = 0 //已停用
	StateOk      = 1 //正常
	BankNoLock   = 0
	BankLocked   = 1
)

const (
	// FavTypeShop 收藏店铺
	FavTypeShop = iota + 1
	// FavTypeGoods 收藏商品
	FavTypeGoods
)

const (
	// PremiumNormal 普通会员
	PremiumNormal int = 0
	// PremiumGold 金会员
	PremiumGold int = 1
	// PremiumWhiteGold 白金会员
	PremiumWhiteGold int = 2
	// PremiumSuper 黑钻会员
	PremiumSuper int = 3
)

const (
	// LAutoUpgrade 自动升级
	LAutoUpgrade = 1
	// LServiceAgentUpgrade 客服更改
	LServiceAgentUpgrade = 2
	// LProgramUpgrade 程序升级，todo: 暂时未对其进行区分
	LProgramUpgrade = 3
)

const (
	// FlagActive 已激活(1)
	FlagActive = 1 << iota
	// FlagLocked 已锁定的(2)
	FlagLocked
	// FlagTrusted 已认证的(4)
	FlagTrusted
	// FlagNoTradePasswd 交易密码未设置(8)
	FlagNoTradePasswd = 64
	// FlagSeller 商户(64)
	FlagSeller = 128
	// FlagProfileCompleted 已完善的资料(16)
	FlagProfileCompleted = 256
)

type (
	IMember interface {
		// GetAggregateRootId 获取聚合根编号
		GetAggregateRootId() int64
		// Complex 会员汇总信息
		Complex() *ComplexMember
		// Profile 会员资料服务
		Profile() IProfileManager
		// Favorite 会员收藏服务
		Favorite() IFavoriteManager
		// GiftCard 礼品卡服务
		GiftCard() IGiftCardManager
		// Invitation 邀请管理
		Invitation() IInvitationManager
		// GetValue 获取值
		GetValue() Member
		// GetAccount 获取账户
		GetAccount() IAccount
		// SendCheckCode 发送验证码,传入操作及消息类型,并返回验证码,及错误
		SendCheckCode(operation string, mssType int) (string, error)
		// CompareCode 对比验证码
		CompareCode(code string) error
		// Active 激活
		Active() error
		// Lock 锁定会员,如minutes为-1, 则永久锁定
		Lock(minutes int, remark string) error
		// Unlock 解锁会员
		Unlock() error
		// ContainFlag 判断是否包含标志
		ContainFlag(f int) bool
		// GetRelation 获取关联的会员
		GetRelation() *InviteRelation
		// BindInviter 绑定邀请人,如果已邀请,force为true时更新
		BindInviter(memberId int64, force bool) error
		// AddExp 增加经验值
		AddExp(exp int) error
		// Premium 升级为高级会员
		Premium(v int, expires int64) error
		// GetLevel 获取等级
		GetLevel() *Level

		// GrantFlag 标志赋值, 如果flag小于零, 则异或运算(去除)
		GrantFlag(flag int) error

		// TestFlag 是否包含标志
		TestFlag(flag int) bool

		// ChangeLevel 更改会员等级,@paymentId:支付单编号,@review:是否需要审核
		ChangeLevel(level int, paymentId int, review bool) error

		// ReviewLevelUp 审核升级请求
		ReviewLevelUp(id int, pass bool) error

		// ConfirmLevelUp 标记已经处理升级
		ConfirmLevelUp(id int) error

		// ChangeUser 更换用户名
		ChangeUser(string) error

		// UpdateLoginTime 更新登录时间
		UpdateLoginTime() error

		// Save 保存
		Save() (int64, error)
	}

	// IProfileManager 会员资料服务
	IProfileManager interface {
		// GetProfile 获取资料
		GetProfile() Profile
		// SaveProfile 保存资料
		SaveProfile(v *Profile) error
		// ChangePhone 更改手机号码,不验证手机格式
		ChangePhone(string) error
		// ChangeAvatar 设置头像
		ChangeAvatar(string) error
		// ProfileCompleted 资料是否完善
		ProfileCompleted() bool
		// CheckProfileComplete 检查资料是否完善
		CheckProfileComplete() error
		// ModifyPassword 修改密码,旧密码可为空; 传入原始密码。密码均为密文
		ModifyPassword(newPwd, oldPwd string) error
		// ModifyTradePassword 修改交易密码，旧密码可为空; 传入原始密码。密码均为密文
		ModifyTradePassword(newPassword, oldPwd string) error
		// GetBankCards 获取提现银行信息
		GetBankCards() []BankCard
		// GetBankCard 获取绑定的银行卡
		GetBankCard(cardNo string) *BankCard
		// AddBankCard 添加银行卡
		AddBankCard(*BankCard) error
		// RemoveBankCard 移除银行卡
		RemoveBankCard(cardNo string) error
		// ReceiptsCodes 获取收款码
		ReceiptsCodes() []ReceiptsCode
		// SaveReceiptsCode 保存收款码
		SaveReceiptsCode(c *ReceiptsCode) error
		// GetTrustedInfo 实名认证信息
		GetTrustedInfo() TrustedInfo
		// SaveTrustedInfo 保存实名认证信息
		SaveTrustedInfo(v *TrustedInfo) error
		// ReviewTrustedInfo 审核实名认证,若重复审核将返回错误
		ReviewTrustedInfo(pass bool, remark string) error
		// CreateDeliver 创建配送地址
		CreateDeliver(*ConsigneeAddress) IDeliverAddress
		// GetDeliverAddress 获取配送地址
		GetDeliverAddress() []IDeliverAddress
		// GetAddress 获取配送地址
		GetAddress(addressId int64) IDeliverAddress
		// SetDefaultAddress 设置默认地址
		SetDefaultAddress(addressId int64) error
		// GetDefaultAddress 获取默认收货地址
		GetDefaultAddress() IDeliverAddress
		// DeleteAddress 删除配送地址
		DeleteAddress(addressId int64) error
	}

	// IFavoriteManager 收藏服务
	IFavoriteManager interface {
		// Favorite 收藏
		Favorite(favType int, referId int64) error
		// Favored 是否已收藏
		Favored(favType int, referId int64) bool
		// Cancel 取消收藏
		Cancel(favType int, referId int64) error
	}

	// ComplexMember 会员概览信息
	ComplexMember struct {
		// 昵称
		Name string
		// 真实姓名
		RealName string
		// 头像
		Avatar string
		// 手机号码
		Phone string
		// 经验值
		Exp int
		// 等级
		Level int
		// 等级名称
		LevelName string
		// 邀请码
		InviteCode string
		// 实名认证状态
		TrustAuthState int
		// 交易密码是否已设置
		TradePasswordHasSet bool
		// 高级会员类型
		PremiumUser int
		// 会员标志
		Flag int
		// 更新时间
		UpdateTime int64
	}

	Member struct {
		// 编号
		Id int64 `db:"id" auto:"yes" pk:"yes"`
		// 用户编码
		Code string `db:"code"`
		// 昵称
		Name string `db:"name"`
		// 真实姓名
		RealName string `db:"real_name"`
		// 用户名
		User string `db:"user"`
		// 加密盐
		Salt string `db:"salt"`
		// 密码
		Pwd string `db:"pwd"`
		// 头像
		Avatar string `db:"avatar"`
		// 交易密码
		TradePassword string `db:"trade_pwd"`
		// 经验值
		Exp int `db:"exp"`
		// 等级
		Level int `db:"level"`
		// 邀请码
		InviteCode string `db:"invite_code"`
		// 高级用户类型
		PremiumUser int `db:"premium_user"`
		// 高级用户过期时间
		PremiumExpires int64 `db:"premium_expires"`
		// 手机号码
		Phone string `db:"phone"`
		// 电子邮箱
		Email string `db:"email"`
		// 注册来源
		RegFrom string `db:"reg_from"`
		// 注册IP
		RegIp string `db:"reg_ip"`
		// 注册时间
		RegTime int64 `db:"reg_time"`
		// 校验码
		CheckCode string `db:"check_code"`
		// 校验码过期时间
		CheckExpires int64 `db:"check_expires"`
		// 会员标志
		Flag int `db:"flag"`
		// 状态
		State int `db:"state"`
		// 登录时间
		LoginTime int64 `db:"login_time"`
		// 最后登录时间
		LastLoginTime int64 `db:"last_login_time"`
		// 更新时间
		UpdateTime int64 `db:"update_time"`
		// 超时时间
		TimeoutTime int64 `db:"-"`
	}

	// 会员资料
	Profile struct {
		//会员编号
		MemberId int64 `db:"member_id" pk:"yes" auto:"no"`
		//昵称
		Name string `db:"name"`
		//头像
		Avatar string `db:"avatar"`
		//性别
		Gender int32 `db:"gender"`
		//生日
		BirthDay string `db:"birthday"`
		//电话
		Phone string `db:"phone"`
		//地址
		Address string `db:"address"`
		//即时通讯
		Im string `db:"im"`
		//电子邮件
		Email string `db:"email"`
		// 省
		Province int32 `db:"province"`
		// 市
		City int32 `db:"city"`
		// 区
		District int32 `db:"district"`
		//备注
		Remark string `db:"remark"`

		// 扩展1
		Ext1 string `db:"ext_1"`
		// 扩展2
		Ext2 string `db:"ext_2"`
		// 扩展3
		Ext3 string `db:"ext_3"`
		// 扩展4
		Ext4 string `db:"ext_4"`
		// 扩展5
		Ext5 string `db:"ext_5"`
		// 扩展6
		Ext6 string `db:"ext_6"`
		//更新时间
		UpdateTime int64 `db:"update_time"`
	}

	// 会员邀请关联表
	InviteRelation struct {
		// 会员编号
		MemberId int64 `db:"member_id" pk:"yes"`
		// 会员卡号
		CardCard string `db:"card_no"`
		// 邀请人（会员）
		InviterId int64 `db:"inviter_id"`
		// 邀请会员编号(depth2)
		InviterD2 int64 `db:"inviter_d2"`
		// 邀请会员编号(depth3)
		InviterD3 int64 `db:"inviter_d3"`
		// 注册关联商户编号
		RegMchId int64 `db:"reg_mchid"`
	}

	// 实名认证信息
	TrustedInfo struct {
		// 会员编号
		MemberId int64 `db:"member_id" pk:"yes"`
		// 真实姓名
		RealName string `db:"real_name"`
		// 国家代码
		CountryCode string `db:"country_code"`
		// 证件类型
		CardType int `db:"card_type"`
		// 证件号码
		CardId string `db:"card_id"`
		// 证件图片
		CardImage string `db:"card_image"`
		// 证件反面图片
		CardReverseImage string `db:"card_reverse_image"`
		// 认证图片,人与身份证的图像等
		TrustImage string `db:"trust_image"`
		// 是否人工审核认证
		ManualReview int `db:"manual_review"`
		// 是否审核通过
		ReviewState int `db:"review_state"`
		// 审核时间
		ReviewTime int64 `db:"review_time"`
		// 审核备注
		Remark string `db:"remark"`
		// 更新时间
		UpdateTime int64 `db:"update_time"`
	}

	// 银行卡信息,因为重要且非频繁更新的数据
	// 所以需要用IsLocked来标记是否锁定
	BankInfo_ struct {
		//会员编号
		MemberId int64 `db:"member_id" pk:"yes"`
		//名称
		BankName string `db:"name"`
		//账号
		Account string `db:"account"`
		//账户名
		AccountName string `db:"account_name"`
		//支行网点
		Network string `db:"network"`
		//状态
		State int `db:"state"`
		//是否锁定
		IsLocked int `db:"is_locked"`
		//更新时间
		UpdateTime int64 `db:"update_time"`
	}

	// 银行卡
	BankCard struct {
		// 编号
		Id int64 `db:"id" pk:"yes" auto:"yes"`
		// 会员编号
		MemberId int64 `db:"member_id"`
		// 银行账号
		BankAccount string `db:"bank_account"`
		// 户名
		AccountName string `db:"account_name"`
		// 银行编号
		BankId int `db:"bank_id"`
		// 银行名称
		BankName string `db:"bank_name"`
		// 银行卡代码
		BankCode string `db:"bank_code"`
		// 网点
		Network string `db:"network"`
		// 快捷支付授权码
		AuthCode string `db:"auth_code"`
		// 状态
		State int16 `db:"state"`
		// 添加时间
		CreateTime int64 `db:"create_time"`
	}

	// 收款码
	ReceiptsCode struct {
		// 编号
		Id int `db:"id" pk:"yes" auto:"yes"`
		// 会员编号
		MemberId int64 `db:"member_id"`
		// 账户标识,如:alipay
		Identity string `db:"identity"`
		// 账户名称
		Name string `db:"name"`
		// 账号
		AccountId string `db:"account_id"`
		// 收款码地址
		CodeUrl string `db:"code_url"`
		// 是否启用
		State int `db:"state"`
	}
	// 收藏
	Favorite struct {
		// 编号
		Id int32 `db:"id" pk:"yes" auto:"yes"`
		// 会员编号
		MemberId int64 `db:"member_id"`
		// 收藏类型
		FavType int `db:"fav_type"`
		// 引用编号
		ReferId int64 `db:"refer_id"`
		// 收藏时间
		CreateTime int64 `db:"create_time"`
	}

	// 收货地址
	IDeliverAddress interface {
		GetDomainId() int64
		GetValue() ConsigneeAddress
		SetValue(*ConsigneeAddress) error
		Save() (int64, error)
	}

	// 收货地址
	ConsigneeAddress struct {
		//编号
		Id int64 `db:"id" pk:"yes" auto:"yes"`
		//会员编号
		MemberId int64 `db:"member_id"`
		//收货人
		ConsigneeName string `db:"consignee_name"`
		//电话
		ConsigneePhone string `db:"consignee_phone"`
		//省
		Province int32 `db:"province"`
		//市
		City int32 `db:"city"`
		//区
		District int32 `db:"district"`
		//地区(省市区连接文本)
		Area string `db:"area"`
		//地址
		DetailAddress string `db:"detail_address"`
		//是否默认
		IsDefault int `db:"is_default"`
	}

	// 会员升级日志
	LevelUpLog struct {
		Id int `db:"id" pk:"yes" auto:"yes"`
		// 会员编号
		MemberId int64 `db:"member_id"`
		// 原来等级
		OriginLevel int `db:"origin_level"`
		// 现在等级
		TargetLevel int `db:"target_level"`
		// 是否为免费升级的会员
		IsFree int `db:"is_free"`
		// 支付单编号
		PaymentId int `db:"payment_id"`
		// 是否审核及处理
		ReviewState int `db:"review_state"`
		// 升级方式,1:自动升级 2:客服更改 3:系统升级
		UpgradeMode int `db:"upgrade_mode"`
		// 升级时间
		CreateTime int64 `db:"create_time"`
	}

	// MmLockInfo 会员锁定记录
	MmLockInfo struct {
		// 编号
		Id int `db:"id" pk:"yes" auto:"yes"`
		// 会员编号
		MemberId int64 `db:"member_id"`
		// 锁定时间
		LockTime int64 `db:"lock_time"`
		// 解锁时间
		UnlockTime int64 `db:"unlock_time"`
		// 备注
		Remark string `db:"remark"`
	}

	// MmLockHistory 会员锁定历史
	MmLockHistory struct {
		// 编号
		Id int `db:"id" pk:"yes" auto:"yes"`
		// 会员编号
		MemberId int64 `db:"member_id"`
		// 锁定时间
		LockTime int64 `db:"lock_time"`
		// 锁定持续分钟数
		Duration int `db:"duration"`
		// 备注
		Remark string `db:"remark"`
	}
)

func (b BankCard) Right() bool {
	return len(b.BankName) > 0 && len(b.BankAccount) > 0 &&
		len(b.AccountName) > 0
}

func (b BankCard) Locked() bool {
	return false
	//panic(errors.New("Not Implemented"))
	//return b.IsLocked == BankLocked
}
