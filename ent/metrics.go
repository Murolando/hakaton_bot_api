package ent

type DataBaseStat struct {
	// database_stat
	RollabackNum     int     `json:"rollback-num"`
	CommitNum        int     `json:"commit-num"`
	SumActiveTime    float64 `json:"active_time"`
	SumSessions      int64   `json:"sessions"`
	SumInsertedDatas int64   `json:"tup_inserted"`
	SumUpdatedDatas  int64   `json:"tup_updated"`
	SumDeletedDatas  int64   `json:"tup_deleted"`
}
type PgActivityStat struct {
	ActualSessions int `json:"actual-sessions"`
	AllSessions    int `json:"all-sessions"`
}
type PgDatabaseSize struct {
	CurrentMem int `json:"current-memory"`
}
type AllMetrics struct {
	DataBaseStat
	PgActivityStat
	PgDatabaseSize

	// 	select count(*)
	// from pg_stat_activity
	// where state = 'active'

	// 	select count(*)
	// from pg_stat_activity
}
type Database struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"dbname"`
}
