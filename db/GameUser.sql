CREATE DATABASE AccountDB;
USE AccountDB;
CREATE TABLE AccountsFace (
	ID INT NOT NULL AUTO_INCREMENT,
	UserID INT NOT NULL COMMENT '用户标识',
	CustomFace NVARCHAR(100) NOT NULL,
	InsertTime DATETIME NOT NULL DEFAULT NOW(),
	InsertAddr NVARCHAR(15) NOT NULL COMMENT '登录地址',
	InserMachine NVARCHAR(12) NOT NULL,	
	PRIMARY KEY (ID)	
);



CREATE TABLE AccountsInfo(
	UserID INT NOT NULL AUTO_INCREMENT COMMENT '用户标识',
	Accounts NVARCHAR(16) NOT NULL COMMENT '用户账号(手机号码)',
	NickName NVARCHAR(31) NOT NULL COMMENT '用户昵称',
	SpreaderID NVARCHAR(16) NOT NULL DEFAULT "0" COMMENT '推广员标识',
	UnderWrite NVARCHAR(63) NOT NULL DEFAULT ""  COMMENT '个性签名',
	Compellation NVARCHAR(16) NOT NULL DEFAULT "N''"  COMMENT '真实姓名',
	LoginPass NCHAR(32) NOT NULL COMMENT '登录密码',
	FaceID SMALLINT NOT NULL DEFAULT 0  COMMENT '头像标识',
	CustomID INT NOT NULL DEFAULT 0  COMMENT '自定标识',
	IsOnline BIT NOT NULL DEFAULT 0  COMMENT '在线标识',
	CustomFaceVer TINYINT NOT NULL DEFAULT 0  COMMENT '头像版本',
	Gender TINYINT NOT NULL DEFAULT 0  COMMENT '用户性别',
	Nullity TINYINT NOT NULL DEFAULT 0  COMMENT '禁止服务',
	NullityOverDate DATETIME NOT NULL DEFAULT "1900-01-01"  COMMENT '禁止时间',
	StunDown TINYINT NOT NULL DEFAULT 0  COMMENT '关闭标识',
	MoorMachine TINYINT NOT NULL DEFAULT 0  COMMENT '固定机器', 
	WebLoginTimes INT NOT NULL DEFAULT 0  COMMENT '登录次数',
	GameLoginTimes INT NOT NULL DEFAULT 0  COMMENT '登录次数',
	PlayTimeCount INT NOT NULL DEFAULT 0  COMMENT '游戏时间',
	OnLineTimeCount INT NOT NULL DEFAULT 0  COMMENT '在线时间',
	LastLoginIP NVARCHAR(15) NOT NULL  COMMENT '登录地址',
	LastLoginDate DATETIME NOT NULL DEFAULT NOW() COMMENT '登录时间',
	LastLoginMacine NVARCHAR(32) NOT NULL DEFAULT "-----------"  COMMENT '登录机器',
	RegisterIP NVARCHAR(15) NOT NULL COMMENT '注册地址',
	RegisterDate DATETIME NOT NULL DEFAULT NOW()  COMMENT '注册时间',
	RegisterMachine NVARCHAR(32) NOT NULL DEFAULT "-----------"  COMMENT '注册机器',
	PRIMARY KEY (UserID)
);

CREATE TABLE ConfineAddress(
	AddrString NVARCHAR(15) NOT NULL  COMMENT '地址字符',
	EnjoinLogin BIT NOT NULL DEFAULT 0  COMMENT '限制登录',
	EnjoinRegister BIT NOT NULL DEFAULT 0  COMMENT '限制注册',
	EnjoinOverDate DATETIME  COMMENT '过期时间',
	CollectDate DATETIME NOT NULL DEFAULT NOW()  COMMENT '收集时间',
	CollectNote NVARCHAR(32) NOT NULL DEFAULT ""  COMMENT '输入备注',
	PRIMARY KEY (AddrString)
);
CREATE TABLE ConfineMachine(
	MachineSerial NVARCHAR(32) NOT NULL COMMENT '机器序列',
	EnjoinLogin BIT NOT NULL  ,
	EnjoinRegister BIT NOT NULL DEFAULT 0  COMMENT '限制注册',
	EnjoinOverDate DATETIME COMMENT '过期时间',
	CollectDate DATETIME NOT NULL DEFAULT NOW()  COMMENT '收集时间',
	CollectNote NVARCHAR(32) NOT NULL DEFAULT ""  COMMENT '输入备注',
	PRIMARY KEY (MachineSerial)
);

CREATE TABLE GameIdentifier(
	UserID INT NOT NULL COMMENT '用户标识',
	Accounts NVARCHAR(16) NOT NULL  COMMENT '游戏标识',
	IDLevel INT NOT NULL DEFAULT 0 COMMENT '标识等级',
	PRIMARY KEY (UserID)
);

CREATE TABLE IndividualDatum(
	UserID INT NOT NULL,
	Compellation NVARCHAR(16) NOT NULL DEFAULT "" COMMENT '真实姓名',
	QQ NVARCHAR(16) NOT NULL COMMENT 'QQ号码',
	Email NVARCHAR(32) NOT NULL COMMENT '电子邮件',
	MobilePhone NVARCHAR(16) NOT NULL DEFAULT "" COMMENT '手机号码',
	CollectDate DATETIME NOT NULL DEFAULT NOW() COMMENT '收集时间',
	UserNote NVARCHAR(256) NOT NULL DEFAULT "" COMMENT '用户备注',
	PRIMARY KEY(UserID)

);


CREATE TABLE SystemGrantCount(
	DateID INT NOT NULL,
	RegisterIP NVARCHAR(15) NOT NULL COMMENT '注册地址',
	RegisterMachine NVARCHAR(32) NOT NULL DEFAULT "---------" COMMENT '注册地址',
	GrantScore BIGINT NOT NULL COMMENT '赠送金币',
	GrantCount BIGINT NOT NULL DEFAULT 0 COMMENT '赠送次数',
	CollectDate DATETIME NOT NULL DEFAULT NOW() COMMENT '收集时间',
	PRIMARY KEY(DateID)
);


CREATE TABLE SystemStatusInfo(
	StatusName NVARCHAR(32) NOT NULL COMMENT '状态名字',
	StatusValue INT NOT NULL DEFAULT 0 COMMENT '状态数值',
	StatusString NVARCHAR(512) NOT NULL DEFAULT "" COMMENT '状态字符',
	StatusTip NVARCHAR(50) NOT NULL DEFAULT "" COMMENT '状态显示名称',
	StatusDescription NVARCHAR(100) NOT NULL DEFAULT "" COMMENT '字符描述',
	PRIMARY KEY(StatusName)
);

CREATE TABLE SystemStreamInfo(
	DateID INT NOT NULL COMMENT '日期标识',
	WebLoginSuccess INT NOT NULL DEFAULT 0 COMMENT '登录成功',
	WebRegisterSuccess INT NOT NULL DEFAULT 0 COMMENT '注册成功',
	GameLoginSuccess INT NOT NULL DEFAULT 0 COMMENT '登录成功',
	GameRegisterSuccess DATETIME NOT NULL DEFAULT 0 COMMENT '注册成功',
	CollectDate DATETIME NOT NULL DEFAULT NOW() COMMENT '登录时间',
	PRIMARY KEY (DateID)
);




