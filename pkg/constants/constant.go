package constants

const (
	//表名
	CommentTableName = "comment"
	FollowTableName  = "follow"
	LikeTableName    = "like"
	UserTableName    = "user"
	VideoTableName   = "video"
	//jwt插件变量
	SecretKey   = "secret key"
	IdentityKey = "user_id"
	//gorm查询时用到的参数(api接口返回的字段名）
	//Total  = "total"
	//Notes  = "notes"
	//NoteID = "note_id"

	//服务名
	ApiServiceName        = "api"
	UserActionServiceName = "useraction"
	UserBaseServiceName   = "userbase"
	VideoBaseServiceName  = "video"

	//mysql默认dsn
	/*
		Data Source Name (DSN)的PDO命名惯例为：PDO驱动程序的名称，
		后面为一个冒号，再后面是可选的驱动程序连接数据库变量信息，
		如主机名、端口和数据库名。
	*/
	MySQLDefaultDSN = "root:Douyin2022@tcp(localhost:3306)/douyin?charset=utf8&parseTime=True&loc=Local"
	//etcd地址
	EtcdAddress = "127.0.0.1:2379"
	//cpu限制
	CPURateLimit float64 = 80.0
	//默认分页
	DefaultLimit = 10
)
