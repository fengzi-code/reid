package model

type YamlSetting struct {
	SysInfo  SysInfo  `yaml:"SysInfo"`
	FileInfo FileInfo `yaml:"FileInfo"`
	Mysql    Mysql    `yaml:"Mysql"`
	Redis    Redis    `yaml:"Redis"`
	WebInfo  WebInfo  `yaml:"WebInfo"`
	Zap      Zap
}

type Datebases struct {
	Database1 string
	Database2 string
}

type SysInfo struct {
	SysName       string `yaml:"SysName"`
	SysHeTongId   string `yaml:"SysHeTongId"`
	PortraitCount int    `yaml:"PortraitCount"`
	Model         int64  `yaml:"Model"`
}

type FileInfo struct {
	YpNameTemp     string `yaml:"YpNameTemp"`
	YpLoginTemp    string `yaml:"YpLoginTemp"`
	AccessTemp     string `yaml:"AccessTemp"`
	PcInfoTemp     string `yaml:"PcInfoTemp"`
	ModStatTemp    string `yaml:"ModStatTemp"`
	YpNameSaveDir  string `yaml:"YpNameSaveDir"`
	YpLoginSaveDir string `yaml:"YpLoginSaveDir"`
	AccessSaveDir  string `yaml:"AccessSaveDir"`
	PcInfoSaveDir  string `yaml:"PcInfoSaveDir"`
	ModStatSaveDir string `yaml:"ModStatSaveDir"`
	UserJson       string `yaml:"UserJson"`
	Logjson        string `yaml:"Logjson"`
	AnjianJson     string `yaml:"AnjianJson"`
	AnjianuserJson string `yaml:"AnjianuserJson"`
	ZipSaveDir     string `yaml:"ZipSaveDir"`
}

type Redis struct {
	Db              int    `yaml:"Db"`
	Password        string `yaml:"Password"`
	Addr            string `yaml:"Addr"`
	Daykey          string `yaml:"Daykey"`
	MonthAlarmKey   string `yaml:"MonthAlarmKey"`
	MonthCaptureKey string `yaml:"MonthCaptureKey"`
}

type Mysql struct {
	Username     string    `yaml:"Username"`
	Password     string    `yaml:"Password"`
	DataBases    Datebases `yaml:"DataBases"`
	Addr         string    `yaml:"Addr"`
	MaxIdleConns int       `yaml:"MaxIdleConns"`
	MaxOpenConns int       `yaml:"MaxOpenConns"`
	Config       string    `yaml:"Config"`
}

type WebInfo struct {
	Ypuser     string `yaml:"Ypuser"`
	YpPassword string `yaml:"YpPassword"`
	YpUrl      string `yaml:"YpUrl"`
	TTuser     string `yaml:"TTuser"`
	TTPassword string `yaml:"TTPassword"`
	TTUrl      string `yaml:"TTUrl"`
}

type Zap struct {
	Level         string `mapstructure:"level" json:"level" yaml:"level"`                           // 级别
	Format        string `mapstructure:"format" json:"format" yaml:"format"`                        // 输出
	Prefix        string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`                        // 日志前缀
	Director      string `mapstructure:"director" json:"director"  yaml:"director"`                 // 日志文件夹
	LinkName      string `mapstructure:"link-name" json:"linkName" yaml:"link-name"`                // 软链接名称
	ShowLine      bool   `mapstructure:"ShowLine" json:"showLine" yaml:"showLine"`                  // 显示行
	EncodeLevel   string `mapstructure:"EncodeLevel" json:"encodeLevel" yaml:"encode-level"`        // 编码级
	StacktraceKey string `mapstructure:"stacktrace-key" json:"stacktraceKey" yaml:"stacktrace-key"` // 栈名
	LogInConsole  bool   `mapstructure:"log-in-console" json:"logInConsole" yaml:"log-in-console"`  // 输出控制台
}
